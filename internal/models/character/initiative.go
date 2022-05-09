package character

import "github.com/koodeex/pathgogen/internal/models/feats"

type InitiativeBonus struct {
	AbilityBonus int
	MiscBonus    int
	Total        int
}

func NewInitiativeBonus(ability, miscBonus int) *InitiativeBonus {
	return &InitiativeBonus{
		AbilityBonus: ability,
		MiscBonus:    miscBonus,
		Total:        ability + miscBonus,
	}
}

func getInitiativeAbilityBonus(character *Character) int {
	return character.Dex.GetBonus()
}

func getInitiativeMiscBonus(character *Character) int {
	var miscBonus int
	for _, feat := range character.Feats {
		switch feat.Name {
		case feats.ImprovedInitiative:
			miscBonus += 4
		}
	}

	for _, traits := range character.Traits {
		switch traits.Name {
		case "Reactionary":
			miscBonus += 2
		}
	}
	return miscBonus
}
