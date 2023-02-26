package strut_test

import (
	"strings"
	"testing"

	"github.com/Zyl9393/strut"
)

func TestSlash(t *testing.T) {
	tests := []struct {
		elements []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab"}, "ab"},
		{[]string{"", ""}, "/"},
		{[]string{"///", "///"}, "/"},
		{[]string{"", "", ""}, "//"},
		{[]string{"///", "///", "///"}, "//"},
		{[]string{"///a///"}, "///a///"},
		{[]string{"///ab///"}, "///ab///"},
		{[]string{"a", "x"}, "a/x"},
		{[]string{"/a", "x/"}, "/a/x/"},
		{[]string{"a/", "/x"}, "a/x"},
		{[]string{"a", "b", "c"}, "a/b/c"},
		{[]string{"/a", "b", "c/"}, "/a/b/c/"},
		{[]string{"///a///", "///b///", "///c///"}, "///a/b/c///"},
		{[]string{"abcdef", "ghij", "klmnopqrst"}, "abcdef/ghij/klmnopqrst"},
		{[]string{"abcdef//", "//ghij//", "//klmnopqrst"}, "abcdef/ghij/klmnopqrst"},
		{[]string{"//abcdef//", "//ghij//", "//klmnopqrst//"}, "//abcdef/ghij/klmnopqrst//"},
	}
	for i, test := range tests {
		result := strut.Slash(test.elements...)
		if result != test.expected {
			t.Errorf("Test %d: Slash(%s) returned `%s`. Expected `%s`.", i, strut.Args(test.elements), result, test.expected)
		}
	}
}

func TestJoinUnary(t *testing.T) {
	tests := []struct {
		elements []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab"}, "ab"},
		{[]string{"", ""}, "/"},
		{[]string{"///", "///"}, "/"},
		{[]string{"", "", ""}, "//"},
		{[]string{"///", "///", "///"}, "//"},
		{[]string{"///a///"}, "///a///"},
		{[]string{"///ab///"}, "///ab///"},
		{[]string{"a", "x"}, "a/x"},
		{[]string{"/a", "x/"}, "/a/x/"},
		{[]string{"a/", "/x"}, "a/x"},
		{[]string{"a", "b", "c"}, "a/b/c"},
		{[]string{"/a", "b", "c/"}, "/a/b/c/"},
		{[]string{"///a///", "///b///", "///c///"}, "///a/b/c///"},
		{[]string{"abcdef", "ghij", "klmnopqrst"}, "abcdef/ghij/klmnopqrst"},
		{[]string{"abcdef//", "//ghij//", "//klmnopqrst"}, "abcdef/ghij/klmnopqrst"},
		{[]string{"//abcdef//", "//ghij//", "//klmnopqrst//"}, "//abcdef/ghij/klmnopqrst//"},
	}
	seps := []string{"/", "12", "123", "121"}
	for j, sep := range seps {
		for i, test := range tests {
			elements := strut.RangeReplaceAll(test.elements, "/", sep)
			expected := strings.ReplaceAll(test.expected, "/", sep)
			result := strut.JoinUnary(elements, sep)
			if result != expected {
				t.Errorf("Test 0.%d.%d: JoinUnary({%s}, `%s`) returned `%s`. Expected `%s`.", j, i, strut.Args(test.elements), sep, result, expected)
			}
		}
	}
	tests2 := []struct {
		elements []string
		sep      string
		expected string
	}{
		{[]string{"//2121212//", "//212//", "//21212//"}, "12", "//2/2/21212//"},
		{[]string{"//2121212//", "//212//", "//21212//"}, "121", "//2121212/212/21212//"},
		{[]string{"//2121212//", "//212//", "//21212//"}, "212", "//2121//12//"},
		{[]string{"//2121212//", "//212//", "//21212//"}, "21", "//2121212/2/2//"},
	}
	for i, test := range tests2 {
		elements := strut.RangeReplaceAll(test.elements, "/", test.sep)
		expected := strings.ReplaceAll(test.expected, "/", test.sep)
		result := strut.JoinUnary(elements, test.sep)
		if result != expected {
			t.Errorf("Test 1.%d: JoinUnary({%s}, `%s`) returned `%s`. Expected `%s`.", i, strut.Args(test.elements), test.sep, result, expected)
		}
	}
}
