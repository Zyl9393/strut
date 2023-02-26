package strut

import (
	"strings"
)

// RangeReplaceAll returns a new slice which is a copy of ss with every contained string having all non-overlapping instances of old replaced by new.
func RangeReplaceAll(ss []string, old string, new string) []string {
	newSlice := make([]string, len(ss))
	for i, s := range ss {
		newSlice[i] = strings.ReplaceAll(s, old, new)
	}
	return newSlice
}
