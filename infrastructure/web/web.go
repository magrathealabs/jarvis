package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/env"
)

// Engine for a web server
type Engine struct {
	Server *gin.Engine
}

// New jarvis web server
func New() *Engine {
	return &Engine{Server: gin.Default()}
}

// Setup this web engine
func (engine *Engine) Setup() {}

// Run using envs
func (engine *Engine) Run() error {
	return engine.Server.Run(engine.ginInterfaceFormat())
}

func (engine *Engine) ginInterfaceFormat() string {
	return fmt.Sprintf("%s:%s", env.Web().Interface(), env.Web().Port())
}
