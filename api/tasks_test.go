package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	dpHTTP "github.com/ONSdigital/dp-net/v2/http"
	"github.com/ONSdigital/dp-search-reindex-api/api"
	apiMock "github.com/ONSdigital/dp-search-reindex-api/api/mock"
	"github.com/ONSdigital/dp-search-reindex-api/apierrors"
	"github.com/ONSdigital/dp-search-reindex-api/config"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/mongo"
	"github.com/ONSdigital/dp-search-reindex-api/url"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// Constants for testing
const (
	invalidJobID          = "UUID3"
	emptyTaskName         = ""
	validTaskName1        = "zebedee"
	validTaskName2        = "dataset-api"
	invalidTaskName       = "any-word-not-in-valid-list"
	validServiceAuthToken = "Bearer fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"
	bindAddress           = "localhost:25700"
	testNoOfDocs          = 5
	validOffset           = 0
	validLimit            = 10
)

// Create Task Payload
var createTaskPayloadFmt = `{
	"task_name": "%s",
	"number_of_documents": 5
}`

func TestCreateTaskHandler(t *testing.T) {
	dataStorerMock := &apiMock.DataStorerMock{
		CreateTaskFunc: func(ctx context.Context, jobID string, taskName string, numDocuments int) (models.Task, error) {
			emptyTask := models.Task{}

			switch taskName {
			case emptyTaskName:
				return emptyTask, apierrors.ErrEmptyTaskNameProvided
			case invalidTaskName:
				return emptyTask, apierrors.ErrTaskInvalidName
			}

			switch jobID {
			case validJobID1:
				return models.NewTask(jobID, taskName, numDocuments, bindAddress), nil
			case invalidJobID:
				return emptyTask, mongo.ErrJobNotFound
			default:
				return emptyTask, errors.New("an unexpected error occurred")
			}
		},
	}

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("When a new reindex task is created and stored", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks", validJobID1), bytes.NewBufferString(
				fmt.Sprintf(createTaskPayloadFmt, validTaskName1)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", validServiceAuthToken)
			resp := httptest.NewRecorder()

			apiInstance.Router.ServeHTTP(resp, req)

			Convey("Then the newly created search reindex task is returned with status code 201", func() {
				So(resp.Code, ShouldEqual, http.StatusCreated)
				payload, err := io.ReadAll(resp.Body)
				So(err, ShouldBeNil)
				newTask := models.Task{}
				err = json.Unmarshal(payload, &newTask)
				So(err, ShouldBeNil)
				zeroTime := time.Time{}.UTC()
				expectedTask, err := expectedTask(validJobID1, zeroTime, 5, validTaskName1)
				So(err, ShouldBeNil)

				Convey("And the new task resource should contain expected values", func() {
					So(newTask.JobID, ShouldEqual, expectedTask.JobID)
					So(newTask.Links, ShouldResemble, expectedTask.Links)
					So(newTask.NumberOfDocuments, ShouldEqual, expectedTask.NumberOfDocuments)
					So(newTask.TaskName, ShouldEqual, expectedTask.TaskName)
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And job ID is invalid", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks", invalidJobID), bytes.NewBufferString(
				fmt.Sprintf(createTaskPayloadFmt, validTaskName2)))

			Convey("When the tasks endpoint is called to create and store a new reindex task", func() {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", validServiceAuthToken)
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then an empty search reindex job is returned with status code 404", func() {
					So(resp.Code, ShouldEqual, http.StatusNotFound)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, "Failed to find job that has the specified id")
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an error occurred for creating and storing a task", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks", "invalid"), bytes.NewBufferString(
				fmt.Sprintf(createTaskPayloadFmt, validTaskName2)))

			Convey("When the tasks endpoint is called to create and store a new reindex task", func() {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", validServiceAuthToken)
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then an empty search reindex job is returned with status code 404", func() {
					So(resp.Code, ShouldEqual, http.StatusInternalServerError)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, apierrors.ErrInternalServer.Error())
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("When the tasks endpoint is called to create and store a new reindex task", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks", validJobID1), bytes.NewBufferString(
				fmt.Sprintf(createTaskPayloadFmt, emptyTaskName)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", validServiceAuthToken)
			resp := httptest.NewRecorder()

			apiInstance.Router.ServeHTTP(resp, req)

			Convey("Then an empty search reindex job is returned with status code 400 because the task name is empty", func() {
				So(resp.Code, ShouldEqual, http.StatusBadRequest)
				errMsg := strings.TrimSpace(resp.Body.String())
				So(errMsg, ShouldEqual, "invalid request body")
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("When the tasks endpoint is called to create and store a new reindex task", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks", validJobID1), bytes.NewBufferString(
				fmt.Sprintf(createTaskPayloadFmt, invalidTaskName)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", validServiceAuthToken)
			resp := httptest.NewRecorder()

			apiInstance.Router.ServeHTTP(resp, req)

			Convey("Then an empty search reindex job is returned with status code 400 because the task name is invalid", func() {
				So(resp.Code, ShouldEqual, http.StatusBadRequest)
				errMsg := strings.TrimSpace(resp.Body.String())
				So(errMsg, ShouldEqual, "invalid request body")
			})
		})
	})
}

func TestGetTaskHandler(t *testing.T) {
	dataStorerMock := &apiMock.DataStorerMock{
		GetTaskFunc: func(ctx context.Context, jobID, taskName string) (models.Task, error) {
			emptyTask := models.Task{}

			switch taskName {
			case emptyTaskName:
				return emptyTask, apierrors.ErrEmptyTaskNameProvided
			case invalidTaskName:
				return emptyTask, mongo.ErrTaskNotFound
			}

			switch jobID {
			case validJobID1:
				return models.NewTask(jobID, taskName, testNoOfDocs, bindAddress), nil
			case invalidJobID:
				return emptyTask, mongo.ErrJobNotFound
			default:
				return emptyTask, errors.New("an unexpected error occurred")
			}
		},
	}

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And valid job id and task name is given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks/%s", validJobID1, validTaskName1), nil)

			Convey("When GetTaskHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 200", func() {
					So(resp.Code, ShouldEqual, http.StatusOK)
					payload, err := io.ReadAll(resp.Body)
					So(err, ShouldBeNil)
					task := models.Task{}
					err = json.Unmarshal(payload, &task)
					So(err, ShouldBeNil)

					Convey("And the task resource should contain expected values", func() {
						zeroTime := time.Time{}.UTC()
						expectedTask, err := expectedTask(validJobID1, zeroTime, testNoOfDocs, validTaskName1)
						So(err, ShouldBeNil)

						So(task.JobID, ShouldEqual, expectedTask.JobID)
						So(task.Links, ShouldResemble, expectedTask.Links)
						So(task.NumberOfDocuments, ShouldEqual, expectedTask.NumberOfDocuments)
						So(task.TaskName, ShouldEqual, expectedTask.TaskName)
					})
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an invalid job id is given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks/%s", invalidJobID, validTaskName1), nil)

			Convey("When GetTaskHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 200", func() {
					So(resp.Code, ShouldEqual, http.StatusNotFound)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, "failed to find task - job id is invalid")
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an invalid task name is given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks/%s", validJobID1, invalidTaskName), nil)

			Convey("When GetTaskHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 200", func() {
					So(resp.Code, ShouldEqual, http.StatusNotFound)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, "failed to find task for the specified task name")
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an unexpected error occurs when getting task", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks/%s", "invalid", validTaskName1), nil)

			Convey("When GetTaskHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 200", func() {
					So(resp.Code, ShouldEqual, http.StatusInternalServerError)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, apierrors.ErrInternalServer.Error())
				})
			})
		})
	})
}

func TestGetTasksHandler(t *testing.T) {
	dataStorerMock := &apiMock.DataStorerMock{
		GetTaskFunc: func(ctx context.Context, jobID, taskName string) (models.Task, error) {
			emptyTask := models.Task{}
			switch jobID {
			case validJobID1:
				return models.NewTask(jobID, taskName, testNoOfDocs, bindAddress), nil
			case invalidJobID:
				return emptyTask, mongo.ErrJobNotFound
			default:
				return emptyTask, errors.New("an unexpected error occurred")
			}
		},
		GetTasksFunc: func(ctx context.Context, offset, limit int, jobID string) (models.Tasks, error) {
			switch jobID {
			case validJobID1:
				return expectedTasks(jobID), nil
			case invalidJobID:
				return models.Tasks{}, mongo.ErrJobNotFound
			default:
				return models.Tasks{}, apierrors.ErrInternalServer
			}
		},
	}

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And valid job id, offset and limit is given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks?offset=%d&limit=%d", validJobID1, validOffset, validLimit), nil)

			Convey("When GetTasksHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the tasks should be retrieved with status code 200", func() {
					So(resp.Code, ShouldEqual, http.StatusOK)
					payload, err := io.ReadAll(resp.Body)
					So(err, ShouldBeNil)
					tasks := models.Tasks{}
					err = json.Unmarshal(payload, &tasks)
					So(err, ShouldBeNil)

					Convey("And the task resource should contain expected values", func() {
						expectedTask := expectedTasks(validJobID1)
						So(err, ShouldBeNil)

						So(tasks.TaskList[0].JobID, ShouldEqual, expectedTask.TaskList[0].JobID)
						So(tasks.TaskList[0].Links, ShouldResemble, expectedTask.TaskList[0].Links)
						So(tasks.TaskList[0].NumberOfDocuments, ShouldEqual, expectedTask.TaskList[0].NumberOfDocuments)
						So(tasks.TaskList[0].TaskName, ShouldEqual, expectedTask.TaskList[0].TaskName)

						So(tasks.Count, ShouldEqual, expectedTask.Count)
						So(tasks.Limit, ShouldEqual, expectedTask.Limit)
						So(tasks.Offset, ShouldEqual, expectedTask.Offset)
						So(tasks.TotalCount, ShouldEqual, expectedTask.TotalCount)
					})
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an invalid pagination values are given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks?offset=%d&limit=invalid", validJobID1, validOffset), nil)

			Convey("When GetTasksHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 400", func() {
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldNotBeEmpty)
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an invalid job id is given", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks?offset=%d&limit=%d", invalidJobID, validOffset, validOffset), nil)

			Convey("When GetTasksHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 404", func() {
					So(resp.Code, ShouldEqual, http.StatusNotFound)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, "failed to find tasks - job id is invalid")
				})
			})
		})
	})

	Convey("Given an API that can create valid search reindex tasks and store their details in a Data Store", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)
		httpClient := dpHTTP.NewClient()
		apiInstance := api.Setup(mux.NewRouter(), dataStorerMock, &apiMock.AuthHandlerMock{}, taskNames, cfg, httpClient, &apiMock.IndexerMock{}, &apiMock.ReindexRequestedProducerMock{})

		Convey("And an error occurred when getting list of tasks", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:25700/jobs/%s/tasks?offset=%d&limit=%d", "invalid", validOffset, validOffset), nil)

			Convey("When GetTasksHandler is called", func() {
				resp := httptest.NewRecorder()

				apiInstance.Router.ServeHTTP(resp, req)

				Convey("Then the search reindex task should be retrieved with status code 404", func() {
					So(resp.Code, ShouldEqual, http.StatusInternalServerError)
					errMsg := strings.TrimSpace(resp.Body.String())
					So(errMsg, ShouldEqual, apierrors.ErrInternalServer.Error())
				})
			})
		})
	})
}

func expectedTask(jobID string, lastUpdated time.Time, numberOfDocuments int, taskName string) (models.Task, error) {
	cfg, err := config.Get()
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to retrieve service configuration: %w", err)
	}
	urlBuilder := url.NewBuilder("http://" + cfg.BindAddr)
	job := urlBuilder.BuildJobURL(jobID)
	self := urlBuilder.BuildJobTaskURL(jobID, taskName)
	return models.Task{
		JobID:       jobID,
		LastUpdated: lastUpdated,
		Links: &models.TaskLinks{
			Self: self,
			Job:  job,
		},
		NumberOfDocuments: numberOfDocuments,
		TaskName:          taskName,
	}, nil
}

func expectedTasks(jobID string) models.Tasks {
	validTask := models.NewTask(jobID, validTaskName1, testNoOfDocs, bindAddress)
	return models.Tasks{
		Count:      1,
		TaskList:   []models.Task{validTask},
		Limit:      validLimit,
		Offset:     validOffset,
		TotalCount: 1,
	}
}
