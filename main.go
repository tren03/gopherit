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
		fmt.Println("too many args",len(os.Args))
		return
	}

	// flag parsing
	openSnipName := ""
	createSnipName := ""
	runSnipName := ""

	flag.StringVar(&runSnipName, "run", "", "search through all snippets and run selected one")
	flag.StringVar(&createSnipName, "create", "", "create a snippet by providing its name")
	flag.StringVar(&openSnipName, "open", "", "search through all snipptets and open the selected one")
	flag.Parse()

	if runSnipName != "" {
		utils.RunSnipFunc(allsnipRef,funcMap,"Test2")
        return
	}

	if openSnipName != "" {
		utils.OpenSnipFunc()
		return
	}

	if createSnipName != "" {
		utils.CreateSnipFunc(createSnipName)
		return
	}

	if len(os.Args) == 1 {
		fmt.Println("Welcome to gopherit \nProvide the name of the snippet you want to run as the argument during cmd call")
		return
	}

	args := os.Args
	snipToRun := strings.ToLower(args[1])

    utils.RunSnipFunc(allsnipRef,funcMap,snipToRun)
}
