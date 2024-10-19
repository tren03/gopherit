package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/tren03/gopherit/snippets"
	"github.com/tren03/gopherit/utils"
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

	if len(os.Args) > 3 {
		fmt.Println("too many args", len(os.Args))
		return
	}

	// flag parsing
	createSnipName := ""

	boolRun := flag.Bool("run", false, "search through all snippets and run selected one")
	flag.StringVar(&createSnipName, "create", "", "create a snippet by providing its name")
	boolOpen := flag.Bool("open", false, "search through all snipptets and open the selected one")
	flag.Parse()

	if *boolRun == true {
		// The selected string will always have go suffix since that is the output of fzf
		selected := utils.GetDirs()
		selected = strings.ToLower(selected)

		// remove .go suffix to check against funcMap
		if strings.HasSuffix(selected, goSuffix) {
			selected = strings.TrimSuffix(selected, goSuffix)
		}

		utils.RunSnipFunc(allsnipRef, funcMap, selected)
		return
	}

	if *boolOpen == true {
		utils.OpenSnipFunc()
		return
	}

	if createSnipName != "" {
		utils.CreateSnipFunc(createSnipName)
		return
	}

	if len(os.Args) == 1 {
		fmt.Println("Welcome to gopherit \nProvide the name of the snippet you want to run as the argument during cmd call")
		fmt.Println("Flags available  -> ")
		fmt.Println("--run <no args>  : ", "search through all snippets using fzf and run selected one")
		fmt.Println("--create         : ", "create a snippet by providing its name")
		fmt.Println("--open <no args> : ", "search through all snippets using fzf and open selected one")
		return
	}

	args := os.Args
	snipToRun := strings.ToLower(args[1])

	utils.RunSnipFunc(allsnipRef, funcMap, snipToRun)
}
