package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	Open            bool
	ClientChan      chan string // we can use structs here
	BarbersDoneChan chan bool
	BarbersCount    int8
	ShopCap         int8
	HairCutDuration time.Duration
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.BarbersCount++
	go func() {
		isSleeping := false
		color.Yellow("Barber %s goes to waiting room to check for clients", barber)

		for {
			// check the waiting room
			if len(shop.ClientChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			// block until someone comes
			client, ok := <-shop.ClientChan
			if ok {
				if isSleeping {
					color.Green("%s wakes %s", client, barber)
					isSleeping = false
				}
				shop.CutHair(barber, client)
			} else {
				shop.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) CutHair(barber, client string) {
	color.Cyan("barber %s is cutting %s's hairs", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Cyan("barber %s is done cutting %s's hairs", barber, client)

}

func (shop *BarberShop) SendBarberHome(barber string) {
	color.Cyan("barber %s is going to home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) Close() {
	color.Cyan("shop is closing...")
	close(shop.ClientChan)

	for i := 0; i < int(shop.BarbersCount); i++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)

	color.Green("barbers all went home...")

}

func (shop *BarberShop) AddClient(client string) {
	color.Green("%s comes to the barber shop", client)

	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Blue("%s takes a sit", client)
		default:
			color.Red("there is not seats available, so %s leaves", client)
		}
	} else {
		color.Red("the shop is already closed. so %s leaves", client)
	}
}
