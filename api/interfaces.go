package api

import (
	"context"

	"github.com/ONSdigital/dp-search-reindex-api/models"
)

//go:generate moq -out mock/job_storer.go -pkg mock . JobStorer

// JobStorer is an interface for a type that can store and retrieve jobs
type JobStorer interface {
	CreateJob(ctx context.Context, id string) (job models.Job, err error)
	GetJob(ctx context.Context, id string) (job models.Job, err error)
	GetJobs(ctx context.Context) (job models.Jobs, err error)
	AcquireJobLock(ctx context.Context, id string) (lockID string, err error)
	UnlockJob(lockID string) error
	PutNumberOfTasks(ctx context.Context, id string, count int) error
}
