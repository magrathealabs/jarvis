package exceptions

import (
	"errors"
	"fmt"
)

func init() {
	ErrorRecordedInMeyNotBeInTheFuture = mayNotBeInTheFuture("recorded_in")
	ErrorTemperatureScaleMstBeCFK = scaleMustBeCFK("temperature")
	ErrorUnfortunatelyWeHeventMetTemperaturesFromOtherPlanetsYet = unfortunatelyWeHeventMetTemperaturesFromOtherPlanetsYet()
	ErrorPercentageMustBeBetween0and1 = mustBeBetweenXandY("percentage", "0.0", "1.0")
	ErrorRecordedByNeedsToBePresent = needsToBePresent("recorded_by")
}

// Model error
var (
	ErrorRecordedInMeyNotBeInTheFuture                           error
	ErrorTemperatureScaleMstBeCFK                                error
	ErrorUnfortunatelyWeHeventMetTemperaturesFromOtherPlanetsYet error
	ErrorPercentageMustBeBetween0and1                            error
	ErrorRecordedByNeedsToBePresent                              error
)

func mayNotBeInTheFuture(field string) error {
	return fmt.Errorf("%s may not be in the future", field)
}

func scaleMustBeCFK(field string) error {
	return fmt.Errorf("%s scale must be C|F|K", field)
}

func unfortunatelyWeHeventMetTemperaturesFromOtherPlanetsYet() error {
	return errors.New("unfortunately we haven't met temperatures from other planets yet")
}

func mustBeBetweenXandY(field, x, y string) error {
	return fmt.Errorf("%s must be between %s and %s", field, x, y)
}

func needsToBePresent(field string) error {
	return fmt.Errorf("%s needs to be present", field)
}
