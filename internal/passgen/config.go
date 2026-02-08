package passgen

import "errors"

type Config struct {
	Length      int
	Lower       bool
	Upper       bool
	Digits      bool
	Symbols     bool
	Count       int
	Exclude     string
	NoAmbiguous bool
	Copy        bool
	QR          bool
}

func (c *Config) Validate() error {
	if c.Length <= 0 {
		return errors.New("-len must be greater than 0")
	}
	if c.Count <= 0 {
		return errors.New("-count must be greater than 0")
	}
	if c.QR && c.Count > 1 {
		return errors.New("-qr supports only -count=1 to avoid terminal clutter")
	}
	return nil
}

func (c *Config) ExcludeSet() map[rune]bool {
	excludeSet := make(map[rune]bool)
	for _, r := range c.Exclude {
		excludeSet[r] = true
	}
	if c.NoAmbiguous {
		for _, r := range AmbiguousChars {
			excludeSet[r] = true
		}
	}
	return excludeSet
}
