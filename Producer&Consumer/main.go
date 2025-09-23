package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// Define a producer
	producer := &Producer{
		DataChannel: make(chan PizzaOrder),
		Quit:        make(chan chan error),
	}

	// Run the pizzeria
	go Pizzeria(producer)

	// Consumer loop
	for order := range producer.DataChannel {
		if order.OrderID <= PizzaNumber {
			if order.Success {
				color.Green(order.Message)
				color.Green("order #%d is delivered to customer", order.OrderID)
			} else {
				color.Red(order.Message)
				color.Red("Customer is angry")
			}
		} else {
			err := producer.Close()
			if err != nil {
				color.Red("Error in closing producer channel")
			}
		}

	}

	// Statistics
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Metric", "Count"})
	table.Append([]string{"Successful", fmt.Sprintf("%d", pizzasMade)})
	table.Append([]string{"Failed", fmt.Sprintf("%d", pizzasFailed)})
	table.Append([]string{"Success Rate", fmt.Sprintf("%.2f%%", float64(pizzasMade)/float64(pizzasMade+pizzasFailed)*100)})

	table.Render()
}
