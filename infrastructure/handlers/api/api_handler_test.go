package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"
)

type APIHandlerSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *APIHandlerSuite) SetupTest() {
	suite.Engine = gin.Default()
	suite.Recorder = httptest.NewRecorder()

	NewHandler(datastore.NewMetricRepositoryFromEnv()).SetupRoutes(suite.Engine)
}

func (suite *APIHandlerSuite) TestNewHandler() {
	suite.NotNil(NewHandler(datastore.NewMetricRepositoryFromEnv()))
}

func (suite *APIHandlerSuite) TestSetupRoutes() {
	suite.Equal(1, len(suite.Engine.Routes()))
}

func TestAPIHandlerSuite(t *testing.T) {
	gspec.Run(t, new(APIHandlerSuite))
}
