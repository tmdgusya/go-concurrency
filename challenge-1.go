package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	msg = s
	wg.Done()
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	msg = "Hello, world!"

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go updateMessage("Hello, Universe", &waitGroup)
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, Cosmos", &waitGroup)
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, world!", &waitGroup)
	waitGroup.Wait()
	printMessage()

}
