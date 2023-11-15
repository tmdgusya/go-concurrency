package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	defer m.Unlock()
	m.Lock()
	msg = s
}

func main() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Universe!", &mutex)
	go updateMessage("Hello, Cosmos!", &mutex)
	wg.Wait() // I'm gonna wait for those wait groups to finish

	fmt.Println(msg)
}
