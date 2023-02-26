package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		value        int
		firstIndex   int
		translations []string
		expected     string
	}{
		{0, 0, []string{}, "0"},
		{99, 100, []string{"cat", "dog", "mouse", "bird"}, "99"},
		{100, 100, []string{"cat", "dog", "mouse", "bird"}, "cat"},
		{101, 100, []string{"cat", "dog", "mouse", "bird"}, "dog"},
		{102, 100, []string{"cat", "dog", "mouse", "bird"}, "mouse"},
		{103, 100, []string{"cat", "dog", "mouse", "bird"}, "bird"},
		{104, 100, []string{"cat", "dog", "mouse", "bird"}, "104"},
	}
	for i, test := range tests {
		result := strut.Translate(test.value, test.firstIndex, test.translations...)
		if result != test.expected {
			t.Errorf("Test %d: Translate(%d, %d%s) returned %s. Expected %s.", i, test.value, test.firstIndex, strut.ArgsAfter(test.translations), result, test.expected)
		}
	}
}
