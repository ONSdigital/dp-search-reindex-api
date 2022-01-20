package event_test

import (
	"context"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	kafka "github.com/ONSdigital/dp-kafka/v2"
	"github.com/ONSdigital/dp-search-reindex-api/event"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestProducer(t *testing.T) {

	fakeKafkaProducer := &FakeKafkaProducer{}
	reindexRequestedEvent := models.ReindexRequested{}
	checkState := healthcheck.NewCheckState("test")
	fakeMarshaller := FakeMarshaller{}

	Convey("Given the producer is initialized without a marshaller", t, func() {
		sut := event.ReindexRequestedProducer{Marshaller: nil, Producer: fakeKafkaProducer}

		Convey("When ProduceReindexRequested is called then an error should be returned", func() {
			err := sut.ProduceReindexRequested(context.Background(), reindexRequestedEvent)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "marshaller is not provided")
		})

		Convey("When Close is called then an error should be called", func() {
			err := sut.Close(context.Background())
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "marshaller is not provided")
		})

		Convey("When Checker is called then an error should be called", func() {
			err := sut.Checker(context.Background(), checkState)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "marshaller is not provided")
		})
	})

	Convey("Given the producer is initialized without a kafka provider", t, func() {
		producer := event.ReindexRequestedProducer{Marshaller: fakeMarshaller, Producer: nil}
		const expectedErrorMessage = "producer is not provided"

		Convey("When ProduceReindexRequested is called then an error should be returned", func() {
			err := producer.ProduceReindexRequested(context.Background(), reindexRequestedEvent)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expectedErrorMessage)
		})

		Convey("When Close is called then an error should be called", func() {
			err := producer.Close(context.Background())
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expectedErrorMessage)
		})

		Convey("When Checker is called then an error should be called", func() {
			err := producer.Checker(context.Background(), checkState)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, expectedErrorMessage)
		})
	})

	Convey("Given an initialized producer", t, func() {
		sut := event.ReindexRequestedProducer{Marshaller: fakeMarshaller, Producer: fakeKafkaProducer}

		Convey("When Close is called, Then the kafka producer should be closed", func() {
			sut.Close(context.Background())
			So(fakeKafkaProducer.CloseCalls, ShouldHaveLength, 1)
			So(fakeKafkaProducer.CloseCalls[0], ShouldEqual, context.Background())

		})

		Convey("When Checker is called, Then the kafka producer checker should be called", func() {
			sut.Checker(context.Background(), checkState)
			So(fakeKafkaProducer.CheckerCalls, ShouldHaveLength, 1)
			So(fakeKafkaProducer.CheckerCalls[0], ShouldResemble,
				checkerCall{context.Background(), checkState})

		})
	})

}

type FakeMarshaller struct {
}

func (FakeMarshaller) Marshal(s interface{}) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

type checkerCall struct {
	Context    context.Context
	CheckState *healthcheck.CheckState
}
type FakeKafkaProducer struct {
	CloseCalls   []context.Context
	CheckerCalls []checkerCall
}

func (f *FakeKafkaProducer) Channels() *kafka.ProducerChannels {
	//TODO implement me
	panic("implement me")
}

func (f *FakeKafkaProducer) IsInitialised() bool {
	//TODO implement me
	panic("implement me")
}

func (f *FakeKafkaProducer) Initialise(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f *FakeKafkaProducer) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	//TODO implement me
	f.CheckerCalls = append(f.CheckerCalls, checkerCall{ctx, state})
	return nil
}

func (f *FakeKafkaProducer) Close(ctx context.Context) (err error) {
	f.CloseCalls = append(f.CloseCalls, ctx)
	return nil
}
