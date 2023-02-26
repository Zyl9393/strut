package strut

import (
	"path/filepath"
	"strings"

	"github.com/rivo/uniseg"
)

// Filepath is a shorthand for filepath.Clean(strings.ReplaceAll(path, `\`, `/`)).
func Filepath(path string) string {
	return filepath.Clean(strings.ReplaceAll(path, `\`, `/`))
}

// Redact is a shorthand for RedactMinMax(secret, 12, 2), which has a minRedact parameter of 12 set according to NIST publications
// regarding recommended minimum password length and a maxReveal parameter of 2, never letting more than the first and the last
// character of secret be revealed.
func Redact(secret string) string {
	return RedactMinMax(secret, 12, 2)
}

// RedactMinMax returns secret with up to minRedact characters redacted before allowing at most maxReveal characters to not be redacted.
// The amount of characters not redacted is distributed between the leading and the trailing characters of secret,
// whereby leading characters take precedence in case the amount of characters to be left revealed is odd.
//
// While this function is agnostic to the contents of secret, it is not agnostic to its length, i.e. this
// function does not execute in constant time and thus is not cryptographically secure to be called multiple times.
// It is intended to be called only once per secret per program execution, e.g. for logging during startup of a service.
func RedactMinMax(secret string, minRedact int, maxReveal int) string {
	if minRedact < 0 {
		panic("minRedact < 0")
	}
	if maxReveal < 0 {
		panic("maxReveal < 0")
	}
	scan := secret
	state := -1
	graphemeClusterIndices := make([]int, 0, len(secret))
	for scan != "" {
		graphemeClusterIndices = append(graphemeClusterIndices, len(secret)-len(scan))
		_, scan, _, state = uniseg.FirstGraphemeClusterInString(scan, state)
	}
	if minRedact > len(graphemeClusterIndices) {
		minRedact = len(graphemeClusterIndices)
	}
	revealCount := maxReveal
	if revealCount > len(graphemeClusterIndices)-minRedact {
		revealCount = len(graphemeClusterIndices) - minRedact
	}
	revealRight := revealCount / 2
	revealLeft := revealCount - revealRight
	redactCount := len(graphemeClusterIndices) - revealLeft - revealRight
	graphemeClusterIndices = append(graphemeClusterIndices, len(secret))
	if revealCount > 0 && redactCount > 3 {
		return secret[:graphemeClusterIndices[revealLeft]] + "..." + secret[graphemeClusterIndices[len(graphemeClusterIndices)-1-revealRight]:]
	}
	return secret[:graphemeClusterIndices[revealLeft]] + strings.Repeat("*", redactCount) + secret[graphemeClusterIndices[len(graphemeClusterIndices)-1-revealRight]:]
}
