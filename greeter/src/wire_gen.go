// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeOrder() Order {
	message := NewMessage()
	cake := NewCake(message)
	order := NewOrder(cake)
	return order
}
