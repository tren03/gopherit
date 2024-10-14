package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func getDirContent(filepath string) ([]fs.DirEntry,error){
	list, err := os.ReadDir(filepath)
	if err != nil {
        return list,err
	}
    return list,nil
}

func main() {
	fmt.Println("welcome to gopherit")
    list,err := getDirContent("./snippets/")
    if err!=nil{
        log.Println("err getting snippet content",err)
    }
	fmt.Println(list)

}
