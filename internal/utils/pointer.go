package utils

// Bool returns a pointer to a bool:
func Bool(value bool) *bool {
	return &value
}

// Int32 returns a pointer to an int32:
func Int32(value int32) *int32 {
	return &value
}

// String returns a pointer to a string:
func String(value string) *string {
	return &value
}
