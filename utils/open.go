package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func OpenSnipFunc() {
	fmt.Println("you need to search")

	arr := []string{}
	files, err := os.ReadDir("./snippets")
	for _, obj := range files {
		arr = append(arr, obj.Name())
	}

	data := strings.Join(arr, "\n")

	// Create an io.Reader from the string
	reader := strings.NewReader(data)

	// Call the fzf function to search in the array
	selected, err := fzf(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Selected item:", selected)
	openNvim(selected)
	return

}

func openNvim(selected string) {

    newFilePath := "./snippets/"+selected
	cmd := exec.Command("nvim", newFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

    err := cmd.Run()
	if err != nil {
		log.Printf("Error running command: %v", err)
		return
	}
}
