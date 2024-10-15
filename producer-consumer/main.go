package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func pizzaria(pizzaMake *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notification

	// try to mkae pizzas

}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//printout out a message
	color.Cyan("Pizzaria")
	color.Cyan("------------------------------")

	//create a producer and consumer

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//start the producer and consumer goroutines
	go pizzaria(pizzaJob)

	//wait for the producer and consumer to finish

	//print out the final result
	//fmt.Println("Final result:", result)
}
