package armor

import (
	"pathgogen/core"
)

type ShieldType string

const (
	LightShield ShieldType = "light"
	TowerShield ShieldType = "tower"
)

type Shield struct {
	Name                     string
	Benefit                  string
	Type                     *ShieldType
	Cost                     int
	ArmorBonus               int
	MaximumDexBonus          int
	ArmorCheckPenalty        int
	ArcaneSpellFailureChance int
	Speed                    *core.Speed
	Weight                   int
	Link                     string
	Source                   string
}
