package strut

import (
	"fmt"
	"strings"
)

// Arg is a shorthand for fmt.Sprintf("%#q", arg) which returns `arg` as a backquoted string where possible and sensible;
// e.g. for arg = "\\n" returns `\n`, but for arg = "\n" returns "\\n".
func Arg(arg string) string {
	return fmt.Sprintf("%#q", arg)
}

// Args returns a string joining the "%#q"-formatted string of every string in args with ", ".
// E.g. Args([]string{"Hello\nWorld!", "D:\\Foo"}) returns a string which prints as "Hello\nWorld", `D:\Foo`.
func Args(args []string) string {
	var sb strings.Builder
	writeArgs(args, &sb, false, false)
	return sb.String()
}

// ArgsBefore is like Args, but also adds the suffix ", " at the end of the returned string if len(args) > 0.
// I.e. args appear *before* some other arguments that follow.
func ArgsBefore(args []string) string {
	var sb strings.Builder
	writeArgs(args, &sb, false, true)
	return sb.String()
}

// ArgsAfter is like Args, but also adds the prefix ", " at the start of the returned string if len(args) > 0.
// I.e. args appear *after* some other preceding arguments.
func ArgsAfter(args []string) string {
	var sb strings.Builder
	writeArgs(args, &sb, true, false)
	return sb.String()
}

// ArgsBetween is like Args, but also adds the prefix ", " at the start and the suffix ", " at the end
// of the returned string if len(args) > 0.
// I.e. args appear *between* some other surrounding arguments.
func ArgsBetween(args []string) string {
	var sb strings.Builder
	writeArgs(args, &sb, true, true)
	return sb.String()
}

func writeArgs(args []string, sb *strings.Builder, prefixComma bool, postfixComma bool) {
	if len(args) == 0 {
		return
	}
	educatedGuess := sb.Len()
	if prefixComma {
		educatedGuess += 2
	}
	if postfixComma {
		educatedGuess += 2
	}
	for _, arg := range args {
		educatedGuess += len(arg)
	}
	educatedGuess += educatedGuess>>3 + len(args)*4
	sb.Grow(educatedGuess)
	if prefixComma {
		sb.WriteString(", ")
	}
	for _, arg := range args[:len(args)-1] {
		sb.WriteString(fmt.Sprintf("%#q, ", arg))
	}
	sb.WriteString(fmt.Sprintf("%#q", args[len(args)-1]))
	if postfixComma {
		sb.WriteString(", ")
	}
}
