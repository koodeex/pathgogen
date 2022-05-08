package character

type CMBBonus struct {
	BAB          int
	StrModifier  int
	SizeModifier int
	MiscModifier int
	Total        int
}

func NewCmbBonus(bab, str, size, misc int) *CMBBonus {
	return &CMBBonus{
		BAB:          bab,
		StrModifier:  str,
		SizeModifier: size,
		MiscModifier: misc,
		Total:        bab + str + size + misc,
	}
}

// TODO finish cmb size modifier getter
func getCmbSizeModifier(character *Character) int {
	if character.Race != nil && character.Race.SizeCategory != nil {
		return character.Race.SizeCategory.Modifier
	}
	return 0
}

// TODO implement cmb misc bonus getter
func getCmbMiscBonus(character *Character) int {
	return 0
}
