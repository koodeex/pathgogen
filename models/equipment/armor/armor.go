package armor

import (
	"pathgogen/core"
)

type Category string

const (
	Light  Category = "light"
	Medium Category = "medium"
	Heavy  Category = "heavy"
)

type Armor struct {
	Name                     string
	Description              string
	Special                  string
	OriginalWording          string
	Type                     Category
	Cost                     int
	ArmorBonus               int
	ArmorCheckPenalty        int
	MaximumDexBonus          int
	ArcaneSpellFailureChance int
	Weight                   int
	Speed                    *core.Speed
	Link                     string
	Source                   string
}
