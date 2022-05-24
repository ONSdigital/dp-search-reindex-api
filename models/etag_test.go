package models_test

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

	Convey("Given an updated job", t, func() {
		updatedJob := currentJob
		updatedJob.State = "completed"

		Convey("When GenerateETagForJob is called", func() {
			newETag, err := models.GenerateETagForJob(testCtx, updatedJob)

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

	Convey("Given an existing job with no new job updates", t, func() {
		Convey("When GenerateETagForJob is called", func() {
			newETag, err := models.GenerateETagForJob(testCtx, currentJob)

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
}

func getTestJob() models.Job {
	zeroTime := time.Time{}.UTC()

	return models.Job{
		ID:          "test1",
		ETag:        `"56b6890f1321590998d5fd8d293b620581ff3531"`,
		LastUpdated: zeroTime,
		Links: &models.JobLinks{
			Tasks: "http://localhost:25700/jobs/test1234/tasks",
			Self:  "http://localhost:25700/jobs/test1234",
		},
		NumberOfTasks:                0,
		ReindexCompleted:             zeroTime,
		ReindexFailed:                zeroTime,
		ReindexStarted:               zeroTime,
		SearchIndexName:              "Test Search Index Name",
		State:                        models.JobStateCreated,
		TotalSearchDocuments:         0,
		TotalInsertedSearchDocuments: 0,
	}
}

func TestGenerateETagForJobs(t *testing.T) {
	ctx := context.Background()

	job := getTestJob()
	jobs := models.Jobs{
		Count:      1,
		JobList:    []models.Job{job},
		Limit:      1,
		Offset:     0,
		TotalCount: 1,
	}

	jobsETag := `"89ffa9ad7f168112ad68de35a7bef9c64535b019"`

	Convey("Given the list of jobs has updated", t, func() {
		updatedJobs := jobs
		updatedJobs.Limit = 10

		Convey("When GenerateETagForJobs is called", func() {
			newETag, err := models.GenerateETagForJobs(ctx, updatedJobs)

			Convey("Then a new eTag is created", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And new eTag should not be the same as the old eTag", func() {
					So(newETag, ShouldNotEqual, jobsETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Given the list of jobs has not updated", t, func() {
		Convey("When GenerateETagForJobs is called", func() {
			newETag, err := models.GenerateETagForJobs(ctx, jobs)

			Convey("Then an eTag is returned", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And the eTag should be the same as the existing eTag", func() {
					So(newETag, ShouldEqual, jobsETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}

func TestGenerateETagForTask(t *testing.T) {
	currentTask := getTestTask()

	Convey("Given an updated task", t, func() {
		updatedTask := currentTask
		updatedTask.TaskName = "zebedee"

		Convey("When GenerateETagForTask is called", func() {
			newETag, err := models.GenerateETagForTask(updatedTask)

			Convey("Then a new eTag is created", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And new eTag should not be the same as the existing eTag", func() {
					So(newETag, ShouldNotEqual, currentTask.ETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Given an existing task with no new updates", t, func() {
		Convey("When GenerateETagForTask is called", func() {
			newETag, err := models.GenerateETagForTask(currentTask)

			Convey("Then an eTag is returned", func() {
				So(newETag, ShouldNotBeEmpty)

				Convey("And the eTag should be the same as the existing eTag", func() {
					So(newETag, ShouldEqual, currentTask.ETag)

					Convey("And no errors should be returned", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}

func getTestTask() models.Task {
	return models.Task{
		ETag:              `"c644f142e428485848c5272759bdd216b5d7560e"`,
		JobID:             "task1234",
		TaskName:          "task",
		NumberOfDocuments: 3,
	}
}
