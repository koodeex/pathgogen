package class

import (
	"github.com/koodeex/pathgogen/internal/core"
	"github.com/koodeex/pathgogen/internal/models/equipment/armor"
	"github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
	"github.com/koodeex/pathgogen/internal/models/skills"
)

var wizardClassSkillList = []string{
	skills.AppraiseName,
	skills.CraftName,
	skills.FlyName,
	skills.KnowledgeArcanaName,
	skills.KnowledgeDungeoneeringName,
	skills.KnowledgeEngineeringName,
	skills.KnowledgeGeographyName,
	skills.KnowledgeHistoryName,
	skills.KnowledgeLocalName,
	skills.KnowledgeNatureName,
	skills.KnowledgeNobilityName,
	skills.KnowledgePlanesName,
	skills.KnowledgeReligionName,
	skills.LinguisticsName,
	skills.ProfessionName,
	skills.SpellcraftName,
}

type Wizard struct {
	*class
}

func NewWizard(level int, isFavored bool) *Wizard {
	var classSkills []string
	classSkills = append(classSkills, wizardClassSkillList...)
	return &Wizard{
		&class{
			Name:        "Wizard",
			Description: "Beyond the veil of the mundane hide the secrets of absolute power. The works of beings beyond mortals, the legends of realms where gods and spirits tread, the lore of creations both wondrous and terrible—such mysteries call to those with the ambition and the intellect to rise above the common folk to grasp true might. Such is the path of the wizard. These shrewd magic-users seek, collect, and covet esoteric knowledge, drawing on cultic arts to work wonders beyond the abilities of mere mortals. While some might choose a particular field of magical study and become masters of such powers, others embrace versatility, reveling in the unbounded wonders of all magic. In either case, wizards prove a cunning and potent lot, capable of smiting their foes, empowering their allies, and shaping the world to their every desire.",
			Role:        "While universalist wizards might study to prepare themselves for any manner of danger, specialist wizards research schools of magic that make them exceptionally skilled within a specific focus. Yet no matter their specialty, all wizards are masters of the impossible and can aid their allies in overcoming any danger.",
			AlignmentAllowed: []Alignment{
				ChaoticEvil,
				ChaoticGood,
				ChaoticNeutral,
				NeutralEvil,
				Neutral,
				NeutralGood,
				LawfulEvil,
				LawfulNeutral,
				LawfulGood,
			},
			HD: &core.Dice{
				Multiple: 1,
				Dice:     core.D6,
			},
			BAB:                core.LowBAB,
			SkillRanksPerLevel: 2,
			Level:              level,
			WeaponProficiency: []weapon.Category{
				weapon.ClubProficiency,
				weapon.DaggerProficiency,
				weapon.HeavyCrossbowProficiency,
				weapon.LightCrossbowProficiency,
				weapon.QuarterstaffProficiency,
			},
			ArmorProficiency: []armor.Category{},
			ClassFeatures:    initWizardFeatures(),
			ClassSkills:      classSkills,
			FortBaseSave:     false,
			RefBaseSave:      false,
			WillBaseSave:     true,
			IsFavored:        isFavored,
		},
	}
}

// TODO: wizard features init
func initWizardFeatures() []*feats.ClassFeature {
	return []*feats.ClassFeature{
		{
			Name:          "Arcane Bond",
			AbilityType:   feats.ExSp,
			Description:   ``,
			Link:          "",
			LevelRequired: 1,
			Special:       struct{}{},
		},
	}
}
