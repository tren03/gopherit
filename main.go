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

func main() {
	func_map := make(map[string]int)
	allsnips := snippets.Snip{}
	allsnipRef := reflect.ValueOf(allsnips)
	for i := 0; i < allsnipRef.NumMethod(); i++ {
		//method := allsnipRef.Type().Method(i)
		method_name := allsnipRef.Type().Method(i).Name
		method_name = method_name[:len(method_name)-4] // to remove the Main word from the key to make it easier for the user
		method_name = strings.ToLower(method_name)
		func_map[method_name] = i

		//fmt.Println("Calling method:", method.Name)
		//allsnipRef.Method(i).Call(nil)
	}

	str := ""
	flag.StringVar(&str, "create", "", "create a snippet by providing its name")
	flag.Parse()

	if len(str) > 0 {

		// prg not accepting names with .go suffix
		if !strings.HasSuffix(str, ".go") {
			str += ".go" // Add .go if not already present
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
			file.Write([]byte(boilerplate))
			//	if err != nil {
			//		log.Println("err creating boilerplate ", err)
			//	}
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
	snip_to_run := strings.ToLower(args[1])

	f_index, ok := func_map[snip_to_run]
	if !ok {
		fmt.Println("snippet not present")
		return
	}

	fmt.Printf("running snippet : %s\n", snip_to_run)
	allsnipRef.Method(f_index).Call(nil)
}
