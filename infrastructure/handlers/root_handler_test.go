package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"
)

type RootHandlerSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *RootHandlerSuite) SetupTest() {
	suite.Engine = gin.Default()
	suite.Recorder = httptest.NewRecorder()

	NewRootHandler(datastore.NewMetricRepositoryFromEnv()).SetupRoutes(suite.Engine)
}

func (suite *RootHandlerSuite) TestNewRootHandler() {
	suite.NotNil(NewRootHandler(datastore.NewMetricRepositoryFromEnv()))
}

func TestRootHandlerSuite(t *testing.T) {
	gspec.Run(t, new(RootHandlerSuite))
}
