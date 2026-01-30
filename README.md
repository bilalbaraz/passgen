# passgen

Command-line password generator written in Go.

## Features
- Length control with `-len`
- Character sets: `-lower`, `-upper`, `-digits`, `-symbols`
- Generate multiple passwords with `-count`
- Exclude specific characters with `-exclude`
- Remove ambiguous chars (0 O 1 l I) with `-no-ambiguous`

## Install

### Go toolchain

```
go install ./...
```

### Build

```
go build -o passgen .
```

## Usage

Basic (defaults to all character sets if none are specified):

```
go run . -len 16 -count 3
```

Only lower/upper/digits:

```
go run . -len 24 -count 5 -lower -upper -digits
```

Symbols only, excluding "@$":

```
go run . -len 20 -symbols -exclude "@$"
```

Exclude ambiguous characters:

```
go run . -len 16 -no-ambiguous
```

Help:

```
go run . -h
```

## CI/CD

- CI runs on tag push matching `v*`.
- Builds the binary and uploads it as a GitHub Release asset.

## Notes
- Uses `crypto/rand` for cryptographically secure randomness.
- If no charset flags are provided, all are enabled by default.
