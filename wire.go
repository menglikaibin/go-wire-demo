//+build wireinject

package main

import "github.com/google/wire"

var kitty = Monster{Name: "哥斯拉"}

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))

func InitEndingA(name string) EndingA {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingA)

	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingB)

	return EndingB{}
}
