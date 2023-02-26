package strut

import (
	"strings"
)

// Slash is like strings.Join(elems, "/"), except that all leading and trailing occurrences of '/'
// between elems are trimmed before they are joined together. Non-trailing leading slashes in the
// first element as well as non-leading trailing slashes in the last element are kept.
func Slash(elems ...string) string {
	return JoinUnary(elems, "/")
}

// JoinUnary is like strings.Join(elems, sep), except that all leading and trailing occurrences of sep
// within elems are trimmed where they are going to be joined with sep beforehand. Non-trailing leading
// separators in the first element as well as non-leading trailing separators in the last element are kept.
func JoinUnary(elems []string, separator string) string {
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	}
	var sb strings.Builder
	const maxGuess = 100
	const guessAverageElementLen = 5
	if len(elems) <= maxGuess {
		sb.Grow((len(elems)-1)*len(separator) + len(elems)*guessAverageElementLen)
	} else {
		sb.Grow((len(elems)-1)*len(separator) + maxGuess*guessAverageElementLen)
	}
	t := TrimSuffixAll(elems[0], separator) + separator
	for _, element := range elems[1 : len(elems)-1] {
		sb.WriteString(t)
		t = TrimSequenceAll(element, separator) + separator
	}
	sb.WriteString(t)
	t = TrimPrefixAll(elems[len(elems)-1], separator)
	sb.WriteString(t)
	return sb.String()
}
