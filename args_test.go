package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
)

func TestArg(t *testing.T) {
	tests := []struct {
		arg      string
		expected string
	}{
		{"", "``"},
		{"x", "`x`"},
		{"x\nz", "\"x\\nz\""},
	}
	for i, test := range tests {
		result := strut.Arg(test.arg)
		if result != test.expected {
			t.Errorf("Test %d: Arg(%s) returned %s. Expected %s.", i, test.arg, result, test.expected)
		}
	}
}

func TestArgs(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, "`x`"},
		{[]string{"x", "y"}, "`x`, `y`"},
		{[]string{"x\nz", "y"}, "\"x\\nz\", `y`"},
	}
	for i, test := range tests {
		result := strut.Args(test.args)
		if result != test.expected {
			t.Errorf("Test %d: Args(%v) returned %s. Expected %s.", i, test.args, result, test.expected)
		}
	}
}

func TestArgsBefore(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, "`x`, "},
		{[]string{"x", "y"}, "`x`, `y`, "},
		{[]string{"x\nz", "y"}, "\"x\\nz\", `y`, "},
	}
	for i, test := range tests {
		result := strut.ArgsBefore(test.args)
		if result != test.expected {
			t.Errorf("Test %d: ArgsBefore(%v) returned %s. Expected %s.", i, test.args, result, test.expected)
		}
	}
}

func TestArgsAfter(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, ", `x`"},
		{[]string{"x", "y"}, ", `x`, `y`"},
		{[]string{"x\nz", "y"}, ", \"x\\nz\", `y`"},
	}
	for i, test := range tests {
		result := strut.ArgsAfter(test.args)
		if result != test.expected {
			t.Errorf("Test %d: ArgsAfter(%v) returned %s. Expected %s.", i, test.args, result, test.expected)
		}
	}
}

func TestArgsBetween(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"x"}, ", `x`, "},
		{[]string{"x", "y"}, ", `x`, `y`, "},
		{[]string{"x\nz", "y"}, ", \"x\\nz\", `y`, "},
	}
	for i, test := range tests {
		result := strut.ArgsBetween(test.args)
		if result != test.expected {
			t.Errorf("Test %d: ArgsBetween(%v) returned %s. Expected %s.", i, test.args, result, test.expected)
		}
	}
}
