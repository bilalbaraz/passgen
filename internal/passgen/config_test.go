package passgen

import "testing"

func TestConfigValidate(t *testing.T) {
	t.Run("invalid length", func(t *testing.T) {
		cfg := Config{Length: 0, Count: 1}
		if err := cfg.Validate(); err == nil || err.Error() != "-len must be greater than 0" {
			t.Fatalf("expected length error, got %v", err)
		}
	})

	t.Run("invalid count", func(t *testing.T) {
		cfg := Config{Length: 8, Count: 0}
		if err := cfg.Validate(); err == nil || err.Error() != "-count must be greater than 0" {
			t.Fatalf("expected count error, got %v", err)
		}
	})

	t.Run("qr count too large", func(t *testing.T) {
		cfg := Config{Length: 8, Count: 2, QR: true}
		if err := cfg.Validate(); err == nil || err.Error() != "-qr supports only -count=1 to avoid terminal clutter" {
			t.Fatalf("expected qr count error, got %v", err)
		}
	})

	t.Run("valid", func(t *testing.T) {
		cfg := Config{Length: 12, Count: 1}
		if err := cfg.Validate(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestConfigExcludeSet(t *testing.T) {
	t.Run("exclude string only", func(t *testing.T) {
		cfg := Config{Exclude: "aZ0"}
		set := cfg.ExcludeSet()
		for _, r := range []rune{'a', 'Z', '0'} {
			if !set[r] {
				t.Fatalf("expected rune %q to be excluded", r)
			}
		}
		if set['b'] {
			t.Fatal("did not expect unrelated rune to be excluded")
		}
	})

	t.Run("no ambiguous adds default set", func(t *testing.T) {
		cfg := Config{Exclude: "x", NoAmbiguous: true}
		set := cfg.ExcludeSet()
		if !set['x'] {
			t.Fatal("expected exclude string to be included")
		}
		for _, r := range AmbiguousChars {
			if !set[r] {
				t.Fatalf("expected ambiguous rune %q to be excluded", r)
			}
		}
	})

	t.Run("empty config yields empty set", func(t *testing.T) {
		cfg := Config{}
		set := cfg.ExcludeSet()
		if len(set) != 0 {
			t.Fatalf("expected empty exclude set, got %d entries", len(set))
		}
	})
}
