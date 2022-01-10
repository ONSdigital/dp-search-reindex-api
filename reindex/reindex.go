package reindex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/headers"
	kafka "github.com/ONSdigital/dp-kafka/v2"
	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/dp-search-reindex-api/config"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/schema"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/pkg/errors"
)

const serviceName = "dp-search-reindex-api"

// Reindex is a type that contains an implementation of the Indexer interface, which can be used for calling the Search API.
type Reindex struct {
}

type NewIndexName struct {
	IndexName string
}

// CreateIndex calls the Search API via the Do function of the dp-net/http/Clienter. It passes in the ServiceAuthToken to identify itself, as the Search Reindex API, to the Search API.
func (r *Reindex) CreateIndex(ctx context.Context, serviceAuthToken, searchAPISearchURL string, httpClient dphttp.Clienter) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, searchAPISearchURL, http.NoBody)
	if err != nil {
		return nil, errors.New("failed to create the request for post search")
	}
	if err := headers.SetServiceAuthToken(req, serviceAuthToken); err != nil {
		//TODO: ideally this needs to return an error when it couldn't set a service auth token.
		fmt.Println("error setting service auth token")
	}
	return httpClient.Do(ctx, req)
}

// GetIndexNameFromResponse unmarshalls the response body, which is passed into the function, and extracts the IndexName, which it then returns.
func (r *Reindex) GetIndexNameFromResponse(ctx context.Context, body io.ReadCloser) (string, error) {
	b, err := io.ReadAll(body)
	logData := log.Data{"response_body": string(b)}
	readBodyFailedMsg := "failed to read response body"
	unmarshalBodyFailedMsg := "failed to unmarshal response body"

	if err != nil {
		log.Error(ctx, readBodyFailedMsg, err, logData)
		return "", errors.New(readBodyFailedMsg)
	}

	if len(b) == 0 {
		b = []byte("[response body empty]")
	}

	var newIndexName NewIndexName

	if err = json.Unmarshal(b, &newIndexName); err != nil {
		log.Error(ctx, unmarshalBodyFailedMsg, err, logData)
		return "", errors.New(unmarshalBodyFailedMsg)
	}

	return newIndexName.IndexName, nil
}

func (r *Reindex) SendReindexRequestedEvent(cfg *config.Config, jobID string, indexName string) error {
	log.Namespace = serviceName
	ctx := context.Background()

	// Create Kafka Producer
	pChannels := kafka.CreateProducerChannels()
	kafkaProducer, err := kafka.NewProducer(ctx, cfg.KafkaConfig.Brokers, cfg.KafkaConfig.ReindexRequestedTopic, pChannels, &kafka.ProducerConfig{
		KafkaVersion: &cfg.KafkaConfig.Version,
	})
	if err != nil {
		return errors.New("failed to create kafka producer")
	}

	// kafka error logging go-routines
	kafkaProducer.Channels().LogErrors(ctx, "kafka producer")

	time.Sleep(500 * time.Millisecond)

	traceID := request.NewRequestID(16)
	reindexReqEvent := models.ReindexRequested{
		JobID:       jobID,
		SearchIndex: indexName,
		TraceID:     traceID,
	}

	log.Info(ctx, "sending reindex-requested event", log.Data{"reindexRequestedEvent": reindexReqEvent})

	bytes, err := schema.ReindexRequestedEvent.Marshal(reindexReqEvent)
	if err != nil {
		return errors.New("reindex-requested event error")
	}

	// Send bytes to Output channel, after calling Initialise just in case it is not initialised.
	kafkaProducer.Initialise(ctx)
	kafkaProducer.Channels().Output <- bytes

	return err
}
