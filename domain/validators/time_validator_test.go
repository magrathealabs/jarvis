package validators

import (
	"testing"
	"time"

	"github.com/krakenlab/gspec"
)

type TimeValidatorSuite struct {
	gspec.Suite
}

func (suite *TimeValidatorSuite) TestNewTimeValidator() {
	suite.NotNil(NewTimeValidator(time.Now()))
}

func (suite *TimeValidatorSuite) TestNotInTheFuture() {
	suite.True(NewTimeValidator(time.Now()).NotInTheFuture())
	suite.False(NewTimeValidator(time.Now().Add(time.Hour)).NotInTheFuture())
}

func (suite *TimeValidatorSuite) TestInTheFuture() {
	suite.False(NewTimeValidator(time.Now()).InTheFuture())
	suite.True(NewTimeValidator(time.Now().Add(time.Hour)).InTheFuture())
}

func TestTimeValidatorSuite(t *testing.T) {
	gspec.Run(t, new(TimeValidatorSuite))
}
