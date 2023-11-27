package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name string
	// Those will be locks upon the forks
	// Nobody else can use this fork if someone already have fork
	rightFork int
	leftFork  int
	eatTime   int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

func main() {
	// print out a welcome message
	fmt.Printf("Dining Philosophers problem")
	fmt.Printf("---------------------------")
	fmt.Println("The table is empty.")
	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// starts the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, group *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer group.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table\n", philosopher.name)
	seated.Done()

	// eat three times
	for i := 3; i > 0; i-- {
		forks[philosopher.leftFork].Lock()
		fmt.Printf("%s takes the left forks\n", philosopher.name)
		forks[philosopher.rightFork].Lock()
		fmt.Printf("%s takes the rightFork forks\n", philosopher.name)

		time.Sleep(3 * time.Second)

		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(3 * time.Second)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
	}
}
