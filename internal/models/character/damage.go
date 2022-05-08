package character

import (
	"fmt"
	"math"

	"github.com/koodeex/pathgogen/internal/core"
	"github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
)

type DamageBonus struct {
	Dice      string
	Ability   int
	MiscBonus int
	Total     string
}

func NewDamageBonus(dice *core.Dice, ability, misc int) *DamageBonus {
	return &DamageBonus{
		Dice:      dice.GetDice(),
		Ability:   ability,
		MiscBonus: misc,
		Total:     fmt.Sprintf("%s + %d + %d", dice.GetDice(), ability, misc),
	}
}

// TODO: finish damage ability bonus
func getDamageAbilityBonus(character *Character) int {
	if character.EquippedWeapon != nil && character.EquippedWeapon.Type == weapon.TwoHanded {
		return int(math.Floor(float64(character.Str.GetBonus()) * 1.5))
	}
	return character.Str.GetBonus()
}

// TODO: finish misc ability bonus
func getDamageMiscBonus(character *Character) int {
	var misc int
	for _, feat := range character.Feats {
		switch feat.Name {
		case feats.WeaponSpecialization:
			if character.EquippedWeapon != nil && core.IsStringInTheArray(character.EquippedWeapon.Name, feat.Extra) {
				misc += 2
			}
		case feats.GreaterWeaponSpecialization:
			if character.EquippedWeapon != nil && core.IsStringInTheArray(character.EquippedWeapon.Name, feat.Extra) {
				misc += 2
			}
		}
	}

	return misc
}
