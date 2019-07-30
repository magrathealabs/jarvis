package v1

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

type TemperatureHandlerSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *TemperatureHandlerSuite) SetupTest() {
	suite.Engine = gin.Default()
	suite.Recorder = httptest.NewRecorder()

	NewTemperatureHandler().SetupRoutes(suite.Engine)
}

func (suite *TemperatureHandlerSuite) TestNewTemperatureHandler() {
	suite.NotNil(NewTemperatureHandler())
}

func (suite *TemperatureHandlerSuite) TestSetupRoutes() {
	suite.Equal(1, len(suite.Engine.Routes()))
}

func (suite *TemperatureHandlerSuite) TestIndex() {
	request, err := http.NewRequest(http.MethodGet, routes.APIV1Temperature, strings.NewReader(""))
	suite.NoError(err)

	suite.Engine.ServeHTTP(suite.Recorder, request)
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func TestTemperatureHandlerSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureHandlerSuite))
}
