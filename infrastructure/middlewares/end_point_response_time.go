package middlewares

import (
	"time"

	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/domain/services"

	"github.com/gin-gonic/gin"
)

// EndpointRuntime middleware
func EndpointRuntime(repository repositories.MetricRepository) gin.HandlerFunc {
	runtimeMetricsService := services.NewRuntimeMetricsService(repository)

	return func(c *gin.Context) {
		startAt := time.Now()
		c.Next()
		url := c.Request.URL.EscapedPath()

		_ = runtimeMetricsService.StoresEndpointRuntime(url, startAt)
	}
}
