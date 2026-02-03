# 🔐 passgen — Generate passwords from the terminal

<p align="center">
  <img src="demo.gif" alt="PassGen Demo" width="500" />
</p>

A small Go CLI that generates secure random passwords with flexible character sets and exclusions.

> Uses `crypto/rand` for cryptographically secure randomness. If no charset flags are provided, all sets are enabled by default.

## Installation (Homebrew)
```bash
brew tap bilalbaraz/tap
brew install passgen
```

## Quickstart
```bash

# basic (defaults to all character sets)
passgen -len 16 -count 3

# lower/upper/digits only
passgen -len 24 -count 5 -lower -upper -digits

# symbols only, excluding "@$"
passgen -len 20 -symbols -exclude "@$"

# exclude ambiguous characters (0 O 1 l I)
passgen -len 16 -no-ambiguous
```

## Command Surface
- **Length:** `-len <n>`
- **Char sets:** `-lower`, `-upper`, `-digits`, `-symbols`
- **Count:** `-count <n>`
- **Exclude chars:** `-exclude "..."`
- **No ambiguous:** `-no-ambiguous`
- **Help:** `-h` / `-help`

## Configuration
No config file. All behavior is controlled via CLI flags.

## Notes
- Errors are printed to stderr and exit with code 1 on invalid input.
