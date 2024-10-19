package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var goSuffix = ".go"

func CreateSnipFunc(createSnipName string) {

	if !strings.HasSuffix(createSnipName, goSuffix) {
		createSnipName += goSuffix // Add .go if not already present
	}

	boilerplate := fmt.Sprintf(`package snippets

import "fmt"

// This is the dynamically generated function for your snippet
func (s Snip) %sMain() {
    fmt.Println("Welcome to your snippet!")
}
        `, strings.TrimSuffix(createSnipName, ".go"))

	snippetsDir := "./snippets"
	path := filepath.Join(snippetsDir, createSnipName)

	newFilePath := filepath.FromSlash(path)
	log.Println("PATH TO NEW FILE ", newFilePath)
	if _, err := os.Stat(newFilePath); err == nil {
		log.Println("Snippet already exists, please choose a different name")
		return
	} else if os.IsNotExist(err) {
		file, err := os.Create(newFilePath)
		if err != nil {
			fmt.Println("error creating new file ", err)
			return
		}
		_, err = file.WriteString(boilerplate)
		if err != nil {
			log.Println("err creating boilerplate ", err)
		}

		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

		fmt.Printf("File created successfully at %s\n", newFilePath)
		createSnipName += ".go"

		cmd := exec.Command("nvim", newFilePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			log.Printf("Error running command: %v", err)
			return
		}
	} else {
		log.Printf("Error checking file: %v\n", err)
		return
	}
	return
}
