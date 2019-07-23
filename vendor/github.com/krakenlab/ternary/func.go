package ternary

// Func ternary
func Func(logic bool, firstFunc func(), secondFunc func()) func() {
	if logic {
		return firstFunc
	}

	return secondFunc
}
