package class

import (
	"github.com/koodeex/pathgogen/internal/core"
	"github.com/koodeex/pathgogen/internal/models/equipment/armor"
	"github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
	"github.com/koodeex/pathgogen/internal/models/feats/ragePower"
	"github.com/koodeex/pathgogen/internal/models/skills"
)

var barbarianClassSkillList = []string{
	skills.AcrobaticsName,
	skills.ClimbName,
	skills.CraftName,
	skills.HandleAnimalName,
	skills.IntimidateName,
	skills.KnowledgeNatureName,
	skills.PerceptionName,
	skills.RideName,
	skills.SurvivalName,
	skills.SwimName,
}

type Barbarian struct {
	*class
}

func NewBarbarianClass(level int, isFavored bool) *Barbarian {
	var classSkills []string
	classSkills = append(classSkills, barbarianClassSkillList...)
	return &Barbarian{
		&class{
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
			HD: &core.Dice{
				Multiple: 1,
				Dice:     core.D12,
			},
			BAB:   core.HighBAB,
			Level: level,
			WeaponProficiency: []weapon.Category{
				weapon.Simple,
				weapon.Martial,
			},
			ArmorProficiency: []armor.Category{
				armor.Light,
				armor.Medium,
			},
			SkillRanksPerLevel: 4,
			ClassFeatures:      initBarbarianFeatures(),
			ClassSkills:        classSkills,
			FortBaseSave:       true,
			RefBaseSave:        false,
			WillBaseSave:       false,
			IsFavored:          isFavored,
		}}
}

func initBarbarianFeatures() []*feats.ClassFeature {
	return []*feats.ClassFeature{
		{
			Name:          "Fast Movement",
			AbilityType:   feats.Ex,
			Description:   `<p>A barbarian’s land speed is faster than the norm for her race by +10 feet. This benefit applies only when he is wearing no armor, light armor, or medium armor, and not carrying a heavy load. Apply this bonus before modifying the barbarian’s speed because of any load carried or armor worn. This bonus stacks with any other bonuses to the barbarian’s land speed.</p>`,
			LevelRequired: 1,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
		},
		{
			Name:          "Rage",
			AbilityType:   feats.Ex,
			Description:   `<p>A barbarian can call upon inner reserves of strength and ferocity, granting her additional combat prowess. Starting at 1st level, a barbarian can rage for a number of rounds per day equal to 4 + her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> modifier. At each level after 1st, she can rage for 2 additional rounds. Temporary increases to <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a>, such as those gained from rage and spells like <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/b/bear-s-endurance">bear’s endurance</a>, do not increase the total number of rounds that a barbarian can rage per day. A barbarian can enter rage as a free action. The total number of rounds of rage per day is renewed after resting for 8 hours, although these hours do not need to be consecutive.</p><p>While in rage, a barbarian gains a +4 morale bonus to her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-" rel="nofollow">Strength</a> and <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a>, as well as a +2 morale bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves. In addition, she takes a –2 penalty to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">Armor Class</a>. The increase to <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> grants the barbarian 2 hit points per Hit Dice, but these disappear when the rage ends and are not lost first like <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Temporary-Hit-Points" rel="nofollow">temporary hit points</a>. While in rage, a barbarian cannot use any <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Charisma-Cha-" rel="nofollow">Charisma</a>-, <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Dexterity-Dex-" rel="nofollow">Dexterity</a>-, or <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Intelligence-Int-" rel="nofollow">Intelligence</a>-based skills (except <a href="https://www.d20pfsrd.com/skills/acrobatics" rel="nofollow">Acrobatics</a>, <a href="https://www.d20pfsrd.com/skills/fly" rel="nofollow">Fly</a>, <a href="https://www.d20pfsrd.com/skills/intimidate" rel="nofollow">Intimidate</a>, and <a href="https://www.d20pfsrd.com/skills/ride" rel="nofollow">Ride</a>) or any ability that requires patience or concentration.</p><p>A barbarian can end her rage as a free action and is <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Fatigued" rel="nofollow">fatigued</a> after rage for a number of rounds equal to 2 times the number of rounds spent in the rage. A barbarian cannot enter a new rage while <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Fatigued" rel="nofollow">fatigued</a> or <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Exhausted" rel="nofollow">exhausted</a> but can otherwise enter rage multiple times during a single encounter or combat. If a barbarian falls <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Unconscious" rel="nofollow">unconscious</a>, her rage immediately ends, placing her in peril of death.</p>`,
			LevelRequired: 1,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 2,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 4,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 6,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 8,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 10,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 12,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 14,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 16,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 18,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Rage Power",
			AbilityType:   feats.Ex,
			Description:   `<p>As a barbarian gains levels, she learns to use her rage in new ways. Starting at 2nd level, a barbarian gains a <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a>. She gains another <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage power</a> for every two levels of barbarian attained after 2nd level. A barbarian gains the benefits of <a href="https://www.d20pfsrd.com/classes/core-classes/barbarian/rage-powers">rage powers</a> only while raging, and some of these powers require the barbarian to take an action first. Unless otherwise noted, a barbarian cannot select an individual power more than once.</p><p>Any barbarian who meets the powers’ prerequisites can select and use rage powers. Totem rage powers grant powers related to a theme. A barbarian cannot select from more than one group of totem rage powers; for example, a barbarian who selects a beast totem rage power cannot later choose to gain any of the dragon totem rage powers (any rage power with “dragon totem” in its title).</p>`,
			LevelRequired: 20,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Rage_Powers_Ex",
			Special:       []*ragePower.RagePower{},
		},
		{
			Name:          "Uncanny Dodge",
			AbilityType:   feats.Ex,
			Description:   `<p>At 2nd level, a barbarian gains the ability to react to danger before her senses would normally allow her to do so. She cannot be caught <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Flat-Footed" rel="nofollow">flat-footed</a>, nor does she lose her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Dexterity-Dex-">Dex</a> bonus to AC if the attacker is invisible. She still loses her <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Dexterity-Dex-" rel="nofollow">Dexterity</a> bonus to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> if immobilized. A barbarian with this ability can still lose her <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Dexterity-Dex-" rel="nofollow">Dexterity</a> bonus to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> if an opponent successfully uses the feint action against her.</p><p>If a barbarian already has <a href="https://www.d20pfsrd.com/classes/core-classes/rogue#TOC-Uncanny-Dodge-Ex-">uncanny dodge</a> from a different class, she automatically gains <a href="https://www.d20pfsrd.com/classes/core-classes/rogue#TOC-Improved-Uncanny-Dodge-Ex-">improved uncanny dodge</a> (see below) instead.</p>`,
			LevelRequired: 2,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Uncanny_Dodge_Ex",
		},
		{
			Name:          "Trap Sense +1",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 3,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       1,
		},
		{
			Name:          "Trap Sense +2",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 6,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       2,
		},
		{
			Name:          "Trap Sense +3",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 9,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       3,
		},
		{
			Name:          "Trap Sense +4",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 12,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       4,
		},
		{
			Name:          "Trap Sense +5",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 15,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       5,
		},
		{
			Name:          "Trap Sense +6",
			AbilityType:   feats.Ex,
			Description:   `<p>At 3rd level, a barbarian gains a +1 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Reflex" rel="nofollow">Reflex</a> saves made to avoid traps and a +1 <a href="https://www.d20pfsrd.com/gamemastering/combat#dodge_bonus" rel="nofollow">dodge bonus</a> to <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Armor-Class" rel="nofollow">armor class</a> against attacks made by traps. These bonuses increase by +1 every three barbarian levels thereafter (6th, 9th, 12th, 15th, and 18th level). Trap sense bonuses gained from multiple classes stack.</p>`,
			LevelRequired: 18,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Trap-Sense-Ex-",
			Special:       6,
		},
		{
			Name:          "Improved Uncanny Dodge",
			AbilityType:   feats.Ex,
			Description:   `<p>At 5th level and higher, a barbarian can no longer be flanked. This defense denies a rogue the ability to sneak attack the barbarian by flanking her, unless the attacker has at least four more rogue levels than the target has barbarian levels.</p><p>If a character already has <a href="https://www.d20pfsrd.com/classes/core-classes/rogue#TOC-Uncanny-Dodge-Ex-">uncanny dodge</a> (see above) from another class, the levels from the classes that grant <a href="https://www.d20pfsrd.com/classes/core-classes/rogue#TOC-Uncanny-Dodge-Ex-">uncanny dodge</a> stack to determine the minimum rogue level required to flank the character.</p>`,
			LevelRequired: 5,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Improved_Uncanny_Dodge_Ex",
		},
		{
			Name:          "Damage Reduction 1/-",
			AbilityType:   feats.Ex,
			Description:   `<p>At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.</p>`,
			LevelRequired: 7,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       1,
		},
		{
			Name:          "Damage Reduction 2/-",
			AbilityType:   feats.Ex,
			Description:   `<p>At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.</p>`,
			LevelRequired: 10,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       2,
		},
		{
			Name:          "Damage Reduction 3/-",
			AbilityType:   feats.Ex,
			Description:   `<p>At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.</p>`,
			LevelRequired: 13,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       3,
		},
		{
			Name:          "Damage Reduction 4/-",
			AbilityType:   feats.Ex,
			Description:   `<p>At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.</p>`,
			LevelRequired: 16,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       4,
		},
		{
			Name:          "Damage Reduction 5/-",
			AbilityType:   feats.Ex,
			Description:   `<p>At 7th level, a barbarian gains <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a>. Subtract 1 from the damage the barbarian takes each time she is dealt damage from a weapon or a natural attack. At 10th level, and every three barbarian levels thereafter (13th, 16th, and 19th level), this <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">damage reduction</a> rises by 1 point. <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction" rel="nofollow">Damage reduction</a> can reduce damage to 0 but not below 0.</p>`,
			LevelRequired: 19,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Damage-Reduction-Ex-",
			Special:       5,
		},
		{
			Name:          "Greater Rage",
			AbilityType:   feats.Ex,
			Description:   `<p>At 11th level, when a barbarian enters rage, the morale bonus to her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-" rel="nofollow">Strength</a> and <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> increases to +6 and the morale bonus on her <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves increases to +3.</p>`,
			LevelRequired: 11,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#TOC-Greater-Rage-Ex-",
		},
		{
			Name:          "Indomitable Will",
			AbilityType:   feats.Ex,
			Description:   `<p>While in rage, a barbarian of 14th level or higher gains a +4 bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves to resist enchantment spells. This bonus stacks with all other modifiers, including the morale bonus on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves she also receives during her rage.</p>`,
			LevelRequired: 14,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Indomitable_Will_Ex",
		},
		{
			Name:          "Tireless Rage",
			AbilityType:   feats.Ex,
			Description:   `<p>Starting at 17th level, a barbarian no longer becomes <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Fatigued" rel="nofollow">fatigued</a> at the end of her rage.</p>`,
			LevelRequired: 17,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Tireless_Rage_Ex",
		},
		{
			Name:          "Mighty Rage",
			AbilityType:   feats.Ex,
			Description:   `<p>At 20th level, when a barbarian enters rage, the morale bonus to her <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-" rel="nofollow">Strength</a> and <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Constitution-Con-" rel="nofollow">Constitution</a> increases to +8 and the morale bonus on her <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Will" rel="nofollow">Will</a> saves increases to +4.</p>`,
			LevelRequired: 20,
			Link:          "https://www.d20pfsrd.com/classes/Core-classes/Barbarian/#Mighty_Rage_Ex",
		},
	}
}
