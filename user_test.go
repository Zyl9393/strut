package strut_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/Zyl9393/strut"
	"github.com/stretchr/testify/assert"
)

func TestFilepath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{``, `.`},
		{`D:/Foo/Bar`, `D:/Foo/Bar`},
		{`/foo/bar`, `/foo/bar`},
	}
	for i, test := range tests {
		path := test.path
		expected := test.expected
		if filepath.Separator == '/' {
			path = strings.ReplaceAll(test.path, `/`, `\`)
		}
		if filepath.Separator == '\\' {
			expected = strings.ReplaceAll(test.expected, `/`, `\`)
		}
		result := strut.Filepath(path)
		if result != expected {
			t.Errorf("Test %d: Filepath(`%s`) returned `%s`. Expected `%s`.", i, path, result, expected)
		}
	}
}

func TestRedact(t *testing.T) {
	tests := []struct {
		secret   string
		expected string
	}{
		{"abcd1234efgh", "************"},
		{"Xabcd1234efgh", "X..."},
		{"Xabcd1234efghY", "X...Y"},
		{"abcd-efgh-ijkl-mnop-qrst-uvwx", "a...x"},
	}
	for i, test := range tests {
		result := strut.Redact(test.secret)
		if result != test.expected {
			t.Errorf("Test %d: Redact(`%s`) returned `%s`. Expected `%s`.", i, test.secret, result, test.expected)
		}
	}
}

func TestRedactMinMax(t *testing.T) {
	tests := []struct {
		secret    string
		minRedact int
		maxReveal int
		expected  string
	}{
		{"", 0, 0, ""},
		{"", 1, 0, ""},
		{"", 0, 1, ""},
		{"", 1, 1, ""},
		{"a", 0, 0, "*"},
		{"a", 1, 0, "*"},
		{"a", 10, 0, "*"},
		{"a", 0, 1, "a"},
		{"a", 1, 1, "*"},
		{"a", 10, 1, "*"},
		{"a", 0, 10, "a"},
		{"a", 1, 10, "*"},
		{"a", 10, 10, "*"},
		{"abcdef", 0, 0, "******"},
		{"abc", 2, 2, "a**"},
		{"abc", 1, 2, "a*c"},
		{"abcd-efgh-ijkl-mnop-qrst-uvwx", 0, 14, "abcd-ef...st-uvwx"},
	}
	for i, test := range tests {
		result := strut.RedactMinMax(test.secret, test.minRedact, test.maxReveal)
		if result != test.expected {
			t.Errorf("Test %d: RedactMinMax(`%s`, %d, %d) returned `%s`. Expected `%s`.", i, test.secret, test.minRedact, test.maxReveal, result, test.expected)
		}
	}
	assert.PanicsWithValue(t, "minRedact < 0", func() { strut.RedactMinMax("", -1, 0) })
	assert.PanicsWithValue(t, "maxReveal < 0", func() { strut.RedactMinMax("", 0, -1) })
}
