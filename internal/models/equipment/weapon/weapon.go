package weapon

import (
	"github.com/koodeex/pathgogen/internal/core"
)

type Category string

const (
	Simple  Category = "simple"
	Martial Category = "martial"
	Exotic  Category = "exotic"
)

type Type string

const (
	OneHanded Type = "One-Handed"
	TwoHanded Type = "Two-Handed"
	Ranged    Type = "Ranged"
)

type DamageType string

const (
	Banishing DamageType = "B"
	Piercing  DamageType = "P"
	Slashing  DamageType = "S"
)

type Weapon struct {
	Name       string
	Category   Category
	Type       Type
	DmgS       *core.Dice
	DmgM       *core.Dice
	Critical   *core.CriticalRange
	DamageType *DamageType
	Range      int
	Special    string
	Source     string
	Link       string
}
