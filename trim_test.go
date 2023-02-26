package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
)

func TestTrimPrefixAll(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"ax", "x", "ax"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"xa", "", "xa"},
		{"xa", "x", "a"},
		{"xxa", "", "xxa"},
		{"xxa", "x", "a"},
		{"xxxa", "", "xxxa"},
		{"xxxa", "x", "a"},
		{"xxxa", "xx", "xa"},
		{"xxxxa", "xx", "a"},
		{"xxxxxa", "xx", "xa"},
	}
	for i, test := range tests {
		result := strut.TrimPrefixAll(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("Test %d: TrimPrefixAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.prefix, result, test.expected)
		}
	}
}

func TestTrimPrefixN(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		n        int
		expected string
	}{
		{"", "", 1, ""},
		{"a", "x", 1, "a"},
		{"ax", "", 1, "ax"},
		{"ax", "x", 1, "ax"},
		{"", "x", 1, ""},
		{"x", "", 1, "x"},
		{"x", "x", 1, ""},
		{"xa", "", 1, "xa"},
		{"xa", "x", 1, "a"},
		{"xxa", "", 1, "xxa"},
		{"xxa", "x", 1, "xa"},
		{"xxxa", "", 1, "xxxa"},
		{"xxxa", "x", 1, "xxa"},
		{"xa", "", 2, "xa"},
		{"xa", "x", 2, "a"},
		{"xxa", "", 2, "xxa"},
		{"xxa", "x", 2, "a"},
		{"xxxa", "", 2, "xxxa"},
		{"xxxa", "x", 2, "xa"},
		{"xxxa", "xx", 1, "xa"},
		{"xxxxa", "xx", 1, "xxa"},
		{"xxxxa", "xx", 2, "a"},
		{"xxxxxa", "xx", 0, "xxxxxa"},
		{"xxxxxa", "xx", 1, "xxxa"},
		{"xxxxxa", "xx", 2, "xa"},
		{"xxxxxa", "xx", 3, "xa"},
		{"xxxxxa", "xx", -1, "xa"},
	}
	for i, test := range tests {
		result := strut.TrimPrefixN(test.s, test.prefix, test.n)
		if result != test.expected {
			t.Errorf("Test %d: TrimPrefixAll(`%s`, `%s`, %d) returned `%s`. Expected `%s`.", i, test.s, test.prefix, test.n, result, test.expected)
		}
	}
}

func TestTrimSuffixAll(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"xa", "x", "xa"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"axx", "", "axx"},
		{"axx", "x", "a"},
		{"axxx", "", "axxx"},
		{"axxx", "x", "a"},
		{"axxx", "xx", "ax"},
		{"axxxx", "xx", "a"},
		{"axxxxx", "xx", "ax"},
	}
	for i, test := range tests {
		result := strut.TrimSuffixAll(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("Test %d: TrimSuffixAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.suffix, result, test.expected)
		}
	}
}

func TestTrimSuffixN(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		n        int
		expected string
	}{
		{"", "", 1, ""},
		{"a", "x", 1, "a"},
		{"xa", "", 1, "xa"},
		{"xa", "x", 1, "xa"},
		{"", "x", 1, ""},
		{"x", "", 1, "x"},
		{"x", "x", 1, ""},
		{"ax", "", 1, "ax"},
		{"ax", "x", 1, "a"},
		{"axx", "", 1, "axx"},
		{"axx", "x", 1, "ax"},
		{"axxx", "", 1, "axxx"},
		{"axxx", "x", 1, "axx"},
		{"ax", "", 2, "ax"},
		{"ax", "x", 2, "a"},
		{"axx", "", 2, "axx"},
		{"axx", "x", 2, "a"},
		{"axxx", "", 2, "axxx"},
		{"axxx", "x", 2, "ax"},
		{"axxx", "xx", 1, "ax"},
		{"axxxx", "xx", 1, "axx"},
		{"axxxx", "xx", 2, "a"},
		{"axxxxx", "xx", 0, "axxxxx"},
		{"axxxxx", "xx", 1, "axxx"},
		{"axxxxx", "xx", 2, "ax"},
		{"axxxxx", "xx", 3, "ax"},
		{"axxxxx", "xx", -1, "ax"},
	}
	for i, test := range tests {
		result := strut.TrimSuffixN(test.s, test.suffix, test.n)
		if result != test.expected {
			t.Errorf("Test %d: TrimSuffixN(`%s`, `%s`, %d) returned `%s`. Expected `%s`.", i, test.s, test.suffix, test.n, result, test.expected)
		}
	}
}

func TestTrimSequenceAll(t *testing.T) {
	tests := []struct {
		s        string
		sequence string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"xa", "", "xa"},
		{"xa", "x", "a"},
		{"xxa", "", "xxa"},
		{"xxa", "x", "a"},
		{"xxxa", "", "xxxa"},
		{"xxxa", "x", "a"},
		{"xxxa", "xx", "xa"},
		{"xxxxa", "xx", "a"},
		{"xxxxxa", "xx", "xa"},

		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"xa", "x", "a"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"axx", "", "axx"},
		{"axx", "x", "a"},
		{"axxx", "", "axxx"},
		{"axxx", "x", "a"},
		{"axxx", "xx", "ax"},
		{"axxxx", "xx", "a"},
		{"axxxxx", "xx", "ax"},

		{"xxxaxxxxx", "xx", "xax"},
		{"xxxxxaxxx", "xx", "xax"},
		{"xxxxaxx", "xx", "a"},
		{"xx", "xx", ""},
	}
	for i, test := range tests {
		result := strut.TrimSequenceAll(test.s, test.sequence)
		if result != test.expected {
			t.Errorf("Test %d: TrimSequenceAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.sequence, result, test.expected)
		}
	}
}
