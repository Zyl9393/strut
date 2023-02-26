package strut

// CommonPrefix returns the longest prefix common to all strings of s.
func CommonPrefix(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	if len(ss) == 1 {
		return ss[0]
	}
	shortestMatch := ss[IndexOfShortest(ss)]
	for i := range ss {
		if l := PrefixLen(shortestMatch, ss[i]); l < len(shortestMatch) {
			shortestMatch = shortestMatch[:l]
			if shortestMatch == "" {
				return ""
			}
		}
	}
	return shortestMatch
}

// PrefixLen returns the length of the longest common prefix of p and q.
func PrefixLen(p string, q string) int {
	limit := len(p)
	if limit > len(q) {
		limit = len(q)
	}
	for i := 0; i < limit; i++ {
		if p[i] != q[i] {
			return i
		}
	}
	return limit
}

// CommonSuffix returns the longest suffix common to all strings of s.
func CommonSuffix(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	if len(ss) == 1 {
		return ss[0]
	}
	shortestMatch := ss[IndexOfShortest(ss)]
	for i := range ss {
		if l := SuffixLen(shortestMatch, ss[i]); l < len(shortestMatch) {
			shortestMatch = shortestMatch[len(shortestMatch)-l:]
			if shortestMatch == "" {
				return ""
			}
		}
	}
	return shortestMatch
}

// SuffixLen returns the length of the longest common suffix of p and q.
func SuffixLen(p string, q string) int {
	limit := len(p)
	if limit > len(q) {
		limit = len(q)
	}
	for i := 1; i <= limit; i++ {
		if p[len(p)-i] != q[len(q)-i] {
			return i - 1
		}
	}
	return limit
}
