package strut_test

import (
	"strings"
	"testing"

	"github.com/Zyl9393/strut"
)

func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		s        []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, "x"},
		{[]string{"x", "y"}, ""},
		{[]string{"xx", "xy"}, "x"},
		{[]string{"xy", "yy"}, ""},
		{[]string{"abcd", "ab", "abc"}, "ab"},
		{[]string{"abcd", "cd", "bcd"}, ""},
		{[]string{"abcd", "cd", "bcd"}, ""},
	}
	for i, test := range tests {
		result := strut.CommonPrefix(test.s)
		if result != test.expected {
			t.Errorf("Test %d: CommonPrefix(%s) returned `%s`. Expected `%s`.", i, strut.Args(test.s), result, test.expected)
		}
	}
}

func TestPrefixLen(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected int
	}{
		{"", "", 0},
		{"warning: foo", "warn: ", 4},
		{"warning: foo", "warning: ", 9},
		{"warn: foo", "warn: ", 6},
		{"warn: foo", "warning: ", 4},
		{"warn: foo", "", 0},
		{"", "warn: ", 0},
	}
	for i, test := range tests {
		result := strut.PrefixLen(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("Test #%d: PrefixLen(`%s`, `%s`) returned %d. Expected %d.", i+1, test.s, test.prefix, result, test.expected)
		}
	}
	for i, test := range tests {
		if !strings.HasPrefix(test.s, test.prefix[:strut.PrefixLen(test.s, test.prefix)]) {
			t.Errorf("Test #%d: strings.HasPrefix(`%s`, `%s`[:PrefixLen(`%s`, `%s`)]) returned false.", i+1, test.s, test.prefix, test.s, test.prefix)
		}
	}
}

func TestCommonSuffix(t *testing.T) {
	tests := []struct {
		s        []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, "x"},
		{[]string{"x", "y"}, ""},
		{[]string{"xx", "xy"}, ""},
		{[]string{"xy", "yy"}, "y"},
		{[]string{"abcd", "ab", "abc"}, ""},
		{[]string{"abcd", "cd", "bcd"}, "cd"},
	}
	for i, test := range tests {
		result := strut.CommonSuffix(test.s)
		if result != test.expected {
			t.Errorf("Test %d: CommonSuffix(%s) returned `%s`. Expected `%s`.", i, strut.Args(test.s), result, test.expected)
		}
	}
}

func TestSuffixLen(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected int
	}{
		{"", "", 0},
		{"warning: foo", "foo", 3},
		{"warning: foo", ": foo", 5},
		{"warn: foo", "foo", 3},
		{"warn: foo", ": foo", 5},
		{"warn: foo", "", 0},
		{"", "warn: ", 0},
	}
	for i, test := range tests {
		result := strut.SuffixLen(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("Test #%d: PrefixLen(`%s`, `%s`) returned %d. Expected %d.", i+1, test.s, test.prefix, result, test.expected)
		}
	}
	for i, test := range tests {
		if !strings.HasSuffix(test.s, test.prefix[:strut.SuffixLen(test.s, test.prefix)]) {
			t.Errorf("Test #%d: strings.HasSuffix(`%s`, `%s`[:SuffixLen(`%s`, `%s`)]) returned false.", i+1, test.s, test.prefix, test.s, test.prefix)
		}
	}
}
