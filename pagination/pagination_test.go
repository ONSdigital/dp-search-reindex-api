package pagination_test

import (
	"testing"

	"github.com/ONSdigital/dp-search-reindex-api/pagination"
	. "github.com/smartystreets/goconvey/convey"
)

// Constants for testing
const (
	defaultLimit    = 20
	defaultOffset   = 0
	defaultMaxLimit = 1000
)

func TestValidatePaginationParametersReturnsErrorWhenOffsetIsNegative(t *testing.T) {

	Convey("Given a minus offset value and a JobStore containing 20 jobs", t, func() {

		offset := "-1"
		limit := ""
		totalCount := 20

		Convey("When ValidatePaginationValues is called", func() {

			paginator := pagination.NewPaginator(defaultLimit, defaultOffset, defaultMaxLimit)
			offset, limit, err := paginator.ValidatePaginationParameters(offset, limit, totalCount)

			Convey("Then the expected error is returned", func() {

				So(err, ShouldEqual, pagination.ErrInvalidOffsetParameter)
				So(offset, ShouldBeZeroValue)
				So(limit, ShouldBeZeroValue)
			})
		})
	})
}

func TestValidatePaginationParametersReturnsErrorWhenLimitIsNegative(t *testing.T) {

	Convey("Given a minus limit value and a JobStore containing 20 jobs", t, func() {

		offset := ""
		limit := "-1"
		totalCount := 20

		Convey("When ValidatePaginationValues is called", func() {

			paginator := pagination.NewPaginator(defaultLimit, defaultOffset, defaultMaxLimit)
			offset, limit, err := paginator.ValidatePaginationParameters(offset, limit, totalCount)

			Convey("Then the expected error is returned", func() {
				So(err, ShouldEqual, pagination.ErrInvalidLimitParameter)
				So(offset, ShouldBeZeroValue)
				So(limit, ShouldBeZeroValue)
			})
		})
	})
}

func TestValidatePaginationParametersReturnsErrorWhenLimitIsGreaterThanMaxLimit(t *testing.T) {

	Convey("Given a request with a limit value over the maximum", t, func() {

		offset := ""
		limit := "1001"
		totalCount := 20

		Convey("When ValidatePaginationValues is called", func() {

			paginator := pagination.NewPaginator(defaultLimit, defaultOffset, defaultMaxLimit)
			offset, limit, err := paginator.ValidatePaginationParameters(offset, limit, totalCount)

			Convey("Then the expected error is returned", func() {
				So(err, ShouldEqual, pagination.ErrLimitOverMax)
				So(offset, ShouldBeZeroValue)
				So(limit, ShouldBeZeroValue)
			})
		})
	})
}

func TestValidatePaginationParametersReturnsLimitAndOffsetProvidedFromQuery(t *testing.T) {

	Convey("Given a request with a valid limit and offset", t, func() {

		offset := "5"
		limit := "10"
		totalCount := 20

		Convey("When ValidatePaginationValues is called", func() {

			paginator := pagination.NewPaginator(defaultLimit, defaultOffset, defaultMaxLimit)
			offset, limit, err := paginator.ValidatePaginationParameters(offset, limit, totalCount)

			Convey("Then the expected values are returned", func() {
				So(err, ShouldBeNil)
				So(offset, ShouldEqual, 5)
				So(limit, ShouldEqual, 10)
			})
		})
	})
}

func TestValidatePaginationParametersReturnsDefaultValuesWhenNotProvided(t *testing.T) {

	Convey("Given a request without pagination parameters", t, func() {

		offset := ""
		limit := ""
		totalCount := 20

		Convey("When ValidatePaginationValues is called", func() {
			expectedLimit := 15
			expectedOffset := 1
			paginator := pagination.NewPaginator(expectedLimit, expectedOffset, defaultMaxLimit)
			offset, limit, err := paginator.ValidatePaginationParameters(offset, limit, totalCount)

			Convey("Then the configured default values are returned", func() {
				So(err, ShouldBeNil)
				So(offset, ShouldEqual, expectedOffset)
				So(limit, ShouldEqual, expectedLimit)
			})
		})
	})
}

func TestNewPaginatorReturnsPaginatorStructWithFilledValues(t *testing.T) {

	Convey("Given a set of expected paginator values", t, func() {

		expectedPaginator := &pagination.Paginator{
			DefaultLimit:    10,
			DefaultOffset:   5,
			DefaultMaxLimit: 100,
		}

		Convey("When NewPaginator is called", func() {

			actualPaginator := pagination.NewPaginator(10, 5, 100)

			Convey("Then the paginator is configured as expected", func() {
				So(actualPaginator, ShouldResemble, expectedPaginator)
			})
		})
	})
}
