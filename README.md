<p align="center">
  <img src="passgen-banner.jpg" alt="PassGen Banner" width="1024" />
</p>

# 🔐 passgen — Generate passwords from the terminal

![GitHub Release](https://img.shields.io/github/v/release/bilalbaraz/passgen)
![Codecov](https://img.shields.io/codecov/c/gh/bilalbaraz/passgen?logo=codecov&label=codecov)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/bilalbaraz/passgen/badge)](https://scorecard.dev/viewer/?uri=github.com/bilalbaraz/passgen)

<p align="center">
  <img src="demo.gif" alt="PassGen Demo" width="500" />
</p>

A small Go CLI that generates secure random passwords with flexible character sets and exclusions.

> Uses `crypto/rand` for cryptographically secure randomness. If no charset flags are provided, all sets are enabled by default.

## Installation

### Homebrew
```bash
brew tap bilalbaraz/tap
brew install passgen
```

### Linux amd64
```bash
wget -L https://github.com/bilalbaraz/passgen/releases/latest/download/passgen_linux_amd64.zip \
&& unzip passgen_linux_amd64.zip \
&& sudo mv passgen /usr/local/bin/ \
&& rm passgen_linux_amd64.zip
```

### Linux arm64
```bash
wget -L https://github.com/bilalbaraz/passgen/releases/latest/download/passgen_linux_arm64.zip \
&& unzip passgen_linux_arm64.zip \
&& sudo mv passgen /usr/local/bin/ \
&& rm passgen_linux_arm64.zip
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

# copy generated passwords to clipboard
passgen -len 20 -count 3 -copy

# render first password as QR code in terminal (requires -count 1)
passgen -len 20 -qr
```

## Command Surface
- **Length:** `-len <n>`
- **Char sets:** `-lower`, `-upper`, `-digits`, `-symbols`
- **Count:** `-count <n>`
- **Exclude chars:** `-exclude "..."`
- **No ambiguous:** `-no-ambiguous`
- **Copy to clipboard:** `-copy`
- **QR code:** `-qr` (only when `-count 1`)
- **Help:** `-h` / `-help`

## Configuration
No config file. All behavior is controlled via CLI flags.

## Notes
- Errors are printed to stderr and exit with code 1 on invalid input.
- `-copy` uses `pbcopy` (macOS), `clip` (Windows), or `wl-copy`/`xclip` (Linux).
