package api_test

import (
	"context"
	"testing"

	"github.com/ONSdigital/dp-search-reindex-api/api"
	"github.com/ONSdigital/dp-search-reindex-api/api/mock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
)

func TestSetup(t *testing.T) {

	Convey("Given an API instance", t, func() {

		api := api.Setup(context.Background(), mux.NewRouter(), &mock.JobStorerMock{}, &mock.AuthHandlerMock{})

		Convey("When created the following routes should have been added", func() {
			So(hasRoute(api.Router, "/jobs", "POST"), ShouldBeTrue)
			So(hasRoute(api.Router, "/jobs/{id}", "GET"), ShouldBeTrue)
			So(hasRoute(api.Router, "/jobs", "GET"), ShouldBeTrue)
		})
	})
}

func TestClose(t *testing.T) {

	Convey("Given an API instance", t, func() {
		ctx := context.Background()
		api := api.Setup(context.Background(), mux.NewRouter(), &mock.JobStorerMock{}, &mock.AuthHandlerMock{})

		Convey("When the api is closed then there is no error returned", func() {
			err := api.Close(ctx)
			So(err, ShouldBeNil)
		})
	})
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}
