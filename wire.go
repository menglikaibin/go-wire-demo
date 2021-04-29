//+build wireinject

package main

import "github.com/google/wire"

func InitMission(p string, m MonsterParam) (Mission, error) {
	wire.Build(NewMonster, NewPlayer, NewMission)

	return Mission{}, nil
}
