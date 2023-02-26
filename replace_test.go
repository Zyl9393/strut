package strut_test

import (
	"reflect"
	"testing"

	"github.com/Zyl9393/strut"
)

func TestRangeReplaceAll(t *testing.T) {
	tests := []struct {
		s        []string
		old      string
		new      string
		expected []string
	}{
		{nil, "", "", []string{}},
		{[]string{}, "", "", []string{}},
		{[]string{"x"}, "x", "y", []string{"y"}},
		{[]string{"in", "if"}, "i", "o", []string{"on", "of"}},
	}
	for i, test := range tests {
		result := strut.RangeReplaceAll(test.s, test.old, test.new)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: RangeReplaceAll(%v, `%s`, `%s`) returned %v. Expected %v.", i, test.s, test.old, test.new, result, test.expected)
		}
	}
}
