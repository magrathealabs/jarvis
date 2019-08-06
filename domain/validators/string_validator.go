package validators

// StringValidator struct
type StringValidator struct {
	Field string
}

// NewStringValidator constructor
func NewStringValidator(field string) *StringValidator {
	return &StringValidator{Field: field}
}

// Nil value
func (validator *StringValidator) Nil() bool {
	return validator.Field == ""
}

// NotNil value
func (validator *StringValidator) NotNil() bool {
	return !validator.Nil()
}
