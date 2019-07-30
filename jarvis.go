package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"
	"github.com/magrathealabs/jarvis/infrastructure/handlers"
	"github.com/magrathealabs/jarvis/infrastructure/middlewares"
	"github.com/magrathealabs/jarvis/libs/env"
)

func engine() *gin.Engine {
	engine := gin.Default()
	metricRepository := datastore.NewMetricRepositoryFromEnv()
	engine.Use(middlewares.EndPointResponseTime(metricRepository))

	handlers.NewRootHandler(metricRepository).SetupRoutes(engine)

	return engine
}

func main() {
	panic(engine().Run(env.Server().Addr()))
}
