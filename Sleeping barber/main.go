package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	seatingCap  = 10
	arrivalRate = 100
	cutDuration = 1000 * time.Millisecond
	timeOpen    = 10 * time.Second
)

func main() {

	color.Cyan("The Sleeping barber")
	color.Cyan("--------------------")

	// Making needed channels
	clientChan := make(chan string, 10)
	doneChan := make(chan bool)

	// Creating the shop
	shop := &BarberShop{
		Open:            true,
		ClientChan:      clientChan,
		HairCutDuration: cutDuration,
		BarbersDoneChan: doneChan,
		BarbersCount:    0,
		ShopCap:         seatingCap,
	}

	color.Green("The shop is open...")

	// Add barbers
	shop.AddBarber("Mehdi")
	shop.AddBarber("Ali")
	shop.AddBarber("Reza")

	// Open the barbershop
	closing := make(chan bool) // prevent new customer
	closed := make(chan bool)  // tell the barbers to finish the work and go home

	go func() {
		<-time.After(timeOpen)
		closing <- true
		shop.Close()
		closed <- true
	}()

	// Add clients
	client := 0

	go func() {
		for {
			randomArrivalTime := rand.Int() % (2 * arrivalRate)

			select {
			case <-closing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomArrivalTime)):
				shop.AddClient(fmt.Sprintf("client #%d", client))
				client++
			}
		}
	}()

	// Wait until closed signal
	<-closed
}
