package character

import "math"

type Ability int

func (a *Ability) GetBonus() int {
	return int(math.Ceil(float64(*a-10) / 2))
}

func (a *Ability) GetValue() int {
	return int(*a)
}

type characterAbilities struct {
	Str Ability
	Dex Ability
	Con Ability
	Int Ability
	Wis Ability
	Cha Ability
}

func newCharacterAbilities(str, dex, con, int, wis, cha Ability) *characterAbilities {
	return &characterAbilities{str, dex, con, int, wis, cha}
}

func (c *characterAbilities) AddStr(val int) {
	c.Str += Ability(val)
}

func (c *characterAbilities) AddDex(val int) {
	c.Dex += Ability(val)
}

func (c *characterAbilities) AddCon(val int) {
	c.Con += Ability(val)
}

func (c *characterAbilities) AddInt(val int) {
	c.Int += Ability(val)
}

func (c *characterAbilities) AddWis(val int) {
	c.Wis += Ability(val)
}

func (c *characterAbilities) AddCha(val int) {
	c.Cha += Ability(val)
}
