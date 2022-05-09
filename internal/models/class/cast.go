package class

import (
	"errors"
	"fmt"

	"github.com/koodeex/pathgogen/internal/models/spells"
)

type CastSlots struct {
	baseSlots   []map[int]int
	actualSlots map[int][]*spells.Spell
}

// GetSlots ...
func (s *CastSlots) GetSlots(level int, castStatBonus int) map[int]int {
	result := make(map[int]int)
	for circle, slots := range s.baseSlots[level] {
		result[circle] = slots + castStatBonus
	}
	return result
}

// AddSpell ...
func (s *CastSlots) AddSpell(level, castStatBonus, slotsRequired int, spell *spells.Spell) error {
	circleSlot := spell.Circle

	actualSlots := s.actualSlots[circleSlot]
	slots := s.GetSlots(level, castStatBonus)

	for i := 0; i < slotsRequired; i++ {
		if len(actualSlots) < slots[circleSlot] {
			s.actualSlots[circleSlot] = append(s.actualSlots[circleSlot], spell)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Cannot add spell to spellbook: max(%d), current(%d)", slots[circleSlot], len(actualSlots)))
}

// InitFullCasterSlots for Wizard, Cleric, Druid etc
func InitFullCasterSlots() *CastSlots {
	return &CastSlots{
		actualSlots: map[int][]*spells.Spell{},
		baseSlots: []map[int]int{
			// lvl1
			{
				0: 3,
				1: 1,
			},
			// lvl2
			{
				0: 4,
				1: 2,
			},
			// lvl3
			{
				0: 4,
				1: 2,
				2: 1,
			},
			// lvl4
			{
				0: 4,
				1: 3,
				2: 2,
			},
			// lvl5
			{
				0: 4,
				1: 3,
				2: 2,
				3: 1,
			},
			// lvl6
			{
				0: 4,
				1: 3,
				2: 3,
				3: 2,
			},
			// lvl7
			{
				0: 4,
				1: 4,
				2: 3,
				3: 2,
				4: 1,
			},
			// lvl8
			{
				0: 4,
				1: 4,
				2: 3,
				3: 3,
				4: 2,
			},
			// lvl9
			{
				0: 4,
				1: 4,
				2: 4,
				3: 3,
				4: 2,
				5: 1,
			},
			// lvl10
			{
				0: 4,
				1: 4,
				2: 4,
				3: 3,
				4: 3,
				5: 2,
			},
			// lvl11
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 3,
				5: 2,
				6: 1,
			},
			// lvl12
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 3,
				5: 3,
				6: 2,
			},
			// lvl13
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 3,
				6: 2,
				7: 1,
			},
			// lvl14
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 3,
				6: 3,
				7: 2,
			},
			// lvl15
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 3,
				7: 2,
				8: 1,
			},
			// lvl16
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 3,
				7: 3,
				8: 2,
			},
			// lvl17
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 4,
				7: 3,
				8: 2,
				9: 1,
			},
			// lvl18
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 4,
				7: 3,
				8: 3,
				9: 2,
			},
			// lvl19
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 4,
				7: 4,
				8: 3,
				9: 3,
			},
			// lvl20
			{
				0: 4,
				1: 4,
				2: 4,
				3: 4,
				4: 4,
				5: 4,
				6: 4,
				7: 4,
				8: 4,
				9: 4,
			},
		},
	}
}
