package validators

// Float32Validator struct
type Float32Validator struct {
	Field float32
}

// NewFloat32Validator constructor
func NewFloat32Validator(field float32) *Float32Validator {
	return &Float32Validator{Field: field}
}

// InBetween two values
func (validator *Float32Validator) InBetween(min, max float32) bool {
	return validator.Field >= min && validator.Field <= max
}
