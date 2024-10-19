package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func RunSnipFunc(allsnipRef reflect.Value, funcMap map[string]int, snipToRun string) {
	snipToRun = strings.ToLower(snipToRun)
	fIndex, ok := funcMap[snipToRun]
	if !ok {
		fmt.Println("snippet not present")
		return
	}

	fmt.Printf("running snippet : %s\n", snipToRun)
	allsnipRef.Method(fIndex).Call(nil)
}
