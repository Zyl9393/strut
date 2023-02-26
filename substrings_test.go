package strut_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/Zyl9393/strut"
	"github.com/stretchr/testify/assert"
)

func TestIterateSubstrings(t *testing.T) {
	tests := []struct {
		s                     string
		minSubstringRuneCount int
		maxSubstringRuneCount int
		expected              [][2]int
	}{
		{"", 1, math.MaxInt, [][2]int{}},
		{"cat", 1, math.MaxInt, [][2]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}, {2, 3}}},
		{"cat", 1, 3, [][2]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}, {2, 3}}},
		{"cat", 1, 2, [][2]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}}},
		{"cat", 1, 1, [][2]int{{0, 1}, {1, 2}, {2, 3}}},
		{"cat", 2, 3, [][2]int{{0, 2}, {0, 3}, {1, 3}}},
		{"cat", 3, 3, [][2]int{{0, 3}}},
		{"cat", 3, 4, [][2]int{{0, 3}}},
		{"cat", 4, 3, [][2]int{}},
		{"cat", math.MaxInt, 3, [][2]int{}},
		{"cat", math.MaxInt, math.MaxInt, [][2]int{}},
		{"äh", 1, 2, [][2]int{{0, 2}, {0, 3}, {2, 3}}},
		{"annn", 1, 2, [][2]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 4}}},
		{"annn", 1, 0, [][2]int{}},
	}
	indexPairs := make([][2]int, 0, 100000)
	for i, test := range tests {
		indexPairs = indexPairs[:0]
		strut.IterateSubstrings(test.s, test.minSubstringRuneCount, test.maxSubstringRuneCount, func(from, to int) {
			indexPairs = append(indexPairs, [2]int{from, to})
		})
		if !reflect.DeepEqual(indexPairs, test.expected) {
			t.Errorf("Test %d: IterateSubstrings(`%s`, %d, %d, *) returned %v. Expected %v.", i, test.s, test.minSubstringRuneCount, test.maxSubstringRuneCount, indexPairs, test.expected)
		}
	}
	assert.PanicsWithValue(t, "minSubstringRuneCount must be greater than 0.", func() { strut.IterateSubstrings("", 0, 0, func(from, to int) {}) })
	assert.PanicsWithValue(t, "maxSubstringRuneCount must not be negative.", func() { strut.IterateSubstrings("", 1, -1, func(from, to int) {}) })
	assert.PanicsWithValue(t, "iterFunc must not be nil.", func() { strut.IterateSubstrings("", 1, 1, nil) })
}

func TestIterateSubstringsUnique(t *testing.T) {
	tests := []struct {
		s                     string
		minSubstringRuneCount int
		maxSubstringRuneCount int
		expected              [][2]int
	}{
		{"", 1, math.MaxInt, [][2]int{}},
		{"cat", 1, math.MaxInt, [][2]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}, {2, 3}}},
		{"cat", 1, 3, [][2]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}, {2, 3}}},
		{"cat", 1, 2, [][2]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}}},
		{"cat", 1, 1, [][2]int{{0, 1}, {1, 2}, {2, 3}}},
		{"cat", 2, 3, [][2]int{{0, 2}, {0, 3}, {1, 3}}},
		{"cat", 3, 3, [][2]int{{0, 3}}},
		{"cat", 3, 4, [][2]int{{0, 3}}},
		{"cat", 4, 3, [][2]int{}},
		{"cat", math.MaxInt, 3, [][2]int{}},
		{"cat", math.MaxInt, math.MaxInt, [][2]int{}},
		{"äh", 1, 2, [][2]int{{0, 2}, {0, 3}, {2, 3}}},
		{"annn", 1, 2, [][2]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}}},
		{"annn", 0, 2, [][2]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}}},
		{"annn", 1, 0, [][2]int{}},
		{"annn", 0, 0, [][2]int{{0, 0}}},
	}
	result := make([][2]int, 0, 64)
	for i, test := range tests {
		result = result[:0]
		strut.IterateSubstringsUnique(test.s, test.minSubstringRuneCount, test.maxSubstringRuneCount, func(from, to int) {
			result = append(result, [2]int{from, to})
		})
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: IterateSubstringsUnique(`%s`, %d, %d, *) returned %v. Expected %v.", i, test.s, test.minSubstringRuneCount, test.maxSubstringRuneCount, result, test.expected)
		}
	}
	assert.PanicsWithValue(t, "minSubstringRuneCount must not be negative.", func() { strut.IterateSubstringsUnique("", -1, 0, func(from, to int) {}) })
	assert.PanicsWithValue(t, "maxSubstringRuneCount must not be negative.", func() { strut.IterateSubstringsUnique("", 0, -1, func(from, to int) {}) })
	assert.PanicsWithValue(t, "iterFunc must not be nil.", func() { strut.IterateSubstringsUnique("", 0, 0, nil) })
}
