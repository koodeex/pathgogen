package character

var baseCmd = 10

type CMDBonus struct {
	Base         int
	BAB          int
	StrModifier  int
	DexModifier  int
	SizeModifier int
	MiscModifier int
	Total        int
}

func NewCmdBonus(bab, str, dex, size, misc int) *CMDBonus {
	return &CMDBonus{
		Base:         baseCmd,
		BAB:          bab,
		StrModifier:  str,
		DexModifier:  dex,
		SizeModifier: size,
		MiscModifier: misc,
		Total:        baseCmd + bab + str + dex + size + misc,
	}
}

// TODO finish cmb size modifier getter
func getCmdSizeModifier(character *Character) int {
	if character.Race != nil && character.Race.SizeCategory != nil {
		return character.Race.SizeCategory.Modifier
	}
	return 0
}

// TODO implement cmb misc bonus getter
func getCmdMiscBonus(character *Character) int {
	return 0
}
