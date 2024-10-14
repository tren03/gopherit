package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/tren03/gopherit/snippets"
)

func testhello() {
	fmt.Println("hello world")
}

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
