package character

import (
	"github.com/koodeex/pathgogen/core"
	"github.com/koodeex/pathgogen/models/feats"
	"math"
)

type AttackBonus struct {
	AbilityBonus int
	BABBonus     int
	SizeModifier int
	MiscModifier int
	Total        int
}

func NewAttackBonus(ability, bab, size, misc int) []*AttackBonus {
	var attackBonus []*AttackBonus
	attacksCount := int(math.Ceil(float64(bab) / 5))
	for attacks := 0; attacks < attacksCount; attacks += 1 {
		babAttack := bab - (attacks * 5)
		attackBonus = append(attackBonus, &AttackBonus{
			AbilityBonus: ability,
			BABBonus:     babAttack,
			SizeModifier: size,
			MiscModifier: misc,
			Total:        ability + babAttack + size + misc,
		})
	}
	return attackBonus
}

// TODO finish attack ability bonus getter
func getAttackAbilityBonus(character *Character) int {
	attackBonus := character.Str.GetBonus()
	for _, feat := range character.Feats {
		switch feat.Name {
		case feats.WeaponFinesse:
			attackBonus = character.Dex.GetBonus()
		}
	}
	return attackBonus
}

// TODO implement attack misc bonus
func getAttackMiscBonus(character *Character) int {
	var miscBonus int
	for _, feat := range character.Feats {
		switch feat.Name {
		case feats.WeaponFocus:
			if character.EquippedWeapon != nil && core.IsStringInTheArray(character.EquippedWeapon.Name, feat.Extra) {
				miscBonus += 1
			}
		}
	}
	return miscBonus
}

// TODO implement attack size bonus
func getAttackSizeCategory(character *Character) int {
	if character.Race != nil && character.Race.SizeCategory != nil {
		return character.Race.SizeCategory.Modifier
	}
	return 0
}
