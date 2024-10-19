package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

var goSuffix = ".go"

func CreateSnipFunc(createSnipName string) {

    if createSnipName == ""{
        fmt.Println("invalid name")
        return
    }

    finalSnipName := createSnipName 
    if !unicode.IsUpper(rune(createSnipName[0])){
        upperChar := unicode.ToUpper(rune(createSnipName[0]))
        finalSnipName = ""
        finalSnipName = string(upperChar) + createSnipName[1:]
    }


	if !strings.HasSuffix(finalSnipName, goSuffix) {
		finalSnipName += goSuffix // Add .go if not already present
	}

	boilerplate := fmt.Sprintf(`package snippets

import "fmt"

// This is the dynamically generated function for your snippet
func (s Snip) %sMain() {
    fmt.Println("Welcome to your snippet!")
}
        `, strings.TrimSuffix(finalSnipName, ".go"))

	snippetsDir := "./snippets"
	path := filepath.Join(snippetsDir, finalSnipName)

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
		finalSnipName += ".go"

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
