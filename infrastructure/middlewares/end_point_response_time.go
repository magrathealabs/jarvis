package middlewares

import (
	"time"

	"github.com/magrathealabs/jarvis/domain/models"
	"github.com/magrathealabs/jarvis/domain/repositories"

	"github.com/gin-gonic/gin"
)

// EndPointResponseTime middleware
func EndPointResponseTime(repository repositories.MetricRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		url := c.Request.URL.EscapedPath()
		elapsed := time.Since(start)

		nano := float64(int64(elapsed / time.Millisecond))

		endPointResponseTime := models.NewEndPointResponseTime(nano, url)
		_ = repository.InsertEndPointResponseTime(endPointResponseTime)
	}
}
