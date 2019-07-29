package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/handlers"
	"github.com/magrathealabs/jarvis/libs/env"
)

func usedHandlers() []handlers.Handler {
	return []handlers.Handler{
		handlers.NewRootHandler(),
	}
}

func engine() *gin.Engine {
	engine := gin.Default()

	for _, handler := range usedHandlers() {
		handler.SetupRoutes(engine)
	}

	return engine
}

func main() {
	panic(engine().Run(env.Server().Addr()))
}
