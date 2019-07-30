package models

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type BaseSuite struct {
	gspec.Suite
}

func (suite *BaseSuite) TestNewBase() {
	suite.NotNil(NewBase())
}

func (suite *BaseSuite) TestToJSON() {
	base := NewBase()
	json := base.ToJSON()

	suite.Contains(json, "id")
	suite.Contains(json, "created_at")
	suite.Contains(json, "deleted_at")
	suite.Contains(json, "updated_at")
}

func TestBaseSuite(t *testing.T) {
	gspec.Run(t, new(BaseSuite))
}
