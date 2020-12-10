//+build wireinject

package main

import "github.com/google/wire"

func InitializeOrder(msg string) Order {
	wire.Build(NewOrder, NewMessage, NewCake)
	return Order{}
}
