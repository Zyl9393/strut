package strut_test

import (
	"reflect"
	"testing"

	"github.com/Zyl9393/strut"
)

func TestSplitTwinUnescape(t *testing.T) {
	tests := []struct {
		s        string
		sep      string
		expected []string
	}{
		{"", ";", []string{""}},
		{".", ";", []string{"."}},
		{";", ";", []string{"", ""}},
		{"", "longsep", []string{""}},
		{"foo", "longsep", []string{"foo"}},
		{"FOO=bar", ";", []string{"FOO=bar"}},
		{"FOO=b;;r", ";", []string{"FOO=b;r"}},
		{"FOO=bar;ENV=var", ";", []string{"FOO=bar", "ENV=var"}},
		{"FOO=b;;r;ENV=v;;r", ";", []string{"FOO=b;r", "ENV=v;r"}},
		{"FOO=b;;r;ENV=v;;r;", ";", []string{"FOO=b;r", "ENV=v;r", ""}},
		{"FOO=bar;", ";", []string{"FOO=bar", ""}},
		{";;;", ";", []string{";", ""}},
		{";;;", ";;", []string{"", ";"}},
		{";;;", ";;;", []string{"", ""}},
		{";;;", ";;;;", []string{";;;"}},
		{"what do you mean 'haha' won't half this string?", "ha", []string{"w", "t do you mean 'ha' won't ", "lf this string?"}},
	}
	for i, test := range tests {
		result := strut.SplitTwinUnescape(test.s, test.sep)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: SplitTwinUnescape(`%s`, `%s`) returned {%v}. Expected: {%v}.", i, test.s, test.sep, strut.Args(result), strut.Args(test.expected))
		}
	}
}
