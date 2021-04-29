package main

import (
	"errors"
	"fmt"
	"time"
)

type PlayerParam string
type MonsterParam string

type Monster struct {
	Name string
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

type Player struct {
	Name string
}

func NewPlayer(name string) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}

	return Player{Name: name}, nil
}


type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

// 使用wire前
func main() {
	var playName string = "瓜皮"
	var monsterName MonsterParam = "里皮"

	player, _ := NewPlayer(playName)
	monster := NewMonster(monsterName)

	mission := NewMission(player, monster)

	mission.Start()
}