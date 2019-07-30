package handler

import "github.com/gin-gonic/gin"

// Handler interface
type Handler interface {
	SetupRoutes(engine *gin.Engine)
}
