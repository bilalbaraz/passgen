package passgen

import "strings"

const (
	LowerChars  = "abcdefghijklmnopqrstuvwxyz"
	UpperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DigitChars  = "0123456789"
	SymbolChars = "!@#$%^&*()-_=+[]{};:,.?/|~"
)

var AmbiguousChars = []rune{'0', 'O', '1', 'l', 'I'}

func BuildCharset(lower, upper, digits, symbols bool, excludeSet map[rune]bool) []rune {
	var b strings.Builder
	if lower {
		b.WriteString(LowerChars)
	}
	if upper {
		b.WriteString(UpperChars)
	}
	if digits {
		b.WriteString(DigitChars)
	}
	if symbols {
		b.WriteString(SymbolChars)
	}

	var out []rune
	for _, r := range b.String() {
		if !excludeSet[r] {
			out = append(out, r)
		}
	}
	return out
}
