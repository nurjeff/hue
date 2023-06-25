package hue_controller

// Returns a pointer to any value
func VPtrs[T any](value T) *T {
	return &value
}
