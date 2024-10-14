package snippets

import (
	"fmt"
	"sync"
	"time"
)

// This is the dynamically generated function for your snippet
func (s Snip) ChannelsMain() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sending message from producer 1")
			ch <- i
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sending message from producer 2")
			ch <- i
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
        fmt.Println("closing channels")
	}()

	for item := range ch {
		fmt.Println(item)
	}

}
