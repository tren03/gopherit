package snippets

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
)

func Check() {

	text := "This is sample text i need to read"

	bText := []byte(text)
	byteReader := bytes.NewReader(bText)
	stringReader := strings.NewReader(text)

	fmt.Println("char left ", stringReader.Len())
	char, n, err := stringReader.ReadRune()
	if err != nil {
		fmt.Println("err ", err)
	}

	fmt.Println("char size ", string(char), n)
	fmt.Println("char left ", stringReader.Len())

	fmt.Println("bytes left ", byteReader.Len())

	b := make([]byte, 1)
	for {
		nb, err := byteReader.Read(b)
		if err != nil {
			if err == io.EOF {
				fmt.Println("finished reading from file ", err)
				break
			}
			fmt.Println("err ", err)

		}
		fmt.Println("bytes array read ", b)
		fmt.Println("bytes array size read ", nb)
		fmt.Println("bytes left ", byteReader.Len())
		fmt.Println("processing next chunk in 2 seconds...")
	}

	fmt.Println("done :)")

}

// This is the dynamically generated function for your snippet
func (s Snip) ReadwriteMain() {
    memoryUsage(Check)
}
func memoryUsage(fn func()) {
	// Record memory stats before function execution
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// Run the function
	fn()

	// Record memory stats after function execution
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	// Calculate memory allocations
	allocBefore := memStatsBefore.TotalAlloc
	allocAfter := memStatsAfter.TotalAlloc
	allocatedMemory := allocAfter - allocBefore

	// Output memory usage
	fmt.Printf("Memory allocated by function: %d bytes\n", allocatedMemory)
}
