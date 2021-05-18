package mongo

import (
	"context"
	"errors"
	//"net/http"
	"time"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dpMongodb "github.com/ONSdigital/dp-mongodb"
	dpMongoLock "github.com/ONSdigital/dp-mongodb/dplock"
	dpMongoHealth "github.com/ONSdigital/dp-mongodb/health"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/log.go/log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// jobs collection name
const jobsCol = "jobs"

// locked jobs collection name
const jobsLockCol = "jobs_locks"

// Mongo represents a simplistic MongoDB configuration, with session, health and lock clients
type Mongo struct {
	Collection   string
	Database     string
	Session      *mgo.Session
	URI          string
	client       *dpMongoHealth.Client
	healthClient *dpMongoHealth.CheckMongoClient
	lockClient   *dpMongoLock.Lock
}

// Init creates a new mgo.Session with a strong consistency and a write mode of "majority".
func (m *Mongo) Init(ctx context.Context) (err error) {
	if m.Session != nil {
		return errors.New("session already exists")
	}

	// Create session
	if m.Session, err = mgo.Dial(m.URI); err != nil {
		return err
	}
	m.Session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	m.Session.SetMode(mgo.Strong, true)

	databaseCollectionBuilder := make(map[dpMongoHealth.Database][]dpMongoHealth.Collection)
	databaseCollectionBuilder[(dpMongoHealth.Database)(m.Database)] = []dpMongoHealth.Collection{(dpMongoHealth.Collection)(m.Collection), (dpMongoHealth.Collection)(jobsLockCol)}
	// Create client and healthclient from session
	m.client = dpMongoHealth.NewClientWithCollections(m.Session, databaseCollectionBuilder)
	m.healthClient = &dpMongoHealth.CheckMongoClient{
		Client:      *m.client,
		Healthcheck: m.client.Healthcheck,
	}

	// Create MongoDB lock client, which also starts the purger loop
	m.lockClient = dpMongoLock.New(ctx, m.Session, m.Database, jobsCol)
	return nil
}

// AcquireJobLock tries to lock the provided jobID.
// If the job is already locked, this function will block until it's released,
// at which point we acquire the lock and return.
func (m *Mongo) AcquireJobLock(ctx context.Context, jobID string) (lockID string, err error) {
	return m.lockClient.Acquire(ctx, jobID)
}

// UnlockJob releases an exclusive mongoDB lock for the provided lockId (if it exists)
func (m *Mongo) UnlockJob(lockID string) error {
	return m.lockClient.Unlock(lockID)
}

// Close closes the mongo session and returns any error
func (m *Mongo) Close(ctx context.Context) error {
	m.lockClient.Close(ctx)
	return dpMongodb.Close(ctx, m.Session)
}

// Checker is called by the healthcheck library to check the health state of this mongoDB instance
func (m *Mongo) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	return m.healthClient.Checker(ctx, state)
}

// GetJobs retrieves all jobs documents corresponding to the provided collectionID
func (m *Mongo) GetJobs(ctx context.Context, collectionID string) ([]models.Jobs, error) {
	s := m.Session.Copy()
	defer s.Close()
	log.Event(ctx, "getting jobs for collectionID", log.Data{"collectionID": collectionID})

	// Filter by collectionID, if provided
	colIDFilter := make(bson.M)
	if collectionID != "" {
		colIDFilter["collection_id"] = collectionID
	}

	iter := s.DB(m.Database).C(jobsCol).Find(colIDFilter).Iter()
	defer func() {
		err := iter.Close()
		if err != nil {
			log.Event(ctx, "error closing iterator", log.ERROR, log.Error(err))
		}
	}()

	results := []models.Jobs{}
	if err := iter.All(&results); err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.New("failed to find job in job store")
		}
		return nil, err
	}

	return results, nil
}

// GetJob retrieves a job document by its ID
func (m *Mongo) GetJob(ctx context.Context, id string) (*models.Job, error) {
	s := m.Session.Copy()
	defer s.Close()
	log.Event(ctx, "getting job by ID", log.Data{"id": id})

	var job models.Job
	err := s.DB(m.Database).C(jobsCol).Find(bson.M{"_id": id}).One(&job)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.New("failed to find job in job store")
		}
		return nil, err
	}

	return &job, nil
}

// UpdateJob updates an existing job document
func (m *Mongo) UpdateJob(ctx context.Context, id string, job *models.Job) (bool, error) {
	s := m.Session.Copy()
	defer s.Close()
	log.Event(ctx, "updating job", log.Data{"id": id})

	updates := createJobUpdateQuery(ctx, id, job)
	if len(updates) == 0 {
		log.Event(ctx, "nothing to update")
		return false, nil
	}

	update := bson.M{"$set": updates, "$setOnInsert": bson.M{"last_updated": time.Now()}}
	if err := s.DB(m.Database).C(jobsCol).UpdateId(id, update); err != nil {
		if err == mgo.ErrNotFound {
			return false, errors.New("failed to find job in job store")
		}
		return false, err
	}

	return true, nil
}

// createJobUpdateQuery generates the bson model to update an job with the provided job update.
// Fields present in mongoDB will not be deleted if they are not present in the job update object.
func createJobUpdateQuery(ctx context.Context, id string, job *models.Job) bson.M {
	updates := make(bson.M)

	log.Event(ctx, "building update query for job resource", log.INFO, log.INFO, log.Data{"job_id": id, "job": job, "updates": updates})

	if job.CollectionID != "" {
		updates["collection_id"] = job.CollectionID
	}
	if job.State != "" {
		updates["state"] = job.State
	}
	//if job.Error != "" {
	//	updates["error"] = job.Error
	//}

	//if job.Upload != nil {
	//	if job.Upload.Path != "" {
	//		updates["upload"] = job.Upload
	//	}
	//}

	return updates
}

// UpsertJob adds or overides an existing job document
func (m *Mongo) UpsertJob(ctx context.Context, id string, job *models.Job) (err error) {
	s := m.Session.Copy()
	defer s.Close()
	log.Event(ctx, "upserting job", log.Data{"id": id})

	update := bson.M{
		"$set": job,
		"$setOnInsert": bson.M{
			"last_updated": time.Now(),
		},
	}

	_, err = s.DB(m.Database).C(jobsCol).UpsertId(id, update)
	return
}
