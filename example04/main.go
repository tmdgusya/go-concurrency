package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var (
	pizzasMade   = 0
	pizzasFailed = 0
	total        = 0
)

type Producer struct {
	// Each channel has to have some kind of type
	Data chan PizzaOrder
	Quit chan chan error
}

type PizzaOrder struct {
	PizzaNumber int
	Message     string
	Success     bool
}

// Close is simply a method of closing the channel when we are done with it
func (p *Producer) Close() error {
	ch := make(chan error)
	p.Quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received Order %d\n, Number:", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d, It will take %d time\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingredients for pizza: %d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza: %d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber,
			msg,
			success,
		}

		return &p
	}
	return &PizzaOrder{
		PizzaNumber: pizzaNumber,
	}
}

func Pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	// this loop will continue to execute, trying to make pizza
	// until the quit channel receive something
	for {
		// try to make pizzas
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.PizzaNumber
			// it only used for channel
			select {
			// we try to make a pizza (we send something to the data channel)
			case pizzaMaker.Data <- *currentPizza:

			case quitChan := <-pizzaMaker.Quit:
				// close channel
				close(pizzaMaker.Data)
				close(quitChan)
				// this goroutine gonna return
				return
			}
		}
	}
}

func main() {
	// seed the random number generator to get a same number at a same time
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		Data: make(chan PizzaOrder),
		Quit: make(chan chan error),
	}

	// run the producer in the background
	go Pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.Data {
		if i.PizzaNumber <= NumberOfPizzas {
			if i.Success {
				color.Green(i.Message)
				color.Green("Order #%d is out for delivery", i.PizzaNumber)
			} else {
				color.Red(i.Message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make: %d, with %d attempts in total", pizzasMade, pizzasFailed, total)
}
