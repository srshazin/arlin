package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func PromptLinux(text string) (bool, error) {
	display := os.Getenv("DISPLAY")
	if display == "" {
		display = ":0" // Default to the first display
	}
	// Run the zenity command with a question prompt
	cmd := exec.Command("zenity", "--question", fmt.Sprintf("--text=%s", text), "--ok-label=Accept", "--cancel-label=Reject")
	cmd.Env = append(os.Environ(), "DISPLAY="+display)
	// Run the command and check the return code
	err := cmd.Run()
	if err != nil {
		// If there’s an error, check if it’s an exit code 1 (user clicked "Reject" or closed the dialog)
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return false, nil // User rejected
		}
		// Other errors mean something went wrong
		return false, fmt.Errorf("failed to show prompt: %w", err)
	}

	return true, nil // User accepted
}
