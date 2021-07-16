// Package steps is used to define the steps that are used in the component test, which is written in godog (Go's version of cucumber).
package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	componentTest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-reindex-api/config"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/mongo"
	"github.com/ONSdigital/dp-search-reindex-api/service"
	serviceMock "github.com/ONSdigital/dp-search-reindex-api/service/mock"
	"github.com/benweissmann/memongo"
	"github.com/cucumber/godog"
	"github.com/pkg/errors"
	"github.com/rdumont/assistdog"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// jobs collection name
const jobsCol = "jobs"

// job id that gets generated by the POST /jobs endpoint
var id = ""

// JobsFeature is a type that contains all the requirements for running a godog (cucumber) feature that tests the /jobs endpoint.
type JobsFeature struct {
	ErrorFeature   componentTest.ErrorFeature
	svc            *service.Service
	errorChan      chan error
	Config         *config.Config
	HTTPServer     *http.Server
	ServiceRunning bool
	ApiFeature     *componentTest.APIFeature
	responseBody   []byte
	MongoClient    *mongo.JobStore
	MongoFeature   *componentTest.MongoFeature
}

// NewJobsFeature returns a pointer to a new JobsFeature, which can then be used for testing the /jobs endpoint.
func NewJobsFeature(mongoFeature *componentTest.MongoFeature) (*JobsFeature, error) {
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
	mongodb := &mongo.JobStore{
		Collection: jobsCol,
		Database:   memongo.RandomDatabase(),
		URI:        mongoFeature.Server.URI(),
	}
	ctx := context.Background()
	if err := mongodb.Init(ctx); err != nil {
		return nil, err
	}

	f.MongoClient = mongodb

	err = runJobsFeatureService(f, err, ctx, cfg, svcErrors)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func runJobsFeatureService(f *JobsFeature, err error, ctx context.Context, cfg *config.Config, svcErrors chan error) error {
	initFunctions := &serviceMock.InitialiserMock{
		DoGetHealthCheckFunc: f.DoGetHealthcheckOk,
		DoGetHTTPServerFunc:  f.DoGetHTTPServer,
		DoGetMongoDBFunc:     f.DoGetMongoDB,
	}

	serviceList := service.NewServiceList(initFunctions)
	f.svc, err = service.Run(ctx, cfg, serviceList, "1", "", "", svcErrors)
	return err
}

// InitAPIFeature initialises the ApiFeature that's contained within a specific JobsFeature.
func (f *JobsFeature) InitAPIFeature() *componentTest.APIFeature {
	f.ApiFeature = componentTest.NewAPIFeature(f.InitialiseService)

	return f.ApiFeature
}

// RegisterSteps defines the steps within a specific JobsFeature cucumber test.
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
	ctx.Step(`^no jobs have been generated in the Job Store$`, f.noJobsHaveBeenGeneratedInTheJobStore)
	ctx.Step(`^I call GET \/jobs\/{"([^"]*)"} using a valid UUID$`, f.iCallGETJobsUsingAValidUUID)
	ctx.Step(`^the response should contain the new number of tasks$`, f.theResponseShouldContainTheNewNumberOfTasks)
	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{(\d+)} using the generated id$`, f.iCallPUTJobsidnumber_of_tasksUsingTheGeneratedId)
	ctx.Step(`^I would expect the response to be an empty list$`, f.iWouldExpectTheResponseToBeAnEmptyList)
	ctx.Step(`^I call PUT \/jobs\/{"([^"]*)"}\/number_of_tasks\/{(\d+)} using a valid UUID$`, f.iCallPUTJobsNumber_of_tasksUsingAValidUUID)
	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{"([^"]*)"} using the generated id with an invalid count$`, f.iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithAnInvalidCount)
	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{"([^"]*)"} using the generated id with a negative count$`, f.iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithANegativeCount)
	ctx.Step(`^the search reindex api loses its connection to mongo DB$`, f.theSearchReindexApiLosesItsConnectionToMongoDB)
}

// Reset sets the resources within a specific JobsFeature back to their default values.
func (f *JobsFeature) Reset(mongoFail bool) *JobsFeature {
	if mongoFail {
		f.MongoClient.Database = "lost database connection"
	} else {
		f.MongoClient.Database = memongo.RandomDatabase()
	}
	ctx := context.Background()
	err := f.MongoClient.Init(ctx)
	if err != nil {
		return nil
	}
	return f
}

// Close stops the *service.Service, which is pointed to from within the specific JobsFeature, from running.
func (f *JobsFeature) Close() error {
	if f.svc != nil && f.ServiceRunning {
		err := f.svc.Close(context.Background())
		if err != nil {
			return err
		}
		f.ServiceRunning = false
	}
	return nil
}

// InitialiseService returns the http.Handler that's contained within a specific JobsFeature.
func (f *JobsFeature) InitialiseService() (http.Handler, error) {
	return f.HTTPServer.Handler, nil
}

// DoGetHTTPServer takes a bind Address (string) and a router (http.Handler), which are used to set up an HTTPServer.
// The HTTPServer is in a specific JobsFeature and is returned.
func (f *JobsFeature) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	f.HTTPServer.Addr = bindAddr
	f.HTTPServer.Handler = router
	return f.HTTPServer
}

// DoGetHealthcheckOk returns a mock HealthChecker service for a specific JobsFeature.
func (f *JobsFeature) DoGetHealthcheckOk(cfg *config.Config, time string, commit string, version string) (service.HealthChecker, error) {
	return &serviceMock.HealthCheckerMock{
		AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
		StartFunc:    func(ctx context.Context) {},
		StopFunc:     func() {},
	}, nil
}

// DoGetMongoDB returns a MongoDB, for the component test, which has a random database name and different URI to the one used by the API under test.
func (f *JobsFeature) DoGetMongoDB(ctx context.Context, cfg *config.Config) (service.MongoJobStorer, error) {
	return f.MongoClient, nil
}

// iWouldExpectIdLast_updatedAndLinksToHaveThisStructure is a feature step that can be defined for a specific JobsFeature.
// It takes a table that contains the expected structures for id, last_updated, and links values. And it asserts whether or not these are found.
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

	id = response.ID
	lastUpdated := response.LastUpdated
	links := response.Links

	err = f.checkStructure(id, lastUpdated, expectedResult, links)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// checkStructure is a utility method that can be called by a feature step to assert that a job contains the expected structure in its values of
// id, last_updated, and links. It confirms that last_updated is a current or past time, and that the tasks and self links have the correct paths.
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

// theResponseShouldAlsoContainTheFollowingValues is a feature step that can be defined for a specific JobsFeature.
// It takes a table that contains the expected values for all the remaining attributes, of a Job resource, and it asserts whether or not these are found.
func (f *JobsFeature) theResponseShouldAlsoContainTheFollowingValues(table *godog.Table) error {
	expectedResult, err := assistdog.NewDefault().ParseMap(table)
	if err != nil {
		panic(err)
	}
	var response models.Job

	_ = json.Unmarshal(f.responseBody, &response)

	f.checkValuesInJob(expectedResult, response)

	return f.ErrorFeature.StepError()
}

// checkValuesInJob is a utility method that can be called by a feature step in order to check that the values
// of certain attributes, in a job, are all equal to the expected ones.
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

// iHaveGeneratedAJobInTheJobStore is a feature step that can be defined for a specific JobsFeature.
// It calls POST /jobs with an empty body, which causes a default job resource to be generated.
// The newly created job resource is stored in the Job Store and also returned in the response body.
func (f *JobsFeature) iHaveGeneratedAJobInTheJobStore() error {
	// call POST /jobs
	err := f.callPostJobs()
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// callPostJobs is a utility method that can be called by a feature step in order to call the POST jobs/ endpoint
// Calling that endpoint results in the creation of a job, in the Job Store, containing a unique id and default values.
func (f *JobsFeature) callPostJobs() error {
	var emptyBody = godog.DocString{}
	err := f.ApiFeature.IPostToWithBody("/jobs", &emptyBody)
	if err != nil {
		return err
	}

	return err
}

// iCallGETJobsidUsingTheGeneratedId is a feature step that can be defined for a specific JobsFeature.
// It gets the id from the response body, generated in the previous step, and then uses this to call GET /jobs/{id}.
func (f *JobsFeature) iCallGETJobsidUsingTheGeneratedId() error {

	if id == "" {
		f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)

		var response models.Job

		err := json.Unmarshal(f.responseBody, &response)
		if err != nil {
			return err
		}

		id = response.ID
	}

	err := f.GetJobByID(id)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iHaveGeneratedThreeJobsInTheJobStore is a feature step that can be defined for a specific JobsFeature.
// It calls POST /jobs with an empty body, three times, which causes three default job resources to be generated.
func (f *JobsFeature) iHaveGeneratedThreeJobsInTheJobStore() error {
	// call POST /jobs three times
	err := f.callPostJobs()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	err = f.callPostJobs()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	err = f.callPostJobs()
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iWouldExpectThereToBeThreeOrMoreJobsReturnedInAList is a feature step that can be defined for a specific JobsFeature.
// It checks the response from calling GET /jobs to make sure that a list containing three or more jobs has been returned.
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

// inEachJobIWouldExpectIdLast_updatedAndLinksToHaveThisStructure is a feature step that can be defined for a specific JobsFeature.
// It checks the response from calling GET /jobs to make sure that each job contains the expected types of values of id,
// last_updated, and links.
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

// eachJobShouldAlsoContainTheFollowingValues is a feature step that can be defined for a specific JobsFeature.
// It checks the response from calling GET /jobs to make sure that each job contains the expected values of
// all the remaining attributes of a job.
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

func (f *JobsFeature) theResponseShouldContainTheNewNumberOfTasks(table *godog.Table) error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)
	expectedResult, err := assistdog.NewDefault().ParseMap(table)
	if err != nil {
		return err
	}
	var response models.Job

	_ = json.Unmarshal(f.responseBody, &response)

	assert.Equal(&f.ErrorFeature, expectedResult["number_of_tasks"], strconv.Itoa(response.NumberOfTasks))

	return f.ErrorFeature.StepError()
}

// theJobsShouldBeOrderedByLast_updatedWithTheOldestFirst is a feature step that can be defined for a specific JobsFeature.
// It checks the response from calling GET /jobs to make sure that the jobs are in ascending order of their last_updated
// times i.e. the most recently updated is last in the list.
func (f *JobsFeature) theJobsShouldBeOrderedByLast_updatedWithTheOldestFirst() error {
	var response models.Jobs
	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}
	jobList := response.JobList
	timeToCheck := jobList[0].LastUpdated

	for j := 1; j < len(jobList); j++ {
		index := strconv.Itoa(j - 1)
		nextIndex := strconv.Itoa(j)
		nextTime := jobList[j].LastUpdated
		assert.True(&f.ErrorFeature, timeToCheck.Before(nextTime),
			"The value of last_updated at job_list["+index+"] should be earlier than that at job_list["+nextIndex+"]")
		timeToCheck = nextTime
	}
	return f.ErrorFeature.StepError()
}

// noJobsHaveBeenGeneratedInTheJobStore is a feature step that can be defined for a specific JobsFeature.
// It resets the Job Store to its default values, which means that it will contain no jobs.
func (f *JobsFeature) noJobsHaveBeenGeneratedInTheJobStore() error {
	f.Reset(false)
	return nil
}

// iCallGETJobsUsingAValidUUID is a feature step that can be defined for a specific JobsFeature.
// It calls GET /jobs/{id} using the id passed in, which should be a valid UUID.
func (f *JobsFeature) iCallGETJobsUsingAValidUUID(id string) error {
	err := f.GetJobByID(id)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iWouldExpectTheResponseToBeAnEmptyList is a feature step that can be defined for a specific JobsFeature.
// It checks the response from calling GET /jobs to make sure that an empty list (0 jobs) has been returned.
func (f *JobsFeature) iWouldExpectTheResponseToBeAnEmptyList() error {

	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)

	var response models.Jobs
	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}
	numJobsFound := len(response.JobList)
	assert.True(&f.ErrorFeature, numJobsFound == 0, "The list should contain no jobs but it contains "+strconv.Itoa(numJobsFound))

	return f.ErrorFeature.StepError()
}

// iCallPUTJobsidnumber_of_taskscountUsingTheGeneratedId is a feature step that can be defined for a specific JobsFeature.
// It gets the id from the response body, generated in the previous step, and then uses this to call PUT /jobs/{id}/number_of_tasks/{count}
func (f *JobsFeature) iCallPUTJobsidnumber_of_tasksUsingTheGeneratedId(count int) error {

	countStr := strconv.Itoa(count)
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)
	var response models.Job

	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	id = response.ID
	err = f.PutNumberOfTasks(countStr)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iCallPUTJobsNumber_of_tasksUsingAValidUUID is a feature step that can be defined for a specific JobsFeature.
// It uses the parameters passed in to call PUT /jobs/{id}/number_of_tasks/{count}
func (f *JobsFeature) iCallPUTJobsNumber_of_tasksUsingAValidUUID(idStr string, count int) error {

	countStr := strconv.Itoa(count)
	id = idStr

	err := f.PutNumberOfTasks(countStr)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithAnInvalidCount is a feature step that can be defined for a specific JobsFeature.
// It gets the id from the response body, generated in the previous step, and then uses this to call PUT /jobs/{id}/number_of_tasks/{invalidCount}
func (f *JobsFeature) iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithAnInvalidCount(invalidCount string) error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)
	var response models.Job

	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	id = response.ID

	err = f.PutNumberOfTasks(invalidCount)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithANegativeCount is a feature step that can be defined for a specific JobsFeature.
// It gets the id from the response body, generated in the previous step, and then uses this to call PUT /jobs/{id}/number_of_tasks/{negativeCount}
func (f *JobsFeature) iCallPUTJobsidnumber_of_tasksUsingTheGeneratedIdWithANegativeCount(negativeCount string) error {
	f.responseBody, _ = ioutil.ReadAll(f.ApiFeature.HttpResponse.Body)
	var response models.Job

	err := json.Unmarshal(f.responseBody, &response)
	if err != nil {
		return err
	}

	id = response.ID

	err = f.PutNumberOfTasks(negativeCount)
	if err != nil {
		return err
	}

	return f.ErrorFeature.StepError()
}

// theSearchReindexApiLosesItsConnectionToMongoDB is a feature step that can be defined for a specific JobsFeature.
// It loses the connection to mongo DB by setting the mongo database to an invalid setting (in the Reset function).
func (f *JobsFeature) theSearchReindexApiLosesItsConnectionToMongoDB() error {
	f.Reset(true)
	return nil
}

// GetJobByID is a utility function that is used for calling the GET /jobs/{id} endpoint.
// It checks that the id string is a valid UUID before calling the endpoint.
func (f *JobsFeature) GetJobByID(id string) error {
	_, err := uuid.FromString(id)
	if err != nil {
		fmt.Println("Got uuid: " + id)
		return err
	}

	// call GET /jobs/{id}
	err = f.ApiFeature.IGet("/jobs/" + id)
	if err != nil {
		os.Exit(1)
	}
	return nil
}

// PutNumberOfTasks is a utility function that is used for calling the PUT /jobs/{id}/number_of_tasks/{count}
// It checks that the id string is a valid UUID before calling the endpoint.
func (f *JobsFeature) PutNumberOfTasks(countStr string) error {

	var emptyBody = godog.DocString{}
	_, err := uuid.FromString(id)
	if err != nil {
		fmt.Println("Got uuid: " + id)
		return err
	}

	// call PUT /jobs/{id}/number_of_tasks/{count}
	err = f.ApiFeature.IPut("/jobs/"+id+"/number_of_tasks/"+countStr, &emptyBody)
	if err != nil {
		os.Exit(1)
	}
	return nil
}
