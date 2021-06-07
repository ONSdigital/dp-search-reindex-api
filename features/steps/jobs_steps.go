//Package steps is used to define the steps that are used in the component test, which is written in godog (Go's version of cucumber).
package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-search-reindex-api/mongo"
	"github.com/benweissmann/memongo"
	"os"
	"strconv"
	"strings"
	"time"

	"io/ioutil"
	"net/http"

	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-reindex-api/config"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/service"
	"github.com/ONSdigital/dp-search-reindex-api/service/mock"
	"github.com/cucumber/godog"
	//"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"github.com/rdumont/assistdog"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

//JobID1 variable is required to make sure that the relevant job will be deleted from mongoDB when the test has finished
var JobID1 string

//JobsFeature is a type that contains all the requirements for running a godog (cucumber) feature that tests the /jobs endpoint.
type JobsFeature struct {
	ErrorFeature   componenttest.ErrorFeature
	svc            *service.Service
	errorChan      chan error
	Config         *config.Config
	HTTPServer     *http.Server
	ServiceRunning bool
	ApiFeature     *componenttest.APIFeature
	responseBody   []byte
	MongoClient    *mongo.MgoDataStore
	MongoFeature   *componenttest.MongoFeature
}

//NewJobsFeature returns a pointer to a new JobsFeature, which can then be used for testing the /jobs endpoint.
func NewJobsFeature(mongoFeature *componenttest.MongoFeature) (*JobsFeature, error) {
	f := &JobsFeature{
		HTTPServer:     &http.Server{},
		errorChan:      make(chan error),
		ServiceRunning: false,
	}
	svcErrors := make(chan error, 1)
	cfg, err := config.Get()
	if err != nil {
		return nil, err
	}
	mongodb := &mongo.MgoDataStore{
		Collection:  "jobs",
		Database:    memongo.RandomDatabase(),
		URI:         mongoFeature.Server.URI(),
	}
	ctx := context.Background()
	if err := mongodb.Init(ctx); err != nil {
		return nil, err
	}

	f.MongoClient = mongodb
	initFunctions := &mock.InitialiserMock{
		DoGetHealthCheckFunc: f.DoGetHealthcheckOk,
		DoGetHTTPServerFunc:  f.DoGetHTTPServer,
		DoGetMongoDBFunc:     f.DoGetMongoDB,
	}

	serviceList := service.NewServiceList(initFunctions)
	f.svc, err = service.Run(ctx, cfg, serviceList, "1", "", "", svcErrors)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//InitAPIFeature initialises the ApiFeature that's contained within a specific JobsFeature.
func (f *JobsFeature) InitAPIFeature() *componenttest.APIFeature {
	f.ApiFeature = componenttest.NewAPIFeature(f.InitialiseService)

	return f.ApiFeature
}

//RegisterSteps defines the steps within a specific JobsFeature cucumber test.
func (f *JobsFeature) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I would expect id, last_updated, and links to have this structure$`, f.iWouldExpectIdLast_updatedAndLinksToHaveThisStructure)
	ctx.Step(`^the response should also contain the following values:$`, f.theResponseShouldAlsoContainTheFollowingValues)
	ctx.Step(`^I have generated a job in the Job Store$`, f.iHaveGeneratedAJobInTheJobStore)
	ctx.Step(`^I call GET \/jobs\/{id} using the generated id$`, f.iCallGETJobsidUsingTheGeneratedId)
	ctx.Step(`^I have generated three jobs in the Job Store$`, f.iHaveGeneratedThreeJobsInTheJobStore)
	ctx.Step(`^I would expect there to be three or more jobs returned in a list$`, f.iWouldExpectThereToBeThreeOrMoreJobsReturnedInAList)
	ctx.Step(`^in each job I would expect id, last_updated, and links to have this structure$`, f.inEachJobIWouldExpectIdLast_updatedAndLinksToHaveThisStructure)
	ctx.Step(`^each job should also contain the following values:$`, f.eachJobShouldAlsoContainTheFollowingValues)
	ctx.Step(`^the jobs should be ordered, by last_updated, with the oldest first$`, f.theJobsShouldBeOrderedByLast_updatedWithTheOldestFirst)
}

//Reset sets the resources within a specific JobsFeature back to their default values.
func (f *JobsFeature) Reset() *JobsFeature {
	return f
}

//Close stops the *service.Service, which is pointed to from within the specific JobsFeature, from running.
func (f *JobsFeature) Close() error {
	if f.svc != nil && f.ServiceRunning {
		f.svc.Close(context.Background())
		f.ServiceRunning = false
	}
	return nil
}

//InitialiseService returns the http.Handler that's contained within a specific JobsFeature.
func (f *JobsFeature) InitialiseService() (http.Handler, error) {
	return f.HTTPServer.Handler, nil
}

//DoGetHTTPServer takes a bind Address (string) and a router (http.Handler), which are used to set up an HTTPServer.
//The HTTPServer is in a specific JobsFeature and is returned.
func (f *JobsFeature) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	f.HTTPServer.Addr = bindAddr
	f.HTTPServer.Handler = router
	return f.HTTPServer
}

//DoGetHealthcheckOk returns a HealthChecker service for a specific JobsFeature.
func (f *JobsFeature) DoGetHealthcheckOk(cfg *config.Config, time string, commit string, version string) (service.HealthChecker, error) {
	versionInfo, _ := healthcheck.NewVersionInfo(time, commit, version)
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)
	return &hc, nil
}

// DoGetMongoDB returns a MongoDB
func (f *JobsFeature) DoGetMongoDB(ctx context.Context, cfg *config.Config) (mongo.MgoJobStore, error) {
	mongodb := &mongo.MgoDataStore{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		URI:        cfg.MongoConfig.BindAddr,
	}
	if err := mongodb.Init(ctx); err != nil {
		return nil, err
	}
	return mongodb, nil
}

//iWouldExpectIdLast_updatedAndLinksToHaveThisStructure is a feature step that can be defined for a specific JobsFeature.
//It takes a table that contains the expected structures for id, last_updated, and links values. And it asserts whether or not these are found.
func (f *JobsFeature) iWouldExpectIdLast_updatedAndLinksToHaveThisStructure(table *godog.Table) error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)
	assist := assistdog.NewDefault()

	expectedResult, err := assist.ParseMap(table)
	if err != nil {
		panic(err)
	}

	var response models.Job

	err = json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	id := response.ID
	lastUpdated := response.LastUpdated
	links := response.Links

	err = f.checkStructure(id, lastUpdated, expectedResult, links)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

//checkStructure is a utility method that can be called by a feature step to assert that a job contains the expected structure in its values of
//id, last_updated, and links. It confirms that last_updated is a current or past time, and that the tasks and self links have the correct paths.
func (f *JobsFeature) checkStructure(id string, lastUpdated time.Time, expectedResult map[string]string, links *models.JobLinks) error {
	_, err := uuid.FromString(id)
	if err != nil {
		fmt.Println("Got uuid: " + id)
		return err
	}

	if lastUpdated.After(time.Now()) {
		return errors.New("expected LastUpdated to be now or earlier but it was: " + lastUpdated.String())
	}

	expectedLinksTasks := strings.Replace(expectedResult["links: tasks"], "{id}", id, 1)

	assert.Equal(&f.ErrorFeature, expectedLinksTasks, links.Tasks)

	expectedLinksSelf := strings.Replace(expectedResult["links: self"], "{id}", id, 1)

	assert.Equal(&f.ErrorFeature, expectedLinksSelf, links.Self)
	return nil
}

//theResponseShouldAlsoContainTheFollowingValues is a feature step that can be defined for a specific JobsFeature.
//It takes a table that contains the expected values for all the remaining attributes, of a Job resource, and it asserts whether or not these are found.
func (f *JobsFeature) theResponseShouldAlsoContainTheFollowingValues(table *godog.Table) error {
	expectedResult, err := assistdog.NewDefault().ParseMap(table)
	if err != nil {
		panic(err)
	}
	var response models.Job

	_ = json.Unmarshal(f.responseBody, &response)

	f.checkValuesInJob(expectedResult, response)

	//Now tidy up by deleting the test job from mongoDB jobs collection
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//f.MongoFeature.Database.Collection("jobs").DeleteOne(ctx, bson.M{"id": JobID1})

	return f.ErrorFeature.StepError()
}

//checkValuesInJob is a utility method that can be called by a feature step in order to check that the values
//of certain attributes, in a job, are all equal to the expected ones.
func (f *JobsFeature) checkValuesInJob(expectedResult map[string]string, job models.Job) {
	assert.Equal(&f.ErrorFeature, expectedResult["number_of_tasks"], strconv.Itoa(job.NumberOfTasks))
	assert.Equal(&f.ErrorFeature, expectedResult["reindex_completed"], job.ReindexCompleted.Format(time.RFC3339))
	assert.Equal(&f.ErrorFeature, expectedResult["reindex_failed"], job.ReindexFailed.Format(time.RFC3339))
	assert.Equal(&f.ErrorFeature, expectedResult["reindex_started"], job.ReindexStarted.Format(time.RFC3339))
	assert.Equal(&f.ErrorFeature, expectedResult["search_index_name"], job.SearchIndexName)
	assert.Equal(&f.ErrorFeature, expectedResult["state"], job.State)
	assert.Equal(&f.ErrorFeature, expectedResult["total_search_documents"], strconv.Itoa(job.TotalSearchDocuments))
	assert.Equal(&f.ErrorFeature, expectedResult["total_inserted_search_documents"], strconv.Itoa(job.TotalInsertedSearchDocuments))
}

//iHaveGeneratedAJobInTheJobStore is a feature step that can be defined for a specific JobsFeature.
//It calls POST /jobs with an empty body, which causes a default job resource to be generated.
//The newly created job resource is stored in the Job Store and also returned in the response body.
func (f *JobsFeature) iHaveGeneratedAJobInTheJobStore() error {
	//call POST /jobs
	f.callPostJobs()

	return f.ErrorFeature.StepError()
}

//callPostJobs is a utility method that can be called by a feature step in order to call the POST jobs/ endpoint
//Calling that endpoint results in the creation of a job, in the Job Store, containing a unique id and default values.
func (f *JobsFeature) callPostJobs() error {
	var emptyBody = godog.DocString{}
	err := f.ApiFeature.IPostToWithBody("/jobs", &emptyBody)
	if err != nil {
		return err
	}

	return err
}

//iCallGETJobsidUsingTheGeneratedId is a feature step that can be defined for a specific JobsFeature.
//It gets the id from the response body, generated in the previous step, and then uses this to call GET /jobs/{id}.
func (f *JobsFeature) iCallGETJobsidUsingTheGeneratedId() error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)

	var response models.Job

	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	JobID1 = response.ID
		fmt.Println("JobID1 is: " + JobID1)
	_, err = uuid.FromString(JobID1)
	if err != nil {
		fmt.Println("Got uuid: " + JobID1)
		return err
	}

	//call GET /jobs/{id}
	err = f.ApiFeature.IGet("/jobs/" + JobID1)
	if err != nil {
		os.Exit(1)
	}

	return f.ErrorFeature.StepError()
}

//iHaveGeneratedThreeJobsInTheJobStore is a feature step that can be defined for a specific JobsFeature.
//It calls POST /jobs with an empty body, three times, which causes three default job resources to be generated.
func (f *JobsFeature) iHaveGeneratedThreeJobsInTheJobStore() error {
	//call POST /jobs three times
	f.callPostJobs()
	f.callPostJobs()
	f.callPostJobs()

	return f.ErrorFeature.StepError()
}

//iWouldExpectThereToBeThreeOrMoreJobsReturnedInAList is a feature step that can be defined for a specific JobsFeature.
//It checks the response from calling GET /jobs to make sure that a list containing three or more jobs has been returned.
func (f *JobsFeature) iWouldExpectThereToBeThreeOrMoreJobsReturnedInAList() error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)

	var response models.Jobs
	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}
	numJobsFound := len(response.JobList)
	assert.True(&f.ErrorFeature, numJobsFound >= 3, "The list should contain three or more jobs but it only contains "+strconv.Itoa(numJobsFound))

	return f.ErrorFeature.StepError()
}

//inEachJobIWouldExpectIdLast_updatedAndLinksToHaveThisStructure is a feature step that can be defined for a specific JobsFeature.
//It checks the response from calling GET /jobs to make sure that each job contains the expected types of values of id,
//last_updated, and links.
func (f *JobsFeature) inEachJobIWouldExpectIdLast_updatedAndLinksToHaveThisStructure(table *godog.Table) error {
	assist := assistdog.NewDefault()
	expectedResult, err := assist.ParseMap(table)
	if err != nil {
		return err
	}
	var response models.Jobs

	err = json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	for j := range response.JobList {
		job := response.JobList[j]
		err := f.checkStructure(job.ID, job.LastUpdated, expectedResult, job.Links)
		if err != nil {
			return err
		}

	}

	return f.ErrorFeature.StepError()
}

//eachJobShouldAlsoContainTheFollowingValues is a feature step that can be defined for a specific JobsFeature.
//It checks the response from calling GET /jobs to make sure that each job contains the expected values of
//all the remaining attributes of a job.
func (f *JobsFeature) eachJobShouldAlsoContainTheFollowingValues(table *godog.Table) error {
	expectedResult, err := assistdog.NewDefault().ParseMap(table)
	if err != nil {
		return err
	}
	var response models.Jobs

	err = json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	for _, job := range response.JobList {
		f.checkValuesInJob(expectedResult, job)
	}

	return f.ErrorFeature.StepError()
}

//theJobsShouldBeOrderedByLast_updatedWithTheOldestFirst is a feature step that can be defined for a specific JobsFeature.
//It checks the response from calling GET /jobs to make sure that the jobs are in ascending order of their last_updated
//times i.e. the most recently updated is last in the list.
func (f *JobsFeature) theJobsShouldBeOrderedByLast_updatedWithTheOldestFirst() error {
	var response models.Jobs
	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}
	job_list := response.JobList
	timeToCheck := job_list[0].LastUpdated

	for j := 1; j < len(job_list); j++ {
		index := strconv.Itoa(j - 1)
		nextIndex := strconv.Itoa(j)
		nextTime := job_list[j].LastUpdated
		assert.True(&f.ErrorFeature, timeToCheck.Before(nextTime),
			"The value of last_updated at job_list["+index+"] should be earlier than that at job_list["+nextIndex+"]")
		timeToCheck = nextTime
	}
	return f.ErrorFeature.StepError()
}
