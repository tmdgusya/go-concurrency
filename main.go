package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println(s)
}

// one go routine
// program has one goroutine
//func main() {
//	var waitGroup sync.WaitGroup
//
//	words := []string{
//		"alpha",
//		"beta",
//		"delta",
//		"gamma",
//	}
//
//	waitGroup.Add(len(words))
//	for i, x := range words {
//		go printSomething(fmt.Sprintf("%d:%s", i, x), &waitGroup)
//	}
//
//	defer waitGroup.Wait()
//
//	printSomething("This is second thing to be printed", &waitGroup)
//}
