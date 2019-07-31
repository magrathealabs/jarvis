package validators

import "time"

// TimeValidator verify time structs
type TimeValidator struct {
	Field time.Time
}

// NewTimeValidator constructor
func NewTimeValidator(time time.Time) *TimeValidator {
	return &TimeValidator{Field: time}
}

// NotInTheFuture verify
func (validator *TimeValidator) NotInTheFuture() bool {
	return validator.Field.Before(time.Now())
}
