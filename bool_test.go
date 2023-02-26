package strut_test

import (
	"testing"

	"github.com/Zyl9393/strut"
)

func TestNewBoolToStringFunc(t *testing.T) {
	bts := strut.NewBoolToStringFunc("up", "down")
	resultTrue := bts(true)
	resultFalse := bts(false)
	if resultTrue != "up" {
		t.Errorf("bts(true) returned `%s`. Expected `up`.", resultTrue)
	}
	if resultFalse != "down" {
		t.Errorf("bts(true) returned `%s`. Expected `down`.", resultFalse)
	}
}
