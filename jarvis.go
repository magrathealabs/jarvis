package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"
	"github.com/magrathealabs/jarvis/infrastructure/handlers"
	"github.com/magrathealabs/jarvis/libs/env"
)

func metricRepository() repositories.MetricRepository {
	return datastore.NewMetricRepository(env.Graphite().Host(), env.Graphite().PortAsInt())
}

func engine() *gin.Engine {
	engine := gin.Default()

	handlers.NewRootHandler(metricRepository()).SetupRoutes(engine)

	return engine
}

func main() {
	panic(engine().Run(env.Server().Addr()))
}
