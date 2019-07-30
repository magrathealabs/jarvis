package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/handlers"
	"github.com/magrathealabs/jarvis/libs/env"
)

func engine() *gin.Engine {
	engine := gin.Default()

	handlers.NewRootHandler().SetupRoutes(engine)

	return engine
}

func main() {
	panic(engine().Run(env.Server().Addr()))
}
