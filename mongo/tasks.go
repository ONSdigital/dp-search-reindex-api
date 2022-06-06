package mongo

import (
	"context"
	"time"

	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// GetTask retrieves the details of a particular task, from the collection, specified by its task name and associated job id
func (m *JobStore) GetTask(ctx context.Context, jobID, taskName string) (*models.Task, error) {
	logData := log.Data{
		"jobID":    jobID,
		"taskName": taskName,
	}

	log.Info(ctx, "getting task from the data store", logData)

	s := m.Session.Copy()
	defer s.Close()

	var task models.Task

	// find task in mongo using job_id and task_name
	err := s.DB(m.Database).C(m.TasksCollection).Find(bson.M{"job_id": jobID, "task_name": taskName}).One(&task)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Error(ctx, "task not found in mongo", err, logData)
			return nil, ErrTaskNotFound
		}

		log.Error(ctx, "error occurred when finding task in mongo", err, logData)
		return nil, err
	}

	return &task, nil
}

// GetTasks retrieves all the tasks, from the collection, and lists them in order of last_updated
func (m *JobStore) GetTasks(ctx context.Context, jobID string, options Options) (*models.Tasks, error) {
	logData := log.Data{
		"job_id":  jobID,
		"options": options,
	}

	log.Info(ctx, "getting list of tasks", logData)

	// get the count of tasks related to jobID
	numTasks, err := m.getTasksCount(ctx, jobID)
	if err != nil {
		log.Error(ctx, "failed to get tasks count", err, logData)
		return nil, err
	}

	s := m.Session.Copy()
	defer s.Close()

	// get tasks from the tasks collection using the given job_id, offset and limit, and order them by last_updated
	tasksQuery := s.DB(m.Database).C(m.TasksCollection).Find(bson.M{"job_id": jobID}).Skip(options.Offset).Limit(options.Limit).Sort("last_updated")

	// populate taskList using tasksQuery
	taskList := make([]models.Task, numTasks)
	if err := tasksQuery.All(&taskList); err != nil {
		log.Error(ctx, "failed to populate task list", err, logData)
		return nil, err
	}

	tasks := &models.Tasks{
		Count:      len(taskList),
		TaskList:   taskList,
		Limit:      options.Limit,
		Offset:     options.Offset,
		TotalCount: numTasks,
	}

	return tasks, nil
}

// getTasksCount returns the total number of tasks stored in the tasks collection in mongo
func (m *JobStore) getTasksCount(ctx context.Context, jobID string) (int, error) {
	logData := log.Data{}

	s := m.Session.Copy()
	defer s.Close()

	// find no of tasks related to job_id
	numTasks, err := s.DB(m.Database).C(m.TasksCollection).Find(bson.M{"job_id": jobID}).Count()
	if err != nil {
		logData["database"] = m.Database
		logData["tasks_collections"] = m.TasksCollection

		log.Error(ctx, "failed to count tasks for given job id", err, logData)
		return 0, err
	}

	if numTasks == 0 {
		log.Info(ctx, "there are no tasks in the data store")
		return 0, nil
	}

	logData["no_of_tasks"] = numTasks
	log.Info(ctx, "number of tasks found in tasks collection", logData)

	return numTasks, nil
}

// UpsertTask creates a new task document or overwrites an existing one
func (m *JobStore) UpsertTask(ctx context.Context, jobID, taskName string, task models.Task) error {
	log.Info(ctx, "upserting task to mongo")

	s := m.Session.Copy()
	defer s.Close()

	selector := bson.M{
		"task_name": taskName,
		"job_id":    jobID,
	}

	task.LastUpdated = time.Now().UTC()

	update := bson.M{
		"$set": task,
	}

	_, err := s.DB(m.Database).C(m.TasksCollection).Upsert(selector, update)
	if err != nil {
		logData := log.Data{
			"job_id":    jobID,
			"task_name": taskName,
			"task":      task,
		}
		log.Error(ctx, "failed to upsert task in mongo", err, logData)
		return err
	}

	return nil
}
