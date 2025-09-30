package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

// Philosopher struct
type Philosopher struct {
	Name  string
	RFork Fork
	LFork Fork
}

type Fork struct {
	ID int
}

var donePhilosophers = make(chan string, len(philosophers))

func Dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks map
	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start dining
	for i := 0; i < len(philosophers); i++ {
		// fire off the goroutine
		go Dining(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
	close(donePhilosophers)
}

func Dining(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher
	color.Yellow(p.Name + " is seated at the table.")
	seated.Done()
	seated.Wait()

	for i := hunger; i > 0; i-- {
		// Get the lock on the forks
		// To prevent the logical deadlock, we make sure that the philosopher
		if p.LFork.ID > p.RFork.ID {
			forks[p.RFork.ID].Lock()
			color.Blue(p.Name + " picked up right fork.")

			forks[p.LFork.ID].Lock()
			color.Blue(p.Name + " picked up left fork.")
		} else {
			forks[p.LFork.ID].Lock()
			color.Blue(p.Name + " picked up left fork.")

			forks[p.RFork.ID].Lock()
			color.Blue(p.Name + " picked up right fork.")
		}

		color.Blue(p.Name + " is eating.")
		time.Sleep(eatTime)

		color.Blue(p.Name + " is thinking.")
		time.Sleep(thinkTime)

		forks[p.LFork.ID].Unlock()
		forks[p.RFork.ID].Unlock()

		color.Blue(p.Name + " put down the fork.")

	}

	donePhilosophers <- p.Name

	color.Yellow(p.Name + " is satisfied and leaves the table.")
}
