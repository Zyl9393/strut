package strut

import "strings"

// TrimPrefixAll returns a new string with all occurrences of prefix at the start of s removed.
// If prefix is the empty string, this function returns s.
func TrimPrefixAll(s string, prefix string) string {
	if prefix == "" {
		return s
	}
	for strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	return s
}

// TrimPrefixN returns a new string with up to n occurrences of prefix at the start of s removed.
// If prefix is the empty string, this function returns s.
// If n is negative, returns TrimPrefixAll(s, prefix).
func TrimPrefixN(s string, prefix string, n int) string {
	if n < 0 {
		return TrimPrefixAll(s, prefix)
	}
	if prefix == "" {
		return s
	}
	for n > 0 && strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
		n--
	}
	return s
}

// TrimSuffixAll returns a new string with all occurrences of suffix at the end of s removed.
// If suffix is the empty string, this function returns s.
func TrimSuffixAll(s string, suffix string) string {
	if suffix == "" {
		return s
	}
	for strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// TrimSuffixN returns a new string with up to n occurrences of suffix at the end of s removed.
// If suffix is the empty string, this function returns s.
// If n is negative, returns TrimSuffixAll(s, suffix).
func TrimSuffixN(s string, suffix string, n int) string {
	if n < 0 {
		return TrimSuffixAll(s, suffix)
	}
	if suffix == "" {
		return s
	}
	for n > 0 && strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
		n--
	}
	return s
}

// TrimSequenceAll returns a new string with all occurrences of sequence at the start and end of s removed.
// If sequence is the empty string, this function returns s.
func TrimSequenceAll(s string, sequence string) string {
	return TrimSuffixAll(TrimPrefixAll(s, sequence), sequence)
}
