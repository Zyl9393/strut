package strut

import "strconv"

// Translate returns translations[value - firstIndex], except that if index
// would be out of bounds, it returns strconv.Itoa(value) instead.
// This function is intended to obtain humanly readable names for package-level
// constants which cannot be found with reflection.
//
// E.g. Translate(Mouse, Cat, "cat", "dog", "mouse", "bird") returns "mouse",
// given that the value of Mouse equals the value of Cat+2.
func Translate(value int, firstIndex int, translations ...string) string {
	index := value - firstIndex
	if index < 0 || index >= len(translations) {
		return strconv.Itoa(value)
	}
	return translations[value-firstIndex]
}
