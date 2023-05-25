# strut (**str**ing **ut**ilities)

Provides additional desirable functions for string processing not present in package `strings`.

# Function overview

## Arguments *(args.go)*

These functions are useful for printing string slice contents and variadic string arguments.

| Function     | Explanation |
| ------------ | ----------- |
| Arg          | Returns `arg` as a backquoted string where possible and sensible; e.g. for `arg = "\\n"` returns \`\n\`, but for `arg = "\n"` returns "\\\\n". |
| Args         | Like `Arg()`, but for `[]string`, joined together with `", "`. |
| ArgsBefore   | Like `Args()`, but appends an extra `", "` if at least one string was passed. |
| ArgsAfter    | Like `Args()`, but prepends an extra `", "` if at least one string was passed. |
| ArgsBetween  | Like `Args()`, but prepends and appends an extra `", "` if at least one string was passed. |

## Bool to string _(bool.go)_

| Function            | Explanation |
| ------------------- | ----------- |
| NewBoolToStringFunc | Returns a function which takes a bool parameter and returns `trueString` if the parameter is true and `falseString` if the parameter is false. |

The functions in the following table are contained in sub-package `but`:

| Function | Explanation |
| -------- | ----------- |
| X        | Returns `"x"` if bool is true and `""` if bool is false. |
| XU       | Returns `"X"` if bool is true and `""` if bool is false. |
| OZ       | Returns `"1"` if bool is true and `"0"` if bool is false. |
| Tf       | Returns `"true"` if bool is true and `"false"` if bool is false. |
| TF       | Returns `"True"` if bool is true and `"False"` if bool is false. |
| TFU      | Returns `"TRUE"` if bool is true and `"FALSE"` if bool is false. |
| Y        | Returns `"y"` if bool is true and `"n"` if bool is false. |
| YU       | Returns `"Y"` if bool is true and `"N"` if bool is false. |
| Yn       | Returns `"yes"` if bool is true and `"no"` if bool is false. |
| YN       | Returns `"Yes"` if bool is true and `"No"` if bool is false. |
| YNU      | Returns `"YES"` if bool is true and `"NO"` if bool is false. |
| Oo       | Returns `"on"` if bool is true and `"off"` if bool is false. |
| OO       | Returns `"On"` if bool is true and `"Off"` if bool is false. |
| OOU      | Returns `"ON"` if bool is true and `"OFF"` if bool is false. |

## Find longest/shortest *(find.go)*

| Function        | Explanation |
| --------------- | ----------- |
| IndexOfLongest  | Returns index of first string in passed `[]string` for which no shorter string exists. |
| IndexOfShortest | Returns index of first string in passed `[]string` for which no longer string exists. |

## Join *(join.go)*

| Function  | Explanation |
| --------- | ----------- |
| JoinUnary | Like `strings.Join()`, except that all leading and trailing occurrences of `sep` within `elems` are trimmed where they are going to be joined with `sep` beforehand. |
| Slash     | Shorthand for `JoinUnary(elems, "/")` which accepts a variadic argument list. |

## Misc *(replace.go, translate.go)*

| Function        | Explanation |
| --------------- | ----------- |
| RangeReplaceAll | Returns a copy of a `[]string` where `strings.ReplaceAll()` was called for every string. |
| Translate       | Map a non-zero-indexed integer constant to a list of translation terms in a `[]string`. |

## Pre- and postfix analysis *(pfix.go)*

| Function     | Explanation |
| ------------ | ----------- |
| CommonPrefix | Returns the longest prefix common to all strings in passed `[]string`. |
| PrefixLen    | Returns the length of the longest common prefix of two strings. |
| CommonSuffix | Returns the longest suffix common to all strings in passed `[]string`. |
| SuffixLen    | Returns the length of the longest common suffix of two strings. |

## Process user input *(user.go)*

| Function     | Explanation |
| ------------ | ----------- |
| Filepath     | Like `filepath.Clean()`, but treats both `/` and `\` as a file path element separator regardless of OS. |
| RedactMinMax | Redact up to `minRedact` runes before leaving no more than `minReveal` runes non-redacted in input string. |
| Redact       | Shorthand for `RedactMinMax(secret, 12, 2)`. E.g. `abcd-1234-efgh` becomes `a...h`. |

## Slicing and indexing *(index.go)*

| Function      | Explanation |
| ------------- | ----------- |
| IndexRune     | Returns byte index for given rune index. |
| RuneIndex     | Returns index of rune at given byte index. |
| RangeLen      | Returns the amount of iterations which take place when ranging over given string. |
| SliceRunes    | Slice a string with rune indices instead of byte indices. |
| SliceClusters | Slice a string with [grapheme cluster](https://github.com/rivo/uniseg#grapheme-clusters) indices instead of byte or rune indices. |

## Splitting *(split.go)*

| Function          | Explanation |
| ----------------- | ----------- |
| SplitTwinUnescape | Split string around single occurrences of separator, reducing double occurrences to a single occurrence. |

## Substrings *(substrings.go)*

| Function                | Explanation |
| ----------------------- | ----------- |
| IterateSubstrings       | Invoke a callback-function for every substring having between `minSubstringRuneCount` and `maxSubstringRuneCount` runes. |
| IterateSubstringsUnique | Like `IterateSubstrings()`, but never calls the callback-function more than once for equal substrings. |

## Trim *(trim.go)*

| Function        | Explanation |
| --------------- | ----------- |
| TrimPrefixAll   | Like `strings.TrimPrefix()`, but repeats until all occurrences of the prefix are gone. |
| TrimPrefixN     | Like `TrimPrefixAll`, but limited to trimming the prefix no more than `n` times. |
| TrimSuffixAll   | Like `strings.TrimSuffix()`, but repeats until all occurrences of the suffix are gone. |
| TrimSuffixN     | Like `TrimSuffixAll`, but limited to trimming the suffix no more than `n` times. |
| TrimSequenceAll | Shorthand for `TrimSuffixAll(TrimSequenceAll(s))`. |
