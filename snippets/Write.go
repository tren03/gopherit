package snippets

import (
	"fmt"
	"os"
)

// This is the dynamically generated function for your snippet
func (s Snip) WriteMain() {
	fd, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("err read", err)
	}
	_, err = fd.Seek(0, 0)

	// we write to sample2, what we read from sample
	wd, err := os.Create("sample2.txt")
	if err != nil {
		fmt.Println("err write", err)
	}
	_, err = wd.Seek(0, 0)

	buff := make([]byte, 5)
	for {
		n, err := fd.Read(buff)
		if err != nil {
			fmt.Println("err while reading", err)
			break
		}
		fmt.Println("bytes read ", string(buff))
		fmt.Println("size  read ", n)
        

		nw, err := wd.Write(buff)
		fmt.Println("bytes written ", string(buff))
		fmt.Println("size  written ", nw)

	}

	defer fd.Close()
	defer wd.Close()
	fmt.Println("done")

}
