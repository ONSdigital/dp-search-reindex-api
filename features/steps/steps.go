package steps

import "github.com/cucumber/godog"

// RegisterSteps defines the steps within a specific SearchReindexAPIFeature cucumber test.
func (f *SearchReindexAPIFeature) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^set the api version to ([^"]*) for incoming requests$`, f.setAPIVersionForPath)
	ctx.Step(`^a new task resource is created containing the following values:$`, f.aNewTaskResourceIsCreatedContainingTheFollowingValues)
	ctx.Step(`^each job should also contain the following values:$`, f.eachJobShouldAlsoContainTheFollowingValues)
	ctx.Step(`^each task should also contain the following values:$`, f.eachTaskShouldAlsoContainTheFollowingValues)
	ctx.Step(`^I am not identified by zebedee$`, f.iAmNotIdentifiedByZebedee)
	ctx.Step(`^the search api is working correctly$`, f.successfulSearchAPIResponse)
	ctx.Step(`^the search api is not working correctly$`, f.unsuccessfulSearchAPIResponse)
	ctx.Step(`^restart the search api$`, f.restartFakeSearchAPI)

	ctx.Step(`^I call GET \/jobs\/{id} using the generated id$`, f.iCallGETJobsidUsingTheGeneratedID)
	ctx.Step(`^I call GET \/jobs\/{"([^"]*)"} using a valid UUID$`, f.iCallGETJobsUsingAValidUUID)
	ctx.Step(`^I call GET \/jobs\/{id}\/tasks\/{"([^"]*)"}$`, f.iCallGETJobsidtasks)
	ctx.Step(`^I call GET \/jobs\/{"([^"]*)"}\/tasks\/{"([^"]*)"} using a valid UUID$`, f.iCallGETJobsTasksUsingAValidUUID)
	ctx.Step(`^I call GET \/jobs\/{id}\/tasks using the same id again$`, f.iCallGETJobsidtasksUsingTheSameIDAgain)
	ctx.Step(`^I call GET \/jobs\/{id}\/tasks\?offset="([^"]*)"&limit="([^"]*)"$`, f.iCallGETJobsidtasksoffsetLimit)
	ctx.Step(`^I GET "\/jobs\/{"([^"]*)"}\/tasks"$`, f.iGETJobsTasks)
	ctx.Step(`^I GET \/jobs\/{id}\/tasks using the generated id$`, f.iGETJobsidtasksUsingTheGeneratedID)

	ctx.Step(`^I call POST \/jobs\/{id}\/tasks to update the number_of_documents for that task$`, f.iCallPOSTJobsidtasksToUpdateTheNumberofdocumentsForThatTask)
	ctx.Step(`^I call POST \/jobs\/{id}\/tasks using the generated id$`, f.iCallPOSTJobsidtasksUsingTheGeneratedID)
	ctx.Step(`^I call POST \/jobs\/{id}\/tasks using the same id again$`, f.iCallPOSTJobsidtasksUsingTheSameIDAgain)

	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{(\d+)} using the generated id$`, f.iCallPUTJobsidnumberTofTasksUsingTheGeneratedID)
	ctx.Step(`^I call PUT \/jobs\/{"([^"]*)"}\/number_of_tasks\/{(\d+)} using a valid UUID$`, f.iCallPUTJobsNumberoftasksUsingAValidUUID)
	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{"([^"]*)"} using the generated id with an invalid count$`, f.iCallPUTJobsidnumberoftasksUsingTheGeneratedIDWithAnInvalidCount)
	ctx.Step(`^I call PUT \/jobs\/{id}\/number_of_tasks\/{"([^"]*)"} using the generated id with a negative count$`, f.iCallPUTJobsidnumberoftasksUsingTheGeneratedIDWithANegativeCount)

	ctx.Step(`^I call PATCH \/jobs\/{id} using the generated id$`, f.iCallPATCHJobsIDUsingTheGeneratedID)

	ctx.Step(`^I have created a task for the generated job$`, f.iHaveCreatedATaskForTheGeneratedJob)
	ctx.Step(`^I have generated (\d+) jobs in the Job Store$`, f.iHaveGeneratedJobsInTheJobStore)
	ctx.Step(`^I set the If-Match header to the generated e-tag$`, f.iSetIfMatchHeaderToTheGeneratedETag)
	ctx.Step(`^I set the "If-Match" header to the old e-tag$`, f.iSetIfMatchHeaderToTheOldGeneratedETag)

	ctx.Step(`^I would expect the response to be an empty list$`, f.iWouldExpectTheResponseToBeAnEmptyList)
	ctx.Step(`^I would expect the response to be an empty list of tasks$`, f.iWouldExpectTheResponseToBeAnEmptyListOfTasks)
	ctx.Step(`^I would expect there to be three or more jobs returned in a list$`, f.iWouldExpectThereToBeThreeOrMoreJobsReturnedInAList)
	ctx.Step(`^I would expect there to be four jobs returned in a list$`, f.iWouldExpectThereToBeFourJobsReturnedInAList)
	ctx.Step(`^I would expect there to be (\d+) tasks returned in a list$`, f.iWouldExpectThereToBeTasksReturnedInAList)
	ctx.Step(`^in each job I would expect the response to contain values that have these structures$`, f.inEachJobIWouldExpectTheResponseToContainValuesThatHaveTheseStructures)
	ctx.Step(`^the response for getting task to look like this$`, f.expectTaskToLookLikeThis)
	ctx.Step(`^no tasks have been created in the tasks collection$`, f.noTasksHaveBeenCreatedInTheTasksCollection)
	ctx.Step(`^the job should only be updated with the following fields and values$`, f.theJobShouldOnlyBeUpdatedWithTheFollowingFieldsAndValues)
	ctx.Step(`^the jobs should be ordered, by last_updated, with the oldest first$`, f.theJobsShouldBeOrderedByLastupdatedWithTheOldestFirst)
	ctx.Step(`^the search reindex api loses its connection to mongo DB$`, f.theSearchReindexAPILosesItsConnectionToMongoDB)
	ctx.Step(`^the tasks should be ordered, by last_updated, with the oldest first$`, f.theTasksShouldBeOrderedByLastupdatedWithTheOldestFirst)
	ctx.Step(`^the task should have the following fields and values$`, f.theTaskShouldHaveTheFollowingFieldsAndValues)
	ctx.Step(`^the reindex-requested event should contain the expected job ID and search index name$`, f.theReindexrequestedEventShouldContainTheExpectedJobIDAndSearchIndexName)

	ctx.Step(`^check the database has the following task document stored in task collection$`, f.theTaskResourceShouldHaveTheFollowingFieldsAndValuesInDatastore)

	ctx.Step(`^the response ETag header should not be empty$`, f.theResponseETagHeaderShouldNotBeEmpty)
	ctx.Step(`^the response should also contain the following values:$`, f.theResponseShouldAlsoContainTheFollowingValues)
	ctx.Step(`^the response should contain a state of "([^"]*)"$`, f.theResponseShouldContainAStateOf)
	ctx.Step(`^the response should contain the new number of tasks$`, f.theResponseShouldContainTheNewNumberOfTasks)
	ctx.Step(`^the response should contain values that have these structures$`, f.theResponseShouldContainValuesThatHaveTheseStructures)
	ctx.Step(`^the search reindex api loses its connection to the search api$`, f.theSearchReindexAPILosesItsConnectionToTheSearchAPI)
}
