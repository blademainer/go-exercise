//+build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
