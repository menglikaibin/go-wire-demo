package main

import (
	"fmt"
)

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{Name: name}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type EndingA struct {
	Player Player
	Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{p, m}
}

func (p EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\n", p.Player.Name, p.Monster.Name)
}

type EndingB struct {
	Player Player
	Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{p, m}
}

func (p EndingB) Appear() {
	fmt.Printf(
		"%s defeats %s, but became monster, world darker!\n",
		p.Player.Name,
		p.Monster.Name,
	)
}

func main() {
	// 使用wire前
	//var playerName = "奥特曼"
	//var monsterName = "哥斯拉"
	//
	player := NewPlayer("奥特曼")
	monster := NewMonster("哥斯拉")

	ending := NewEndingA(player, monster)
	ending.Appear()

	// 使用wire后
	//endingA := InitEndingA("奥特曼")
	//endingA.Appear()
	//
	//endingB := InitEndingB("哥斯拉")
	//endingB.Appear()
}