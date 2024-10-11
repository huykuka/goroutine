package main

import (
	"fmt"
	"sync"
)


//
func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha", "delta", "gamma",
	}

	wg.Add(len(words))
	for i, v := range words {
		go printsth(fmt.Sprintf("%d:%s", i, v), &wg)
	}
	wg.Wait()


}

func printsth(s string, wg *sync.WaitGroup) {
	defer wg.Done()


	fmt.Println(s)
}
