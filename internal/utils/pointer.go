package utils

// String returns a pointer to a string:
func String(value string) *string {
	return &value
}
