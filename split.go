package strut

import (
	"strings"
)

// SplitTwinUnescape splits s around all occurrences of sep,
// except where sep occurs twice in a row which will be reduced
// to a single occurrence.
//
// If s does not contain an isolated occurrence of sep,
// this function returns a slice whose only element is s.
//
// If sep is the empty string, this function panics.
func SplitTwinUnescape(s string, sep string) (unescapedSubstrings []string) {
	if sep == "" {
		panic("sep must not be the empty string")
	}
	var from, relativeSepIndex, to int
	var sb strings.Builder
	for relativeSepIndex != -1 {
		relativeSepIndex = strings.Index(s[from:], sep)
		to = from + relativeSepIndex
		if relativeSepIndex == -1 {
			to = len(s)
		}
		sb.WriteString(s[from:to])
		if to+len(sep) <= len(s) && strings.HasPrefix(s[to+len(sep):], sep) {
			sb.WriteString(sep)
			from = to + 2*len(sep)
			continue
		}
		unescapedSubstrings = append(unescapedSubstrings, sb.String())
		sb.Reset()
		from = to + len(sep)
	}
	return unescapedSubstrings
}
