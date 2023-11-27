package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter q to quit)")

	for {
		// print a prompt
		fmt.Printf("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		// wait for a response
		response := <-pong

		fmt.Printf("response : %s\n", response)
	}
}
