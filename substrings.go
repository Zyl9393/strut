package strut

// IterateSubstrings calls iterFunc for every substring of s which has at least a length of minSubstringRuneCount
// and at most a length of maxSubstringRuneCount.
// When iterFunc is called, the substring is given by s[from:to].
func IterateSubstrings(s string, minSubstringRuneCount int, maxSubstringRuneCount int, iterFunc func(from, to int)) {
	if minSubstringRuneCount < 1 {
		panic("minSubstringRuneCount must be greater than 0.")
	}
	if maxSubstringRuneCount < 0 {
		panic("maxSubstringRuneCount must not be negative.")
	}
	if minSubstringRuneCount > maxSubstringRuneCount {
		return
	}
	if iterFunc == nil {
		panic("iterFunc must not be nil.")
	}
	indices := make([]int, len(s)+1)
	runeIndex := 0
	for i := range s + "." {
		indices[runeIndex] = i
		jStart := -maxSubstringRuneCount
		if jStart < -runeIndex {
			jStart = -runeIndex
		}
		for j := jStart; j <= -minSubstringRuneCount; j++ {
			iterFunc(indices[runeIndex+j], i)
		}
		runeIndex++
	}
}

// Like IterateSubstrings, but never invokes iterFunc more than once for multiple occurrences of equal substrings.
// This function also allows to set minSubstringRuneCount to 0 to cause iterFunc(0, 0) to be called with respect to the empty substring.
func IterateSubstringsUnique(s string, minSubstringRuneCount int, maxSubstringRuneCount int, iterFunc func(from, to int)) {
	if minSubstringRuneCount < 0 {
		panic("minSubstringRuneCount must not be negative.")
	}
	if maxSubstringRuneCount < 0 {
		panic("maxSubstringRuneCount must not be negative.")
	}
	if minSubstringRuneCount > maxSubstringRuneCount {
		return
	}
	if iterFunc == nil {
		panic("iterFunc must not be nil.")
	}
	if minSubstringRuneCount == 0 {
		iterFunc(0, 0)
		minSubstringRuneCount++
	}
	if minSubstringRuneCount > maxSubstringRuneCount {
		return
	}
	indices := make([]int, len(s)+1)
	runeIndex := 0
	seen := make(map[string]bool)
	for i := range s + "." {
		indices[runeIndex] = i
		jStart := -maxSubstringRuneCount
		if jStart < -runeIndex {
			jStart = -runeIndex
		}
		for j := jStart; j <= -minSubstringRuneCount; j++ {
			if _, ok := seen[s[indices[runeIndex+j]:i]]; !ok {
				seen[s[indices[runeIndex+j]:i]] = true
				iterFunc(indices[runeIndex+j], i)
			}
		}
		runeIndex++
	}
}
