package utils

// Bool returns a pointer to a bool:
func Bool(value bool) *bool {
	return &value
}

// String returns a pointer to a string:
func String(value string) *string {
	return &value
}
