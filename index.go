package strut

import "github.com/rivo/uniseg"

// IndexRune returns the byte index of rune index runeIndex in O(n).
// This function panics if runeIndex is less than zero or greater than the amount of runes in s.
func IndexRune(s string, runeIndex int) int {
	if runeIndex < 0 {
		panic("runeIndex < 0")
	}
	currentRuneIndex := 0
	for currentByteIndex := range s {
		if currentRuneIndex == runeIndex {
			return currentByteIndex
		}
		currentRuneIndex++
	}
	if currentRuneIndex == runeIndex {
		return len(s)
	}
	panic("runeIndex is out of bounds")
}

// RuneIndex returns the rune index of the rune at byte index i in O(n).
// If i does not index the firts byte of a rune, returns the index of the rune which that byte is a part of.
// This function panics if i is less than zero or greater than len(s).
func RuneIndex(s string, i int) int {
	if i < 0 {
		panic("i < 0")
	}
	if i > len(s) {
		panic("i is out of bounds")
	}
	currentRuneIndex := 0
	for j := range s {
		if j == i {
			return currentRuneIndex
		} else if j > i {
			return currentRuneIndex - 1
		}
		currentRuneIndex++
	}
	if len(s) == i {
		return currentRuneIndex
	}
	return currentRuneIndex - 1
}

// RangeLen returns the amount of iterations which would take place when iterating over s using the range-expression.
func RangeLen(s string) int {
	rangeLen := 0
	for range s {
		rangeLen++
	}
	return rangeLen
}

// SliceRunes returns the result of s[IndexRune(s, fromRuneIndex) : IndexRune(s, toRuneIndex)] in O(n).
// This function panics if either of the passed indices is less than zero or greater than the amount of runes in s,
// or when fromRuneIndex is greater than toRuneIndex.
func SliceRunes(s string, fromRuneIndex, toRuneIndex int) string {
	if fromRuneIndex < 0 {
		panic("fromRuneIndex must not be less than zero")
	}
	if toRuneIndex < 0 {
		panic("toRuneIndex must not be less than zero")
	}
	if fromRuneIndex > toRuneIndex {
		panic("fromRuneIndex must not be greater than toRuneIndex")
	}
	runeIndex := 0
	fromIndex := len(s)
	for i := range s {
		if runeIndex == fromRuneIndex {
			fromIndex = i
		}
		if runeIndex == toRuneIndex {
			return s[fromIndex:i]
		}
		runeIndex++
	}
	if toRuneIndex != runeIndex {
		panic("toRuneIndex overflows number of runes in s")
	}
	return s[fromIndex:]
}

// SliceClusters is like SliceRunes, except that it operates on grapheme cluster indices rather than rune indices.
func SliceClusters(s string, fromGraphemeClusterIndex, toGraphemeClusterIndex int) string {
	if fromGraphemeClusterIndex < 0 {
		panic("fromGraphemeClusterIndex must not be less than zero")
	}
	if toGraphemeClusterIndex < 0 {
		panic("toGraphemeClusterIndex must not be less than zero")
	}
	if fromGraphemeClusterIndex > toGraphemeClusterIndex {
		panic("fromGraphemeClusterIndex must not be greater than toGraphemeClusterIndex")
	}
	b := []byte(s)
	graphemeClusterIndex := 0
	fromIndex := len(s)
	state := -1
	for len(b) != 0 {
		if graphemeClusterIndex == fromGraphemeClusterIndex {
			fromIndex = len(s) - len(b)
		}
		if graphemeClusterIndex == toGraphemeClusterIndex {
			return s[fromIndex : len(s)-len(b)]
		}
		_, b, _, state = uniseg.FirstGraphemeCluster(b, state)
		graphemeClusterIndex++
	}
	if toGraphemeClusterIndex != graphemeClusterIndex {
		panic("toGraphemeClusterIndex overflows number of grapheme clusters in s")
	}
	return s[fromIndex:]
}
