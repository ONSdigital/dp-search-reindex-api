package etag

import (
	"context"
	"testing"
	"time"

	"github.com/ONSdigital/dp-search-reindex-api/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateETagForJobPatch(t *testing.T) {
	testCtx := context.Background()
	currentJob := getTestJob()

	Convey("Given an existing job with no new job updates", t, func() {

		Convey("When GenerateETagForJob is called", func() {
			newETag, err := GenerateETagForJob(testCtx, currentJob)

			Convey("Then an eTag is returned", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And the eTag should be the same as the existing eTag", func() {
					So(newETag, ShouldEqual, currentJob.ETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Given an updated job", t, func() {
		updatedJob := currentJob
		updatedJob.State = "completed"

		Convey("When GenerateETagForJob is called", func() {
			newETag, err := GenerateETagForJob(testCtx, updatedJob)

			Convey("Then a new eTag is created", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And new eTag should not be the same as the existing eTag", func() {
					So(newETag, ShouldNotEqual, currentJob.ETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}

func getTestJob() models.Job {
	zeroTime := time.Time{}.UTC()

	return models.Job{
		ID:          "test1",
		ETag:        `"56b6890f1321590998d5fd8d293b620581ff3531"`,
		LastUpdated: time.Now().UTC(),
		Links: &models.JobLinks{
			Tasks: "http://localhost:25700/jobs/test1234/tasks",
			Self:  "http://localhost:25700/jobs/test1234",
		},
		NumberOfTasks:                0,
		ReindexCompleted:             zeroTime,
		ReindexFailed:                zeroTime,
		ReindexStarted:               zeroTime,
		SearchIndexName:              "Test Search Index Name",
		State:                        models.JobCreatedState,
		TotalSearchDocuments:         0,
		TotalInsertedSearchDocuments: 0,
	}
}
