package weapon

import (
	"github.com/koodeex/pathgogen/core"
)

type Category string

const (
	Simple  Category = "simple"
	Martial Category = "martial"
	Exotic  Category = "exotic"
)

type DamageType string

const (
	Banishing DamageType = "B"
	Piercing  DamageType = "P"
	Slashing  DamageType = "S"
)

type Weapon struct {
	Name       string
	Type       Category
	DmgS       *core.Dice
	DmgM       *core.Dice
	Critical   *core.CriticalRange
	DamageType *DamageType
	Range      int
	Special    string
	Source     string
	Link       string
}
