package passgen

import "testing"

func TestBuildCharset_AllGroups(t *testing.T) {
	exclude := map[rune]bool{}
	got := BuildCharset(true, true, true, true, exclude)
	expected := []rune(LowerChars + UpperChars + DigitChars + SymbolChars)

	if len(got) != len(expected) {
		t.Fatalf("length mismatch: expected %d, got %d", len(expected), len(got))
	}
	for i, r := range expected {
		if got[i] != r {
			t.Fatalf("mismatch at index %d: expected %q, got %q", i, r, got[i])
		}
	}
}

func TestBuildCharset_ExcludeSet(t *testing.T) {
	exclude := map[rune]bool{'a': true, 'Z': true, '0': true, '!': true}
	got := BuildCharset(true, true, true, true, exclude)

	for _, r := range []rune{'a', 'Z', '0', '!'} {
		for _, gr := range got {
			if gr == r {
				t.Fatalf("excluded rune still present: %q", r)
			}
		}
	}
}

func TestBuildCharset_NoneSelected(t *testing.T) {
	exclude := map[rune]bool{}
	got := BuildCharset(false, false, false, false, exclude)
	if len(got) != 0 {
		t.Fatalf("expected empty charset, got length %d", len(got))
	}
}
