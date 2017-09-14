package ctx

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/persistence"
	handler "bitbucket.org/instinctools/gluten/core/result_handlers"
	"bitbucket.org/instinctools/gluten/master/backend/clustering"
	"bitbucket.org/instinctools/gluten/master/backend/config"
	"bitbucket.org/instinctools/gluten/master/backend/service"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
)

var (
	GlobalContext = loadContext()
)

type ApplicationContext struct {
	AppConfig        *config.Config
	NodeStore        *clustering.NodeStore
	Runner           core.TestRunner
	RawExecutionRepo *gorm.RawExecutionRepo
	RawResultRepo    *gorm.RawResultsRepo
	ExecutionRepo    persistence.ExecutionRepo
	ResultRepo       persistence.ResultRepo
	ExecutionService *service.ExecutionService
	ProjectStore     service.ProjectStore
}

func loadContext() *ApplicationContext {
	appConfig := config.GetDefaultConfig()
	dbConn := gorm.NewDBConnection(appConfig.DB.Connection.URL)
	rawExecutionRepo := gorm.NewRawExecutionRepo(dbConn)
	rawResultRepo := gorm.NewRawResultRepo(dbConn)
	executionRepo := gorm.NewExecutionRepoWrapper(rawExecutionRepo)

	runner := core.NewDefaultRunner(&handler.LoggableResultHandler{})
	projectStore := service.NewInMemoryProjectStore()

	return &ApplicationContext{
		AppConfig:        appConfig,
		NodeStore:        clustering.NewNodeStore(appConfig.Node.RetrieveTimeout, appConfig.Node.ExitTimeout, clustering.SubmitOverRPC),
		Runner:           runner,
		RawExecutionRepo: rawExecutionRepo,
		RawResultRepo:    rawResultRepo,
		ExecutionRepo:    executionRepo,
		ExecutionService: service.NewExecutionService(runner, executionRepo, projectStore),
		ProjectStore:     projectStore,
	}
}
