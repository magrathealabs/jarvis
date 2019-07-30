package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

type RootHandlerSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *RootHandlerSuite) SetupTest() {
	suite.Engine = gin.Default()
	suite.Recorder = httptest.NewRecorder()

	NewRootHandler(nil).SetupRoutes(suite.Engine)
}

func (suite *RootHandlerSuite) TestNewRootHandler() {
	suite.NotNil(NewRootHandler(nil))
}

func (suite *RootHandlerSuite) TestIndex() {
	request, err := http.NewRequest(http.MethodGet, routes.Root, strings.NewReader(""))
	suite.NoError(err)

	suite.Engine.ServeHTTP(suite.Recorder, request)
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func TestRootHandlerSuite(t *testing.T) {
	gspec.Run(t, new(RootHandlerSuite))
}
