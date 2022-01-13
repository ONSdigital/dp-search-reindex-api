package event

import (
	"context"
	"fmt"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	kafka "github.com/ONSdigital/dp-kafka/v2"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/log.go/v2/log"
)

//go:generate moq -out ./mock/producer.go -pkg mock . Marshaller

// Marshaller defines a type for marshalling the requested object into the required format.
type Marshaller interface {
	Marshal(s interface{}) ([]byte, error)
}

// ReindexRequestedProducer produces kafka messages to send to the reindex-requested topic.
type ReindexRequestedProducer struct {
	Marshaller Marshaller
	Producer   kafka.IProducer
}

// ProduceReindexRequested produce a kafka message for an instance which has been successfully processed.
func (p ReindexRequestedProducer) ProduceReindexRequested(ctx context.Context, event models.ReindexRequested) error {
	bytes, err := p.Marshaller.Marshal(event)
	if err != nil {
		log.Fatal(ctx, "Marshaller.Marshal", err)
		return fmt.Errorf(fmt.Sprintf("Marshaller.Marshal returned an error: event=%v: %%w", event), err)
	}
	p.Producer.Channels().Output <- bytes
	log.Info(ctx, "completed successfully", log.Data{"event": event, "package": "event.ReindexRequestedProducer"})
	return nil
}

// Close is called when the service shuts down gracefully
func (p ReindexRequestedProducer) Close(ctx context.Context) error {
	return p.Producer.Close(ctx)
}

// Checker is called by the healthcheck library to check the health state of this kafka producer instance
func (p *ReindexRequestedProducer) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	return p.Producer.Checker(ctx, state)
}
