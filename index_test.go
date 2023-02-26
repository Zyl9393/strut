package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
	"github.com/stretchr/testify/assert"
)

func TestIndexRune(t *testing.T) {
	tests := []struct {
		s         string
		runeIndex int
		expected  int
	}{
		{"", 0, 0},
		{"a", 0, 0},
		{"a", 1, 1},
		{"Käse", 0, 0},
		{"Käse", 1, 1},
		{"Käse", 2, 3},
		{"Käse", 3, 4},
		{"Käse", 4, 5},
	}
	for i, test := range tests {
		result := strut.IndexRune(test.s, test.runeIndex)
		if result != test.expected {
			t.Errorf("Test %d: IndexRune(`%s`, %d) returned %d. Expected %d. (len(`%s`) == %d)", i, test.s, test.runeIndex, result, test.expected, test.s, len(test.s))
		}
	}
	assert.PanicsWithValue(t, "runeIndex < 0", func() { strut.IndexRune("", -1) })
	assert.PanicsWithValue(t, "runeIndex is out of bounds", func() { strut.IndexRune("", 1) })
}

func TestRuneIndex(t *testing.T) {
	tests := []struct {
		s         string
		byteIndex int
		expected  int
	}{
		{"", 0, 0},
		{"a", 0, 0},
		{"a", 1, 1},
		{"Käse", 0, 0},
		{"Käse", 1, 1},
		{"Käse", 2, 1},
		{"Käse", 3, 2},
		{"Käse", 4, 3},
		{"Käse", 5, 4},
		{"ä", 0, 0},
		{"ä", 1, 0},
		{"ä", 2, 1},
	}
	for i, test := range tests {
		result := strut.RuneIndex(test.s, test.byteIndex)
		if result != test.expected {
			t.Errorf("Test %d: RuneIndex(`%s`, %d) returned %d. Expected %d. (len(`%s`) == %d)", i, test.s, test.byteIndex, result, test.expected, test.s, len(test.s))
		}
	}
	assert.PanicsWithValue(t, "i < 0", func() { strut.RuneIndex("", -1) })
	assert.PanicsWithValue(t, "i is out of bounds", func() { strut.RuneIndex("", 1) })
}

func TestNumIter(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"", 0},
		{"x", 1},
		{"xx", 2},
		{"xy", 2},
		{"\n", 1},
		{"\r", 1},
		{"\r\n", 2},
		{"Käse", 4},
		{"Käse", 5},
	}
	for i, test := range tests {
		result := strut.RangeLen(test.s)
		if result != test.expected {
			t.Errorf("Test %d: NumIter(`%s`) returned %d. Expected %d.", i, test.s, result, test.expected)
		}
	}
}

func TestSliceRunes(t *testing.T) {
	tests := []struct {
		s             string
		fromRuneIndex int
		toRuneIndex   int
		expected      string
	}{
		{"", 0, 0, ""},
		{"x", 0, 0, ""},
		{"x", 1, 1, ""},
		{"x", 0, 1, "x"},
		{"Käse", 1, 2, "ä"},
		{"Käse", 2, 4, "se"},
	}
	for i, test := range tests {
		result := strut.SliceRunes(test.s, test.fromRuneIndex, test.toRuneIndex)
		if result != test.expected {
			t.Errorf("Test %d: SliceRunes(`%s`, %d, %d) returned `%s`. Expected `%s`.", i, test.s, test.fromRuneIndex, test.toRuneIndex, result, test.expected)
		} else {
			altResult := test.s[strut.IndexRune(test.s, test.fromRuneIndex):strut.IndexRune(test.s, test.toRuneIndex)]
			if altResult != result {
				t.Errorf("Test %d: %#q[IndexRune(*, %d) : IndexRune(*, %d)] returned `%s`. Expected `%s`.", i, test.s, test.fromRuneIndex, test.toRuneIndex, altResult, result)
			}
		}
	}
	assert.PanicsWithValue(t, "fromRuneIndex must not be less than zero", func() { strut.SliceRunes("", -1, 0) })
	assert.PanicsWithValue(t, "toRuneIndex must not be less than zero", func() { strut.SliceRunes("", 0, -1) })
	assert.PanicsWithValue(t, "fromRuneIndex must not be greater than toRuneIndex", func() { strut.SliceRunes("", 1, 0) })
	assert.PanicsWithValue(t, "toRuneIndex overflows number of runes in s", func() { strut.SliceRunes("", 0, 1) })
}

func TestSliceClusters(t *testing.T) {
	tests := []struct {
		s                        string
		fromGraphemeClusterIndex int
		toGraphemeClusterIndex   int
		expected                 string
	}{
		{"", 0, 0, ""},
		{"x", 0, 0, ""},
		{"x", 1, 1, ""},
		{"x", 0, 1, "x"},
		{"Käse", 1, 2, "ä"},
		{"Käse", 2, 4, "se"},
	}
	for i, test := range tests {
		result := strut.SliceClusters(test.s, test.fromGraphemeClusterIndex, test.toGraphemeClusterIndex)
		if result != test.expected {
			t.Errorf("Test %d: SliceClusters(`%s`, %d, %d) returned `%s`. Expected `%s`.", i, test.s, test.fromGraphemeClusterIndex, test.toGraphemeClusterIndex, result, test.expected)
		}
	}
	assert.PanicsWithValue(t, "fromGraphemeClusterIndex must not be less than zero", func() { strut.SliceClusters("", -1, 0) })
	assert.PanicsWithValue(t, "toGraphemeClusterIndex must not be less than zero", func() { strut.SliceClusters("", 0, -1) })
	assert.PanicsWithValue(t, "fromGraphemeClusterIndex must not be greater than toGraphemeClusterIndex", func() { strut.SliceClusters("", 1, 0) })
	assert.PanicsWithValue(t, "toGraphemeClusterIndex overflows number of grapheme clusters in s", func() { strut.SliceClusters("", 0, 1) })
}
