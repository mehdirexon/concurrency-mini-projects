package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const PizzaNumber = 10

var pizzasFailed, pizzasMade int

type Producer struct {
	DataChannel chan PizzaOrder
	Quit        chan chan error
}

type PizzaOrder struct {
	OrderID int
	Message string
	Success bool
}

// makePizza creates a pizza. it gets order number in input and returns a full filled PizzaOrder struct or just ID
func makePizza(OrderNum int) *PizzaOrder {
	OrderNum++

	if OrderNum <= PizzaNumber {
		color.Cyan("Received order #%d", OrderNum)

		delay := rand.Intn(3) + 1 // 1 to 3 seconds delay
		rnd := rand.Intn(12) + 1

		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
			if rnd <= 2 {
				msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", OrderNum)
			} else if rnd <= 4 {
				msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", OrderNum)
			}
		} else {
			pizzasMade++
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready!", OrderNum)
		}
		color.Yellow("Making pizza #%d. it will take %d seconds ...\n", OrderNum, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		return &PizzaOrder{
			OrderID: OrderNum,
			Message: msg,
			Success: success,
		}

	}

	return &PizzaOrder{
		OrderID: OrderNum,
	}
}

// Pizzeria is the main goroutine responsible for creating pizza and hand it to the consumer.
func Pizzeria(pizzaMaker *Producer) {
	var i = 0

	for {
		// Make a pizza
		pizza := makePizza(i)

		if pizza != nil {
			i = pizza.OrderID

			select {
			case pizzaMaker.DataChannel <- *pizza:

			case quitChan := <-pizzaMaker.Quit:
				close(pizzaMaker.DataChannel)
				close(quitChan)
				return
			}
		}

	}
}

// Close is a receiver function for producer struct to close all channels.
func (p *Producer) Close() error {
	quitChan := make(chan error)
	p.Quit <- quitChan
	return <-quitChan
}
