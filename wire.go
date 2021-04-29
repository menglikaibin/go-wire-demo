//+build wireinject

package main

import "github.com/google/wire"

func InitMission(p PlayerParam, m MonsterParam) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)

	return Mission{}
}
