package main

import (
	"fmt"
	"sync"
)

var msgMutex string
var wg sync.WaitGroup

func updateMessageMutex(message string, m *sync.Mutex) {
	defer wg.Done()
	// Lock the mutex to ensure exclusive access to the shared resource
	m.Lock()
	msgMutex = message

	// Unlock the mutex when the function is done
	defer m.Unlock()
}

func main() {
	msgMutex = "hello world"
	var mutext sync.Mutex
	wg.Add(2)
	go updateMessageMutex("Ben", &mutext)
	go updateMessageMutex("uyen", &mutext)
	wg.Wait()

	fmt.Println(msgMutex)
}
