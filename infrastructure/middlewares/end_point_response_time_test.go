package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"

	"github.com/krakenlab/gspec"
)

type EndpointRuntimeSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *EndpointRuntimeSuite) SetupTest() {
	suite.Engine = gin.Default()
	suite.Recorder = httptest.NewRecorder()

	repository := datastore.NewMetricRepositoryFromEnv()
	suite.Engine.Use(EndpointRuntime(repository))

	suite.Engine.GET("/", func(c *gin.Context) {
		c.String(200, "Ok")
	})
}

func (suite *EndpointRuntimeSuite) TestEndpointRuntime() {
	repository := datastore.NewMetricRepositoryFromEnv()
	suite.NotNil(EndpointRuntime(repository))
}

func (suite *EndpointRuntimeSuite) TestGet() {
	request, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	suite.NoError(err)

	suite.Engine.ServeHTTP(suite.Recorder, request)
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func TestRootHandlerSuite(t *testing.T) {
	gspec.Run(t, new(EndpointRuntimeSuite))
}
