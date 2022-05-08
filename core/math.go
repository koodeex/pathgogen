package core

import "fmt"

type Speed struct {
	Ft30 int
	Ft20 int
}

func NewSpeed(ft30, ft20 int) *Speed {
	return &Speed{Ft30: ft30, Ft20: ft20}
}

type DiceType int

const (
	D2  DiceType = 2
	D3  DiceType = 3
	D4  DiceType = 4
	D6  DiceType = 6
	D8  DiceType = 8
	D10 DiceType = 10
	D12 DiceType = 12
	D20 DiceType = 20
)

type Dice struct {
	Multiple int
	Dice     DiceType
}

func (d *Dice) GetDice() string {
	return fmt.Sprintf("%dd%d", d.Multiple, d.Dice)
}

type CriticalRange struct {
	Multiple int
	Minimal  int
	Maximum  int
}
