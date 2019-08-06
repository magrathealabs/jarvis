package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/magrathealabs/jarvis/domain/enums"
	"github.com/magrathealabs/jarvis/domain/services/forms"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/datastore"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

type TemperatureHandlerSuite struct {
	gspec.Suite

	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

func (suite *TemperatureHandlerSuite) TestNewTemperatureHandler() {
	suite.Engine = gin.Default()

	NewTemperatureHandler(datastore.NewMetricRepositoryFromEnv()).SetupRoutes(suite.Engine)

	suite.NotNil(NewTemperatureHandler(datastore.NewMetricRepositoryFromEnv()))
}

func (suite *TemperatureHandlerSuite) TestIndex() {
	form := forms.NewStoresTemperature()
	form.RecordedAt = time.Now()
	form.RecordedBy = "beacon"
	form.Temperature = 45
	form.TemperatureScale = enums.CelsiusTemperaureScale

	data, marshalErr := json.Marshal(form)
	suite.NoError(marshalErr)

	suite.Engine = gin.Default()

	NewTemperatureHandler(datastore.NewMetricRepositoryFromEnv()).SetupRoutes(suite.Engine)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodPost, routes.APIV1Temperature, bytes.NewReader(data))
	suite.NoError(err)

	request.Header.Set("Content-Type", "application/json")

	suite.Engine.ServeHTTP(recorder, request)
	suite.Equal(http.StatusOK, recorder.Code)
}

func TestTemperatureHandlerSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureHandlerSuite))
}
