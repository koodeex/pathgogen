package class

import (
	"github.com/koodeex/pathgogen/models/equipment/armor"
	"github.com/koodeex/pathgogen/models/equipment/weapon"
	"github.com/koodeex/pathgogen/models/feats"
)

var barbarianClassSkills = []string{
	"Acrobatics",
	"Climb",
	"Craft",
	"Handle Animal",
	"Intimidate",
	"Knowledge (nature)",
	"Perception",
	"Ride",
	"Survival",
	"Swim",
}

type Barbarian struct {
	class
}

func NewBarbarianClass(level int, isFavored bool) *Barbarian {
	var classSkills []string
	classSkills = append(classSkills, barbarianClassSkills...)
	return &Barbarian{class{
		Name:        "Barbarian",
		Description: "For some, there is only rage. In the ways of their people, in the fury of their passion, in the howl of battle, conflict is all these brutal souls know. Savages, hired muscle, masters of vicious martial techniques, they are not soldiers or professional warriors—they are the battle possessed, creatures of slaughter and spirits of war. Known as barbarians, these warmongers know little of training, preparation, or the rules of warfare; for them, only the moment exists, with the foes that stand before them and the knowledge that the next moment might hold their death. They possess a sixth sense in regard to danger and the endurance to weather all that might entail. These brutal warriors might rise from all walks of life, both civilized and savage, though whole societies embracing such philosophies roam the wild places of the world. Within barbarians storms the primal spirit of battle, and woe to those who face their rage.",
		Role:        "Barbarians excel in combat, possessing the martial prowess and fortitude to take on foes seemingly far superior to themselves. With rage granting them boldness and daring beyond that of most other warriors, barbarians charge furiously into battle and ruin all who would stand in their way.",
		AlignmentAllowed: []Alignment{
			ChaoticEvil,
			ChaoticGood,
			ChaoticNeutral,
			NeutralEvil,
			Neutral,
			NeutralGood,
		},
		BAB:   1,
		Level: level,
		WeaponProficiency: []weapon.Category{
			weapon.Simple,
			weapon.Martial,
		},
		ArmorProficiency: []armor.Category{
			armor.Light,
			armor.Medium,
		},
		ClassFeatures: InitBarbarianFeatures(),
		ClassSkills:   classSkills,
		FortBaseSave:  true,
		RefBaseSave:   false,
		WillBaseSave:  false,
		IsFavored:     isFavored,
	}}
}

func InitBarbarianFeatures() []*feats.ClassFeature {
	return []*feats.ClassFeature{
		{
			Name:          "Fast Movement",
			AbilityType:   feats.Ex,
			Description:   "A barbarian’s land speed is faster than the norm for her race by +10 feet. This benefit applies only when he is wearing no armor, light armor, or medium armor, and not carrying a heavy load. Apply this bonus before modifying the barbarian’s speed because of any load carried or armor worn. This bonus stacks with any other bonuses to the barbarian’s land speed.",
			LevelRequired: 1,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
		},
		{
			Name:          "Rage",
			AbilityType:   feats.Ex,
			Description:   `A barbarian can call upon inner reserves of strength and ferocity, granting her additional combat prowess. Starting at 1st level, a barbarian can rage for a number of rounds per day equal to 4 + her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> modifier. At each level after 1st, she can rage for 2 additional rounds. Temporary increases to <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a>, such as those gained from rage and spells like <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/b/bear-s-endurance">bear’s endurance</a>, do not increase the total number of rounds that a barbarian can rage per day. A barbarian can enter rage as a free action. The total number of rounds of rage per day is renewed after resting for 8 hours, although these hours do not need to be consecutive.`,
			LevelRequired: 1,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 2,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 4,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 6,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 8,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 10,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 12,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 14,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 16,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 18,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.`,
			LevelRequired: 20,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*feats.RagePower{},
		},
		{
			Name:          "Uncanny Dodge",
			AbilityType:   feats.Ex,
			Description:   `At 2nd level, a barbarian gains the ability to react to danger before her senses would normally allow her to do so. She cannot be caught <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Flat-Footed" rel="nofollow">flat-footed</a>, nor does she lose her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Dexterity-Dex-">Dex</a> bonus to AC if the attacker is invisible. She still loses her <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Dexterity-Dex-" rel="nofollow">Dexterity</a> bonus to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> if immobilized. A barbarian with this ability can still lose her <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Dexterity-Dex-" rel="nofollow">Dexterity</a> bonus to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> if an opponent successfully uses the feint action against her.`,
			LevelRequired: 2,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Uncanny_Dodge_Ex",
		},
		{
			Name:          "Trap Sense +1",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 3,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       1,
		},
		{
			Name:          "Trap Sense +2",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 6,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       2,
		},
		{
			Name:          "Trap Sense +3",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 9,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       3,
		},
		{
			Name:          "Trap Sense +4",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 12,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       4,
		},
		{
			Name:          "Trap Sense +5",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 15,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       5,
		},
		{
			Name:          "Trap Sense +6",
			AbilityType:   feats.Ex,
			Description:   `At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.`,
			LevelRequired: 18,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       6,
		},
		{
			Name:          "Improved Uncanny Dodge",
			AbilityType:   feats.Ex,
			Description:   `At 5th level and higher, a barbarian can no longer be flanked. This defense denies a rogue the ability to sneak attack the barbarian by flanking her, unless the attacker has at least four more rogue levels than the target has barbarian levels.`,
			LevelRequired: 5,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Improved_Uncanny_Dodge_Ex",
		},
		{
			Name:          "Damage Reduction 1/-",
			AbilityType:   feats.Ex,
			Description:   `At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.`,
			LevelRequired: 7,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       1,
		},
		{
			Name:          "Damage Reduction 2/-",
			AbilityType:   feats.Ex,
			Description:   `At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.`,
			LevelRequired: 10,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       2,
		},
		{
			Name:          "Damage Reduction 3/-",
			AbilityType:   feats.Ex,
			Description:   `At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.`,
			LevelRequired: 13,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       3,
		},
		{
			Name:          "Damage Reduction 4/-",
			AbilityType:   feats.Ex,
			Description:   `At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.`,
			LevelRequired: 16,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       4,
		},
		{
			Name:          "Damage Reduction 5/-",
			AbilityType:   feats.Ex,
			Description:   `At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.`,
			LevelRequired: 19,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       5,
		},
		{
			Name:          "Greater Rage",
			AbilityType:   feats.Ex,
			Description:   `At 11th level, when a barbarian enters rage, the morale bonus to her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-" rel="nofollow">Strength</a> and <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> increases to +6 and the morale bonus on her <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves increases to +3.`,
			LevelRequired: 11,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Greater-Rage-Ex-",
		},
		{
			Name:          "Indomitable Will",
			AbilityType:   feats.Ex,
			Description:   `While in rage, a barbarian of 14th level or higher gains a +4 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves to resist enchantment spells. This bonus stacks with all other modifiers, including the morale bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves she also receives during her rage.`,
			LevelRequired: 14,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Indomitable_Will_Ex",
		},
		{
			Name:          "Tireless Rage",
			AbilityType:   feats.Ex,
			Description:   `Starting at 17th level, a barbarian no longer becomes <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Fatigued" rel="nofollow">fatigued</a> at the end of her rage.`,
			LevelRequired: 17,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Tireless_Rage_Ex",
		},
		{
			Name:          "Mighty Rage",
			AbilityType:   feats.Ex,
			Description:   `At 20th level, when a barbarian enters rage, the morale bonus to her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-" rel="nofollow">Strength</a> and <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> increases to +8 and the morale bonus on her <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves increases to +4.`,
			LevelRequired: 20,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Mighty_Rage_Ex",
		},
	}
}
