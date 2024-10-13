package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

// The idea of wait group is to define a counter, where every registered goroutines go through have to decrease the counter to value 0 (wg.Wait())
// To be noticed that the counter should be always greater than 0
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	updateMessage("Hello, universe!")
	go printMessage(&wg)
	wg.Wait()

	wg.Add(1)
	updateMessage("Hello, cosmos!")
	go printMessage(&wg)
	wg.Wait()

	wg.Add(1)
	updateMessage("Hello, world!")
	go printMessage(&wg)
	wg.Wait()

}
