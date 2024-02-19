package main

import (
	"fmt"
	"sync"
	"time"
)

func LongOperation(index int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Done", index)
}

func main() {
	fmt.Println("starting")

	wg := &sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(i int) {
			LongOperation(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}
