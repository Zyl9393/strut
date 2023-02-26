package strut

// Returns a function which takes a bool parameter that returns trueString if the parameter is true and falseString if the parameter is false.
func NewBoolToStringFunc(trueString, falseString string) func(bool) string {
	return func(b bool) string {
		if b {
			return trueString
		}
		return falseString
	}
}
