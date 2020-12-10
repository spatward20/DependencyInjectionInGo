//+build wireinject

package main

import "github.com/google/wire"

func InitializeOrder() Order {
	wire.Build(NewOrder, NewMessage, NewCake)
	return Order{}
}
