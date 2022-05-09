package class

import (
	"github.com/koodeex/pathgogen/internal/core"
	a "github.com/koodeex/pathgogen/internal/models/equipment/armor"
	w "github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
	"math"
)

type Class interface {
	LevelUp()
	LevelDown()
	GetName() string
	GetDescription() string
	GetRole() string
	GetSkillRanksPerLevel() int
	GetHD() *core.Dice
	GetAverageHP() int
	GetStartingWealth() string
	GetLevel() int
	FilterWeapon(weapons ...w.Weapon) []w.Weapon
	VerifyWeapon(weapon w.Weapon) bool
	FilterArmor(armors ...a.Armor) []a.Armor
	VerifyArmor(armor a.Armor) bool
	VerifyAlignment(alignment Alignment) bool
	GetBAB() int
	GetClassSkills() []string
	GetClassFeatures() []*feats.ClassFeature
	GetFortSave() int
	GetRefSave() int
	GetWillSave() int
	IsCaster() bool
}

type class struct {
	Name               string
	Description        string
	Role               string
	AlignmentAllowed   []Alignment
	SkillRanksPerLevel int
	HD                 *core.Dice
	StartingWealth     string
	BAB                float64
	Level              int
	WeaponProficiency  []w.Category
	ArmorProficiency   []a.Category
	ClassFeatures      []*feats.ClassFeature
	ClassSkills        []string
	FortBaseSave       bool
	RefBaseSave        bool
	WillBaseSave       bool
	IsFavored          bool
	Casting            *CastSlots
}

func (c *class) IsCaster() bool {
	return c.Casting != nil
}

func (c *class) LevelUp() {
	c.Level += 1
}

func (c *class) LevelDown() {
	c.Level -= 1
}

func (c *class) GetName() string {
	return c.Name
}

func (c *class) GetDescription() string {
	return c.Description
}

func (c *class) GetRole() string {
	return c.Role
}

func (c *class) GetSkillRanksPerLevel() int {
	return c.SkillRanksPerLevel
}

func (c *class) GetHD() *core.Dice {
	return c.HD
}

func (c *class) GetAverageHP() int {
	return (int(math.Ceil(float64(c.HD.Dice)/2)) + 1) * c.Level
}

func (c *class) GetStartingWealth() string {
	return c.StartingWealth
}

func (c *class) GetLevel() int {
	return c.Level
}

func (c *class) FilterWeapon(weapons ...w.Weapon) []w.Weapon {
	var result []w.Weapon
	for _, weapon := range weapons {
		if c.VerifyWeapon(weapon) {
			result = append(result, weapon)
		}
	}
	return result
}

func (c *class) VerifyWeapon(weapon w.Weapon) bool {
	for _, proficiency := range c.WeaponProficiency {
		if weapon.Category == proficiency || weapon.Name == string(proficiency) {
			return true
		}
	}
	return false
}

func (c *class) FilterArmor(armors ...a.Armor) []a.Armor {
	var result []a.Armor
	for _, armor := range armors {
		if c.VerifyArmor(armor) {
			result = append(result, armor)
		}
	}
	return result
}

func (c *class) VerifyArmor(armor a.Armor) bool {
	for _, proficiency := range c.ArmorProficiency {
		if armor.Type == proficiency {
			return true
		}
	}
	return false
}

func (c *class) VerifyAlignment(alignment Alignment) bool {
	for _, alignmentAllowed := range c.AlignmentAllowed {
		if alignmentAllowed == alignment {
			return true
		}
	}
	return false
}

func (c *class) GetBAB() int {
	return int(math.Floor(c.BAB * float64(c.Level)))
}

func (c *class) GetClassSkills() []string {
	return c.ClassSkills
}

func (c *class) GetClassFeatures() []*feats.ClassFeature {
	var result []*feats.ClassFeature
	for _, feature := range c.ClassFeatures {
		if feature.LevelRequired <= c.Level {
			result = append(result, feature)
		}
	}
	return result
}

func (c *class) GetFortSave() int {
	return getBaseSave(c.FortBaseSave, c.Level)
}

func (c *class) GetRefSave() int {
	return getBaseSave(c.RefBaseSave, c.Level)
}

func (c *class) GetWillSave() int {
	return getBaseSave(c.WillBaseSave, c.Level)
}

func getBaseSave(isGood bool, level int) int {
	if isGood {
		return int(math.Floor(2 + 0.5*float64(level)))
	}
	return int(math.Floor(2 + 0.33*float64(level)))
}
