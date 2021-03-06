// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitEndingA(name string) EndingA {
	player := NewPlayer(name)
	monster := _wireMonsterValue
	endingA := NewEndingA(player, monster)
	return endingA
}

var (
	_wireMonsterValue = kitty
)

func InitEndingB(name string) EndingB {
	player := NewPlayer(name)
	monster := _wireMainMonsterValue
	endingB := NewEndingB(player, monster)
	return endingB
}

var (
	_wireMainMonsterValue = kitty
)

// wire.go:

var kitty = Monster{Name: "哥斯拉"}

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))

var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))
