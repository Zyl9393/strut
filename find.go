package strut

// IndexOfLongest returns the index of the first string for which no longer string exists in s,
// or -1 if len(s) == 0.
func IndexOfLongest(s []string) int {
	if len(s) == 0 {
		return -1
	}
	longestIndex := 0
	for i, element := range s {
		if len(element) > len(s[longestIndex]) {
			longestIndex = i
		}
	}
	return longestIndex
}

// IndexOfShortest returns the index of the first string for which no shorter string exists in s,
// or -1 if len(s) == 0.
func IndexOfShortest(s []string) int {
	if len(s) == 0 {
		return -1
	}
	shortestIndex := 0
	for i, element := range s {
		if len(element) < len(s[shortestIndex]) {
			shortestIndex = i
		}
	}
	return shortestIndex
}
