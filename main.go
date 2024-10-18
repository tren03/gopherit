package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/tren03/gopherit/snippets"
)

var goSuffix = ".go"

func main() {
	funcMap := make(map[string]int)
	allsnips := snippets.Snip{}
	allsnipRef := reflect.ValueOf(allsnips)
	for i := 0; i < allsnipRef.NumMethod(); i++ {
		methodName := allsnipRef.Type().Method(i).Name
		methodName = methodName[:len(methodName)-4] // to remove the Main word from the key to make it easier for the user
		methodName = strings.ToLower(methodName)
		funcMap[methodName] = i
	}

	str := ""
	flag.StringVar(&str, "create", "", "create a snippet by providing its name")
	flag.Parse()

	if str != "" {
		// prg not accepting names with .go suffix
		if !strings.HasSuffix(str, goSuffix) {
			str += goSuffix // Add .go if not already present
		}

		boilerplate := fmt.Sprintf(`package snippets

import "fmt"

// This is the dynamically generated function for your snippet
func (s Snip) %sMain() {
    fmt.Println("Welcome to your snippet!")
}
        `, strings.TrimSuffix(str, ".go"))

		snippetsDir := "./snippets"
		path := filepath.Join(snippetsDir, str)

		newFilePath := filepath.FromSlash(path)
		log.Println("PATH TO NEW FILE ", newFilePath)
		if _, err := os.Stat(newFilePath); err == nil {
			log.Println("Snippet already exists, please choose a different name")
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
			defer file.Close()
			fmt.Printf("File created successfully at %s\n", newFilePath)
			str += ".go"
			cmd := exec.Command("nvim", newFilePath)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("error executing editor command ", err)
			}
		} else {
			log.Printf("Error checking file: %v\n", err)
			return
		}
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("Welcome to gopherit \nProvide the name of the snippet you want to run as the argument during cmd call")
		return
	}

	args := os.Args
	snipToRun := strings.ToLower(args[1])

	fIndex, ok := funcMap[snipToRun]
	if !ok {
		fmt.Println("snippet not present")
		return
	}

	fmt.Printf("running snippet : %s\n", snipToRun)
	allsnipRef.Method(fIndex).Call(nil)
}
