package passgen

import "testing"

func TestGeneratePassword_LengthAndCharset(t *testing.T) {
	charset := []rune{'a', 'b', 'c', '1', '2'}
	got, err := GeneratePassword(32, charset)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len([]rune(got)) != 32 {
		t.Fatalf("expected length 32, got %d", len([]rune(got)))
	}
	for _, r := range got {
		found := false
		for _, allowed := range charset {
			if r == allowed {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("unexpected rune %q not in charset", r)
		}
	}
}

func TestGeneratePassword_EmptyCharset(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for empty charset, got none")
		}
	}()
	_, _ = GeneratePassword(1, nil)
}

func TestGeneratePassword_ZeroLength(t *testing.T) {
	got, err := GeneratePassword(0, []rune{'x'})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}
