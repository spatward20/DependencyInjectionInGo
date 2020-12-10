package main

import (
	"fmt"
)

// Message to write on teh cake
type Message string

// NewMessage creates a default Message.
func NewMessage() Message {
	return "Happy Birthday"
}

//A Cake
type Cake struct {
	Message Message
}

// Cake constructor
func NewCake(msg Message) Cake {
	return Cake{Message: msg}
}

// Bake the Cake!
func (c Cake) Bake() Message {
	return c.Message
}

// Create an order for the cake
func NewOrder(c Cake) Order {
	return Order{Cake: c}
}

// Order for a new cake
type Order struct {
	Cake Cake
}

// Process starts processing the order
func (o Order) Process() {
	msg := o.Cake.Bake()
	fmt.Println(msg)
}

func main() {
	// // person walks into the store
	// p := NewMessage()

	// // select a cake
	// c := NewCake(p)

	// // place the order
	// o := NewOrder(c)

	o := InitializeOrder()

	// process the order
	o.Process()

}
