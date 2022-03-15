package models

import (
	"time"

	dpresponse "github.com/ONSdigital/dp-net/v2/handlers/response"
	"github.com/globalsign/mgo/bson"
)

// GenerateETagForJob generates a new eTag for a job resource
func GenerateETagForJob(job Job) (eTag string, err error) {
	// ignoring the metadata LastUpdated and currentEtag when generating new eTag
	zeroTime := time.Time{}.UTC()
	job.ETag = ""
	job.LastUpdated = zeroTime

	b, err := bson.Marshal(job)
	if err != nil {
		return "", err
	}

	eTag = dpresponse.GenerateETag(b, false)

	return eTag, nil
}
