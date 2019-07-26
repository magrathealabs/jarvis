package web

import (
	"github.com/gin-gonic/gin"
)

// Web for a web server
type Web struct {
	Server *gin.Engine
	Addr   string
}

// New jarvis web server
func New(addr string) *Web {
	return &Web{Server: gin.Default()}
}

// Setup this web Web
func (web *Web) Setup() {}

// Run using envs
func (web *Web) Run() error {
	return web.Server.Run(web.Addr)
}
