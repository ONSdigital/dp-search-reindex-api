package main

import (
	"context"
	"flag"
	"fmt"

	componentTest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-search-reindex-api/features/steps"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

// Mongo version here is overridden in the pipeline by the URL provided in the component.sh
const MongoVersion = "4.4.8"
const DatabaseName = "testing"

var componentFlag = flag.Bool("component", false, "perform component tests")

type ComponentTest struct {
	MongoFeature *componentTest.MongoFeature
	AuthFeature *componentTest.AuthorizationFeature
}

func (f *ComponentTest) InitializeScenario(godogCtx *godog.ScenarioContext) {
	ctx := context.Background()
	jobsFeature, err := steps.NewJobsFeature(f.MongoFeature, f.AuthFeature)
	if err != nil {
		log.Error(ctx, "error occurred while creating a new jobsFeature", err)
		os.Exit(1)
	}
	apiFeature := jobsFeature.InitAPIFeature()

	godogCtx.BeforeScenario(func(*godog.Scenario) {
		apiFeature.Reset()
		err := jobsFeature.Reset(false)
		if err != nil {
			log.Error(ctx, "error occurred while resetting the jobsFeature", err)
			os.Exit(1)
		}
	})
	godogCtx.AfterScenario(func(*godog.Scenario, error) {
		err := jobsFeature.Close()
		if err != nil {
			log.Error(ctx, "error occurred while closing the jobsFeature", err)
			os.Exit(1)
		}
	})
	jobsFeature.RegisterSteps(godogCtx)
	apiFeature.RegisterSteps(godogCtx)
	f.AuthFeature.RegisterSteps(godogCtx)
}
func (f *ComponentTest) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctxBackground := context.Background()
	ctx.BeforeSuite(func() {
		f.MongoFeature = componentTest.NewMongoFeature(componentTest.MongoOptions{MongoVersion: MongoVersion, DatabaseName: DatabaseName})
		f.AuthFeature = componentTest.NewAuthorizationFeature()
		f.AuthFeature.FakeAuthService.NewHandler().Get("/serviceInstancePermissions").Reply(200).BodyString(`{ "permissions": ["DELETE", "READ", "CREATE", "UPDATE"]}`)
	})
	ctx.AfterSuite(func() {
		err := f.MongoFeature.Close()
		if err != nil {
			log.Error(ctxBackground, "error occurred while closing the MongoFeature", err)
			os.Exit(1)
		}
		f.AuthFeature.Close()
	})
}
func TestComponent(t *testing.T) {
	if *componentFlag {
		status := 0
		var opts = godog.Options{
			Output: colors.Colored(os.Stdout),
			Format: "pretty",
			Paths:  flag.Args(),
		}
		f := &ComponentTest{}
		status = godog.TestSuite{
			Name:                 "feature_tests",
			ScenarioInitializer:  f.InitializeScenario,
			TestSuiteInitializer: f.InitializeTestSuite,
			Options:              &opts,
		}.Run()
		fmt.Println("=================================")
		fmt.Printf("Component test coverage: %.2f%%\n", testing.Coverage()*100)
		fmt.Println("=================================")
		if status != 0 {
			t.FailNow()
		}
	} else {
		t.Skip("component flag required to run component tests")
	}
}
