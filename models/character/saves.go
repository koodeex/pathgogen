package character

type SaveBonus struct {
	BaseBonus      int
	AbilityBonus   int
	MagicBonus     int
	MiscBonus      int
	TemporaryBonus int
	Total          int
}

func NewSaveBonus(base, ability, magic, misc, temporary int) *SaveBonus {
	return &SaveBonus{
		BaseBonus:      base,
		AbilityBonus:   ability,
		MagicBonus:     magic,
		MiscBonus:      misc,
		TemporaryBonus: temporary,
		Total:          base + ability + magic + misc + temporary,
	}
}

// TODO implement get magic bonus
func getSaveMagicBonus(character *Character, save string) int {
	return 0
}

// TODO implement get misc bonus
func getSaveMiscBonus(character *Character, save string) int {
	return 0
}

// TODO implement get temporary bonus
func getTemporaryBonus(character *Character, save string) int {
	return 0
}
