package passgen

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func writeScript(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}
	return path
}

func TestCopyToClipboardEmpty(t *testing.T) {
	if err := CopyToClipboard(""); err == nil {
		t.Fatal("expected error for empty input")
	}
}

func TestCopyToClipboardLinuxNoTools(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("linux-only")
	}
	tmp := t.TempDir()
	t.Setenv("PATH", tmp)

	if err := CopyToClipboard("data"); err == nil {
		t.Fatal("expected error when no clipboard tools are available")
	}
}

func TestCopyToClipboardLinuxWlCopy(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("linux-only")
	}
	tmp := t.TempDir()
	outPath := filepath.Join(tmp, "out.txt")
	t.Setenv("CLIP_OUT", outPath)
	writeScript(t, tmp, "wl-copy", "#!/bin/sh\n/bin/cat - > \"$CLIP_OUT\"\n")
	t.Setenv("PATH", tmp)

	if err := CopyToClipboard("hello"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("unexpected clipboard content: %q", string(data))
	}
}

func TestCopyToClipboardLinuxXclipFallback(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("linux-only")
	}
	tmp := t.TempDir()
	outPath := filepath.Join(tmp, "out.txt")
	t.Setenv("CLIP_OUT", outPath)
	writeScript(t, tmp, "xclip", "#!/bin/sh\n/bin/cat - > \"$CLIP_OUT\"\n")
	t.Setenv("PATH", tmp)

	if err := CopyToClipboard("hello-xclip"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != "hello-xclip" {
		t.Fatalf("unexpected clipboard content: %q", string(data))
	}
}

func TestCopyToClipboardDarwinPbcopy(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("darwin-only")
	}
	tmp := t.TempDir()
	outPath := filepath.Join(tmp, "out.txt")
	t.Setenv("CLIP_OUT", outPath)
	writeScript(t, tmp, "pbcopy", "#!/bin/sh\n/bin/cat - > \"$CLIP_OUT\"\n")
	t.Setenv("PATH", tmp)

	if err := CopyToClipboard("mac"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != "mac" {
		t.Fatalf("unexpected clipboard content: %q", string(data))
	}
}
