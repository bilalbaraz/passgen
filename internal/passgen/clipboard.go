package passgen

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func CopyToClipboard(text string) error {
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
