package but_test

import (
	"testing"

	"github.com/Zyl9393/strut/but"
)

func TestBoolToStringFunctions(t *testing.T) {
	tests := []struct {
		funcName      string
		f             func(bool) string
		expectedTrue  string
		expectedFalse string
	}{
		{"X", but.X, "x", ""},
		{"XU", but.XU, "X", ""},
		{"OZ", but.OZ, "1", "0"},
		{"Tf", but.Tf, "true", "false"},
		{"TF", but.TF, "True", "False"},
		{"TFU", but.TFU, "TRUE", "FALSE"},
		{"Y", but.Y, "y", "n"},
		{"YU", but.YU, "Y", "N"},
		{"Yn", but.Yn, "yes", "no"},
		{"YN", but.YN, "Yes", "No"},
		{"YNU", but.YNU, "YES", "NO"},
		{"Oo", but.Oo, "on", "off"},
		{"OO", but.OO, "On", "Off"},
		{"OOU", but.OOU, "ON", "OFF"},
	}
	for i, test := range tests {
		resultTrue := test.f(true)
		resultFalse := test.f(false)
		if resultTrue != test.expectedTrue {
			t.Errorf("Test %d: Ex(true) returned `%s`. Expected `%s`", i, resultTrue, test.expectedTrue)
		}
		if resultFalse != test.expectedFalse {
			t.Errorf("Test %d: Ex(false) returned `%s`. Expected `%s`", i, resultFalse, test.expectedFalse)
		}
	}
}
