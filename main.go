package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars  = "0123456789"
	symbolChars = "!@#$%^&*()-_=+[]{};:,.?/|~"
)

var ambiguousChars = []rune{'0', 'O', '1', 'l', 'I'}

func main() {
	length := flag.Int("len", 16, "password length")
	lower := flag.Bool("lower", false, "include lowercase letters")
	upper := flag.Bool("upper", false, "include uppercase letters")
	digits := flag.Bool("digits", false, "include digits")
	symbols := flag.Bool("symbols", false, "include symbols")
	count := flag.Int("count", 1, "number of passwords to generate")
	exclude := flag.String("exclude", "", "exclude specific characters")
	noAmbiguous := flag.Bool("no-ambiguous", false, "exclude ambiguous characters like 0 O 1 l I")
	flag.Parse()

	if *length <= 0 {
		fatal(errors.New("-len must be greater than 0"))
	}
	if *count <= 0 {
		fatal(errors.New("-count must be greater than 0"))
	}

	if !charsetFlagsSpecified() {
		*lower = true
		*upper = true
		*digits = true
		*symbols = true
	}

	excludeSet := make(map[rune]bool)
	for _, r := range *exclude {
		excludeSet[r] = true
	}
	if *noAmbiguous {
		for _, r := range ambiguousChars {
			excludeSet[r] = true
		}
	}

	charset := buildCharset(*lower, *upper, *digits, *symbols, excludeSet)
	if len(charset) == 0 {
		fatal(errors.New("character set is empty after exclusions"))
	}

	for i := 0; i < *count; i++ {
		pwd, err := generatePassword(*length, charset)
		if err != nil {
			fatal(err)
		}
		fmt.Println(pwd)
	}
}

func charsetFlagsSpecified() bool {
	specified := false
	flag.CommandLine.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "lower", "upper", "digits", "symbols":
			specified = true
		}
	})
	return specified
}

func buildCharset(lower, upper, digits, symbols bool, excludeSet map[rune]bool) []rune {
	var b strings.Builder
	if lower {
		b.WriteString(lowerChars)
	}
	if upper {
		b.WriteString(upperChars)
	}
	if digits {
		b.WriteString(digitChars)
	}
	if symbols {
		b.WriteString(symbolChars)
	}

	var out []rune
	for _, r := range b.String() {
		if !excludeSet[r] {
			out = append(out, r)
		}
	}
	return out
}

func generatePassword(length int, charset []rune) (string, error) {
	var b strings.Builder
	b.Grow(length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b.WriteRune(charset[idx.Int64()])
	}
	return b.String(), nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
