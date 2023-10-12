package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

// one go routine
// program has one goroutine
func main() {
	go printSomething("This is first thing to be printed")
	printSomething("This is second thing to be printed")
}
