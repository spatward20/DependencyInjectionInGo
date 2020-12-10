package main

import (
	"fmt"
)

//A Cake
type Cake struct {
	Message string
}

// Cake constructor
func NewCake(msg string) Cake {
	return Cake{Message: msg}
}

// Bake the Cake!
func (c Cake) Bake() string {
	return c.Message
}

// A Person to place the order
type Person struct {
	Name  string
	Email string
}

// Constructor for the person
func NewPerson(name string, email string) Person {
	return Person{Name: name, Email: email}
}

// Create an order for the cake
func NewOrder(c Cake) Order {
	return Order{Cake: c}
}

// Order for a new cake
type Order struct {
	Person Person
	Cake   Cake
}

// Process starts processing the order
func (o Order) Process() {
	msg := o.Cake.Bake()
	fmt.Println(msg)
}

func main() {
	// person walks into the store
	p := Person{Name: "Harry Potter", Email: "hpotter@verticalscope.com"}

	// select a cake
	c := Cake{Message: "Happy Birthday!"}

	// place the order
	o := Order{Person: p, Cake: c}

	// process the order
	o.Process()

}
