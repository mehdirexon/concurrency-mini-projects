package main

import (
	"time"

	"github.com/fatih/color"
)

// philosophers list
var philosophers = []Philosopher{
	{"Plato", Fork{4}, Fork{0}},
	{"Aristotle", Fork{0}, Fork{1}},
	{"Socrates", Fork{1}, Fork{2}},
	{"Kant", Fork{2}, Fork{3}},
	{"Nietzsche", Fork{3}, Fork{4}},
}

// define some variables

var hunger = 3 // how many times a philosopher will eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// welcome message
	color.Green("Dining Philosophers Problem")
	color.Green("----------------------------")
	color.Green("The table is empty!")

	// start the meal
	Dine()

	// end message
	color.Green("the table is empty!")

	color.Cyan("Philosophers did finish in this order:")
	for i := range donePhilosophers {
		color.White(i)
	}

}
