package main

import "fmt"

// Message to write on teh cake
type Message string

// NewMessage creates a default Message.
func NewMessage() Message {
	return "Happy Birthday"
}

type Cake struct {
	Message Message
}

func NewCake(msg Message) Cake {
	return Cake{Message: msg}
}

func (c Cake) Bake() Message {
	return c.Message
}

type Order struct {
	Cake Cake
}

func NewOrder(c Cake) Order {
	return Order{Cake: c}
}

func (o Order) Process() {
	msg := o.Cake.Bake()
	fmt.Println(msg)
}

func main() {

	// m := NewMessage()

	// c := NewCake(m)

	// o := NewOrder(c)

	o := InitializeOrder()
	o.Process()
}
