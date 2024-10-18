package snippets

import (
	"fmt"
	"os/exec"
	"strings"
)

// This is the dynamically generated function for your snippet
func (s Snip) FzfpipeMain() {
	items := []string{"Option 1", "Option 2", "Option 3", "Option 4"}

	// Create a command to run fzf
	cmd := exec.Command("fzf")

	// Create a pipe to pass input to fzf
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return
	}

	// Start the fzf process
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting fzf:", err)
		return
	}

	// Write the array items to fzf's stdin
	for _, item := range items {
		_, err := stdin.Write([]byte(item + "\n"))
		if err != nil {
			fmt.Println("Error writing to stdin:", err)
			return
		}
	}

	// Close stdin after writing the input
	err = stdin.Close()
	if err != nil {
		fmt.Println("Error closing stdin:", err)
		return
	}

	// Capture the selected output from fzf
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running fzf:", err)
		return
	}

	// Print the selected item
	fmt.Printf("Selected: %s", strings.TrimSpace(string(output)))
}
