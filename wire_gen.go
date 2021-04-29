// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitEndingA(name string) EndingA {
	player := NewPlayer(name)
	monster := NewMonster(name)
	endingA := NewEndingA(player, monster)
	return endingA
}

func InitEndingB(name string) EndingB {
	player := NewPlayer(name)
	monster := NewMonster(name)
	endingB := NewEndingB(player, monster)
	return endingB
}
