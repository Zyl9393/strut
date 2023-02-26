package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
)

func TestIndexOfLongest(t *testing.T) {
	tests := []struct {
		s        []string
		expected int
	}{
		{nil, -1},
		{[]string{}, -1},
		{[]string{""}, 0},
		{[]string{"x"}, 0},
		{[]string{"x", "x"}, 0},
		{[]string{"x", "xx"}, 1},
		{[]string{"x", "xx", "xx"}, 1},
		{[]string{"x", "xx", "xxx"}, 2},
		{[]string{"x", "xxx", "xx"}, 1},
		{[]string{"xxx", "xx", "x"}, 0},
		{[]string{"xxx", "xx", "xx"}, 0},
	}
	for i, test := range tests {
		result := strut.IndexOfLongest(test.s)
		if result != test.expected {
			t.Errorf("Test %d: IndexOfLongest({%s}) returned %d. Expected %d.", i, strut.Args(test.s), result, test.expected)
		}
	}
}

func TestIndexOfShortest(t *testing.T) {
	tests := []struct {
		s        []string
		expected int
	}{
		{nil, -1},
		{[]string{}, -1},
		{[]string{""}, 0},
		{[]string{"x"}, 0},
		{[]string{"x", "x"}, 0},
		{[]string{"x", "xx"}, 0},
		{[]string{"x", "xx", "xx"}, 0},
		{[]string{"x", "xx", "xxx"}, 0},
		{[]string{"x", "xxx", "xx"}, 0},
		{[]string{"xxx", "xx", "x"}, 2},
		{[]string{"xxx", "xx", "xx"}, 1},
	}
	for i, test := range tests {
		result := strut.IndexOfShortest(test.s)
		if result != test.expected {
			t.Errorf("Test %d: IndexOfShortest({%s}) returned %d. Expected %d.", i, strut.Args(test.s), result, test.expected)
		}
	}
}
