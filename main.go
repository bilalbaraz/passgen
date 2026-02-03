package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars  = "0123456789"
	symbolChars = "!@#$%^&*()-_=+[]{};:,.?/|~"
)

var ambiguousChars = []rune{'0', 'O', '1', 'l', 'I'}

const banner = " " +
	"                                       \n" +
	"  _ __   __ _ ___ ___  __ _  ___ _ __  \n" +
	" | '_ \\ / _` / __/ __|/ _` |/ _ \\ '_ \\ \n" +
	" | |_) | (_| \\__ \\__ \\ (_| |  __/ | | |\n" +
	" | .__/ \\__,_|___/___/\\__, |\\___|_| |_|\n" +
	" | |                   __/ |           \n" +
	" |_|                  |___/            \n"

func main() {
	length := flag.Int("len", 16, "password length")
	lower := flag.Bool("lower", false, "include lowercase letters")
	upper := flag.Bool("upper", false, "include uppercase letters")
	digits := flag.Bool("digits", false, "include digits")
	symbols := flag.Bool("symbols", false, "include symbols")
	count := flag.Int("count", 1, "number of passwords to generate")
	exclude := flag.String("exclude", "", "exclude specific characters")
	noAmbiguous := flag.Bool("no-ambiguous", false, "exclude ambiguous characters like 0 O 1 l I")
	copyOut := flag.Bool("copy", false, "copy generated passwords to clipboard")
	flag.Parse()

	fmt.Println(banner)

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

	passwords := make([]string, 0, *count)
	for i := 0; i < *count; i++ {
		pwd, err := generatePassword(*length, charset)
		if err != nil {
			fatal(err)
		}
		passwords = append(passwords, pwd)
		fmt.Println(pwd)
	}

	if *copyOut {
		if err := copyToClipboard(strings.Join(passwords, "\n")); err != nil {
			fatal(err)
		}
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

func copyToClipboard(text string) error {
	if text == "" {
		return errors.New("nothing to copy")
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else {
			return errors.New("no clipboard tool found (need wl-copy or xclip)")
		}
	}

	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("clipboard copy failed: %w", err)
	}
	return nil
}
