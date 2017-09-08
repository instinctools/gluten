package ctx

import (
	"bitbucket.org/instinctools/gluten/master/backend/clustering"
	"bitbucket.org/instinctools/gluten/master/backend/config"
)

var (
	GlobalContext = loadContext()
)

type ApplicationContext struct {
	AppConfig *config.Config
	NodeStore *clustering.NodeStore
}

func loadContext() *ApplicationContext {
	appConfig := config.GetDefaultConfig()

	return &ApplicationContext{
		AppConfig: appConfig,
		NodeStore: clustering.NewNodeStore(appConfig.Node.RetrieveTimeout, appConfig.Node.ExitTimeout, clustering.SubmitOverRPC),
	}

}
