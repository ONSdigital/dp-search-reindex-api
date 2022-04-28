package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/mongo"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"

	dpresponse "github.com/ONSdigital/dp-net/v2/handlers/response"
)

var invalidBodyErrorMessage = "invalid request body"

const v1 = "v1"

// CreateTaskHandler returns a function that generates a new TaskName resource containing default values in its fields.
func (api *API) CreateTaskHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	host := req.Host
	vars := mux.Vars(req)
	jobID := vars["id"]

	// Unmarshal task to create and validate it
	taskToCreate := &models.TaskToCreate{}
	if err := ReadJSONBody(req.Body, taskToCreate); err != nil {
		log.Error(ctx, "reading request body failed", err)
		http.Error(w, invalidBodyErrorMessage, http.StatusBadRequest)
		return
	}

	if err := taskToCreate.Validate(api.taskNames); err != nil {
		log.Error(ctx, "CreateTask endpoint: Invalid request body", err)
		http.Error(w, invalidBodyErrorMessage, http.StatusBadRequest)
		return
	}

	newTask, err := api.dataStore.CreateTask(ctx, jobID, taskToCreate.TaskName, taskToCreate.NumberOfDocuments)
	if err != nil {
		log.Error(ctx, "creating and storing a task failed", err, log.Data{"job id": jobID})
		if err == mongo.ErrJobNotFound {
			http.Error(w, "Failed to find job that has the specified id", http.StatusNotFound)
		} else {
			http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		}
		return
	}

	newTask.Links.Job = fmt.Sprintf("%s/%s%s", host, v1, newTask.Links.Job)
	newTask.Links.Self = fmt.Sprintf("%s/%s%s", host, v1, newTask.Links.Self)

	jsonResponse, err := json.Marshal(newTask)
	if err != nil {
		log.Error(ctx, "marshalling response failed", err)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}

	// set eTag on ETag response header
	dpresponse.SetETag(w, newTask.ETag)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error(ctx, "writing response failed", err)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}

// GetTaskHandler returns a function that gets a specific task, associated with an existing Job resource, using the job id and task name passed in.
func (api *API) GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	host := req.Host
	vars := mux.Vars(req)
	id := vars["id"]
	taskName := vars["task_name"]
	logData := log.Data{"job_id": id, "task_name": taskName}

	task, err := api.dataStore.GetTask(ctx, id, taskName)
	if err != nil {
		log.Error(ctx, "getting task failed", err, logData)
		if err == mongo.ErrJobNotFound {
			http.Error(w, "failed to find task - job id is invalid", http.StatusNotFound)
		} else if err == mongo.ErrTaskNotFound {
			http.Error(w, "failed to find task for the specified job id", http.StatusNotFound)
		} else {
			http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		}
		return
	}

	task.Links.Job = fmt.Sprintf("%s/%s%s", host, v1, task.Links.Job)
	task.Links.Self = fmt.Sprintf("%s/%s%s", host, v1, task.Links.Self)

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(task)
	if err != nil {
		log.Error(ctx, "marshalling response failed", err, logData)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error(ctx, "writing response failed", err, logData)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}

// GetTasksHandler gets a list of existing Task resources, from the data store, sorted by their values of
// last_updated time (ascending).
func (api *API) GetTasksHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	host := req.Host
	offsetParam := req.URL.Query().Get("offset")
	limitParam := req.URL.Query().Get("limit")
	vars := mux.Vars(req)
	id := vars["id"]
	logData := log.Data{"job_id": id}

	offset, limit, err := api.setUpPagination(offsetParam, limitParam)
	if err != nil {
		log.Error(ctx, "pagination validation failed", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := api.dataStore.GetTasks(ctx, offset, limit, id)
	if err != nil {
		log.Error(ctx, "getting tasks failed", err, logData)
		switch {
		case err == mongo.ErrJobNotFound:
			http.Error(w, "failed to find tasks - job id is invalid", http.StatusNotFound)
			return
		default:
			log.Error(ctx, "getting list of tasks failed", err)
			http.Error(w, serverErrorMessage, http.StatusInternalServerError)
			return
		}
	}

	for i := range tasks.TaskList {
		tasks.TaskList[i].Links.Job = fmt.Sprintf("%s/%s%s", host, v1, tasks.TaskList[i].Links.Job)
		tasks.TaskList[i].Links.Self = fmt.Sprintf("%s/%s%s", host, v1, tasks.TaskList[i].Links.Self)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(tasks)
	if err != nil {
		log.Error(ctx, "marshalling response failed", err)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error(ctx, "writing response failed", err)
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}
