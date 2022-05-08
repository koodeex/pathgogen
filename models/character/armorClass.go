package character

import (
	"strings"

	"pathgogen/models/equipment/armor"
	"pathgogen/models/feats"
)

const baseAC = 10

type ArmorClassBonus struct {
	Base            int
	AbilityBonus    int
	ShieldBonus     int
	ArmorBonus      int
	DodgeBonus      int
	DeflectionBonus int
	NaturalBonus    int
	SizeBonus       int
	MiscModifier    int
	Total           int
	FlatFooted      int
	Touch           int
}

func NewArmorClassBonus(ability, shield, armor, dodge, deflection, natural, size, misc int) *ArmorClassBonus {
	return &ArmorClassBonus{
		Base:            baseAC,
		AbilityBonus:    ability,
		ShieldBonus:     shield,
		ArmorBonus:      armor,
		DodgeBonus:      dodge,
		DeflectionBonus: deflection,
		NaturalBonus:    natural,
		SizeBonus:       size,
		MiscModifier:    misc,
		Total:           baseAC + ability + shield + armor + dodge + deflection + natural + size + misc,
		FlatFooted:      baseAC + shield + armor + deflection + natural + size + misc,
		Touch:           baseAC + ability + dodge + deflection + size + misc,
	}
}

func (a *ArmorClassBonus) GetTotal() int {
	a.Total = a.Base + a.AbilityBonus + a.ShieldBonus + a.ArmorBonus + a.DodgeBonus + a.DeflectionBonus + a.NaturalBonus + a.SizeBonus + a.MiscModifier
	return a.Total
}

func (a *ArmorClassBonus) GetFlatFooted() int {
	a.FlatFooted = a.Base + a.ShieldBonus + a.ArmorBonus + a.DeflectionBonus + a.NaturalBonus + a.SizeBonus + a.MiscModifier
	return a.FlatFooted
}

func (a *ArmorClassBonus) GetTouch() int {
	a.Touch = a.Base + a.AbilityBonus + a.DodgeBonus + a.DeflectionBonus + a.SizeBonus + a.MiscModifier
	return a.Touch
}

func getACAbilityModifier(ability int, armor *armor.Armor, shield *armor.Shield) int {
	var maxDexBonus int
	if armor != nil {
		maxDexBonus = armor.MaximumDexBonus
	}
	if shield != nil && shield.MaximumDexBonus < maxDexBonus {
		maxDexBonus = shield.MaximumDexBonus
	}
	if ability > maxDexBonus && maxDexBonus != 0 {
		ability = maxDexBonus
	}
	return ability
}

// TODO finish dodge bonus getter
func getDodgeBonus(character *Character) int {
	var dodgeBonus int
	for _, feat := range character.Feats {
		switch feat.Name {
		case feats.Dodge:
			dodgeBonus += 1
		}
	}
	return dodgeBonus
}

func getArmorBonus(character *Character) int {
	var armorBonus int
	if character.EquippedArmor != nil {
		armorBonus = character.EquippedArmor.ArmorBonus
	}
	wrists := character.WondrousItems.Wrists
	if wrists != nil && strings.Contains(wrists.Name, "Bracers of armor") && wrists.Special.(int) > armorBonus {
		armorBonus = wrists.Special.(int)
	}

	return armorBonus
}

// TODO finish deflection bonus getter
func getACDeflectionBonus(character *Character) int {
	var deflectionBonus1 int
	var deflectionBonus2 int

	ring1 := character.WondrousItems.RingRight
	if ring1 != nil && strings.Contains(ring1.Name, "Ring of protection") && ring1.Special.(int) > deflectionBonus1 {
		deflectionBonus1 = ring1.Special.(int)
	}

	ring2 := character.WondrousItems.RingRight
	if ring2 != nil && strings.Contains(ring2.Name, "Ring of protection") && ring2.Special.(int) > deflectionBonus2 {
		deflectionBonus2 = ring2.Special.(int)
	}

	if deflectionBonus1 < deflectionBonus2 {
		return deflectionBonus2
	}

	return deflectionBonus1
}

// TODO finish natural bonus getter
func getACNaturalBonus(character *Character) int {
	var naturalBonus int

	neck := character.WondrousItems.Neck
	if neck != nil && strings.Contains(neck.Name, "Amulet of natural armor") && neck.Special.(int) > naturalBonus {
		naturalBonus = neck.Special.(int)
	}

	return naturalBonus
}

// TODO implement size bonus getter
func getACSizeBonus(character *Character) int {
	return 0
}

// TODO implement misc bonus getter
func getACMiscBonus(character *Character) int {
	return 0
}
