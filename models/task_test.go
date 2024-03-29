package models

import (
	"context"
	"testing"
	"time"

	"github.com/ONSdigital/dp-search-reindex-api/apierrors"
	. "github.com/smartystreets/goconvey/convey"
)

var taskNamesTest = map[string]bool{
	"zebedee":     true,
	"dataset_api": true,
}

func TestParseTaskName(t *testing.T) {
	Convey("Given a task name which is available in list of task names", t, func() {
		givenTaskName := "zebedee"

		Convey("When ParseTaskName is called", func() {
			err := ParseTaskName(givenTaskName, taskNamesTest)

			Convey("Then no error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a invalid task name which is not available in list of task names", t, func() {
		givenTaskName := "invalid"

		Convey("When ParseTaskName is called", func() {
			err := ParseTaskName(givenTaskName, taskNamesTest)

			Convey("Then invalid task name error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, apierrors.ErrTaskInvalidName)
			})
		})
	})
}

func TestNewTask(t *testing.T) {
	Convey("Given jobID, task name and no of documents and bind address", t, func() {
		ctx := context.Background()

		jobID := "task1234"

		taskToCreate := &TaskToCreate{
			TaskName:          "task",
			NumberOfDocuments: 3,
		}

		currentTime := time.Now().UTC()

		Convey("When NewTask is called", func() {
			task, err := NewTask(ctx, jobID, taskToCreate)
			So(err, ShouldBeNil)

			Convey("Then a new task resource is created", func() {
				So(task, ShouldNotBeEmpty)
				So(task.JobID, ShouldEqual, jobID)

				// check LastUpdated to a degree of seconds
				So(task.LastUpdated.Day(), ShouldEqual, currentTime.Day())
				So(task.LastUpdated.Month().String(), ShouldEqual, currentTime.Month().String())
				So(task.LastUpdated.Year(), ShouldEqual, currentTime.Year())
				So(task.LastUpdated.Hour(), ShouldEqual, currentTime.Hour())
				So(task.LastUpdated.Minute(), ShouldEqual, currentTime.Minute())
				So(task.LastUpdated.Second(), ShouldEqual, currentTime.Second())

				So(task.Links.Self, ShouldEqual, "/search-reindex-jobs/task1234/tasks/task")
				So(task.Links.Job, ShouldEqual, "/search-reindex-jobs/task1234")

				So(task.NumberOfDocuments, ShouldEqual, taskToCreate.NumberOfDocuments)
				So(task.TaskName, ShouldEqual, taskToCreate.TaskName)
			})
		})
	})
}

func TestValidate(t *testing.T) {
	Convey("Given a task to create with valid taskName", t, func() {
		task := TaskToCreate{
			TaskName:          "zebedee",
			NumberOfDocuments: 5,
		}

		Convey("When Validate is called", func() {
			err := task.Validate(taskNamesTest)

			Convey("Then no error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a task to create with empty taskName", t, func() {
		task := TaskToCreate{
			TaskName:          "",
			NumberOfDocuments: 5,
		}

		Convey("When Validate is called", func() {
			err := task.Validate(taskNamesTest)

			Convey("Then empty task name error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, apierrors.ErrEmptyTaskNameProvided)
			})
		})
	})

	Convey("Given a task to create with invalid taskName", t, func() {
		task := TaskToCreate{
			TaskName:          "invalid",
			NumberOfDocuments: 5,
		}

		Convey("When Validate is called", func() {
			err := task.Validate(taskNamesTest)

			Convey("Then invalid task name error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, apierrors.ErrTaskInvalidName)
			})
		})
	})
}
