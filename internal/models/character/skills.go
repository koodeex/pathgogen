package character

import (
	"fmt"
	"github.com/koodeex/pathgogen/internal/core"
	c "github.com/koodeex/pathgogen/internal/models/class"
	"github.com/koodeex/pathgogen/internal/models/feats"
	r "github.com/koodeex/pathgogen/internal/models/races"
	"github.com/koodeex/pathgogen/internal/models/skills"
	t "github.com/koodeex/pathgogen/internal/models/traits"
)

type Skill struct {
	Name              string
	Ranks             int
	Ability           *Ability
	Untrained         bool
	ArmorCheckPenalty bool
	Special           string
	Link              string
}

func NewSkill(name string, ability *Ability, untrained bool, armorCheckPenalty bool, link string) *Skill {
	return &Skill{
		Name:              name,
		Ability:           ability,
		Untrained:         untrained,
		ArmorCheckPenalty: armorCheckPenalty,
		Link:              link,
	}
}

type SkillBonus struct {
	AbilityBonus int
	RankBonus    int
	MiscBonus    int
	Total        int
}

func NewSkillBonus(abilityBonus, rankBonus, miscBonus int) *SkillBonus {
	return &SkillBonus{
		AbilityBonus: abilityBonus,
		RankBonus:    rankBonus,
		MiscBonus:    miscBonus,
		Total:        abilityBonus + rankBonus + miscBonus,
	}
}

type CraftType string

const (
	Alchemy      CraftType = "Alchemy"
	Armor        CraftType = "Armor"
	Bows         CraftType = "Bows"
	Firearms     CraftType = "Firearms"
	SiegeEngines CraftType = "SiegeEngines"
	Weapons      CraftType = "Weapons"
	Traps        CraftType = "Traps"
	Varies       CraftType = "Varies"
)

const (
	AcrobaticsLink     string = "https://www.d20pfsrd.com/skills/acrobatics"
	AppraiseLink       string = "https://www.d20pfsrd.com/skills/appraise"
	BluffLink          string = "https://www.d20pfsrd.com/skills/bluff"
	ClimbLink          string = "https://www.d20pfsrd.com/skills/climb"
	CraftLink          string = "https://www.d20pfsrd.com/skills/craft"
	DiplomacyLink      string = "https://www.d20pfsrd.com/skills/diplomacy"
	DisableDeviceLink  string = "https://www.d20pfsrd.com/skills/disable-device"
	DisguiseLink       string = "https://www.d20pfsrd.com/skills/disguise"
	EscapeArtistLink   string = "https://www.d20pfsrd.com/skills/escape-artist"
	FlyLink            string = "https://www.d20pfsrd.com/skills/fly"
	HandleAnimalLink   string = "https://www.d20pfsrd.com/skills/handle-animal"
	HealLink           string = "https://www.d20pfsrd.com/skills/heal"
	IntimidateLink     string = "https://www.d20pfsrd.com/skills/intimidate"
	KnowledgeLink      string = "https://www.d20pfsrd.com/skills/knowledge"
	LinguisticsLink    string = "https://www.d20pfsrd.com/skills/linguistics"
	PerceptionLink     string = "https://www.d20pfsrd.com/skills/perception"
	PerformLink        string = "https://www.d20pfsrd.com/skills/perform"
	ProfessionLink     string = "https://www.d20pfsrd.com/skills/profession"
	RideLink           string = "https://www.d20pfsrd.com/skills/ride"
	SenseMotiveLink    string = "https://www.d20pfsrd.com/skills/sense-motive"
	SleightOfHandLink  string = "https://www.d20pfsrd.com/skills/sleight-of-hand"
	SpellcraftLink     string = "https://www.d20pfsrd.com/skills/spellcraft"
	StealthLink        string = "https://www.d20pfsrd.com/skills/stealth"
	SurvivalLink       string = "https://www.d20pfsrd.com/skills/survival"
	SwimLink           string = "https://www.d20pfsrd.com/skills/swim"
	UseMagicDeviceLink string = "https://www.d20pfsrd.com/skills/use-magic-device"
)

type characterSkills struct {
	Acrobatics             *Skill
	Appraise               *Skill
	Bluff                  *Skill
	Climb                  *Skill
	Craft                  []*Skill
	Diplomacy              *Skill
	DisableDevice          *Skill
	Disguise               *Skill
	EscapeArtist           *Skill
	Fly                    *Skill
	HandleAnimal           *Skill
	Heal                   *Skill
	Intimidate             *Skill
	KnowledgeArcana        *Skill
	KnowledgeDungeoneering *Skill
	KnowledgeEngineering   *Skill
	KnowledgeGeography     *Skill
	KnowledgeHistory       *Skill
	KnowledgeLocal         *Skill
	KnowledgeNature        *Skill
	KnowledgeNobility      *Skill
	KnowledgePlanes        *Skill
	KnowledgeReligion      *Skill
	Linguistics            *Skill
	Perception             *Skill
	Perform                []*Skill
	Profession             []*Skill
	Ride                   *Skill
	SenseMotive            *Skill
	SleightOfHand          *Skill
	Spellcraft             *Skill
	Stealth                *Skill
	Survival               *Skill
	Swim                   *Skill
	UseMagicDevice         *Skill
}

func newCharacterSkills(str, dex, int, wis, cha *Ability) *characterSkills {
	return &characterSkills{
		Acrobatics:             NewSkill("Acrobatics", dex, true, true, AcrobaticsLink),
		Appraise:               NewSkill("Appraise", int, true, false, AppraiseLink),
		Bluff:                  NewSkill("Bluff", cha, true, false, BluffLink),
		Climb:                  NewSkill("Climb", str, true, true, ClimbLink),
		Diplomacy:              NewSkill("Diplomacy", cha, true, false, DiplomacyLink),
		DisableDevice:          NewSkill("DisableDevice", dex, false, true, DisableDeviceLink),
		Disguise:               NewSkill("Disguise", cha, true, false, DisguiseLink),
		EscapeArtist:           NewSkill("EscapeArtist", dex, true, true, EscapeArtistLink),
		Fly:                    NewSkill("Fly", dex, true, true, FlyLink),
		HandleAnimal:           NewSkill("HandleAnimal", cha, false, false, HandleAnimalLink),
		Heal:                   NewSkill("Heal", wis, true, false, HealLink),
		Intimidate:             NewSkill("Intimidate", cha, true, false, IntimidateLink),
		KnowledgeArcana:        NewSkill("KnowledgeArcana", int, false, false, KnowledgeLink),
		KnowledgeDungeoneering: NewSkill("KnowledgeDungeoneering", int, false, false, KnowledgeLink),
		KnowledgeEngineering:   NewSkill("KnowledgeEngineering", int, false, false, KnowledgeLink),
		KnowledgeGeography:     NewSkill("KnowledgeGeography", int, false, false, KnowledgeLink),
		KnowledgeHistory:       NewSkill("KnowledgeHistory", int, false, false, KnowledgeLink),
		KnowledgeLocal:         NewSkill("KnowledgeLocal", int, false, false, KnowledgeLink),
		KnowledgeNature:        NewSkill("KnowledgeNature", int, false, false, KnowledgeLink),
		KnowledgeNobility:      NewSkill("KnowledgeNobility", int, false, false, KnowledgeLink),
		KnowledgePlanes:        NewSkill("KnowledgePlanes", int, false, false, KnowledgeLink),
		KnowledgeReligion:      NewSkill("KnowledgeReligion", int, false, false, KnowledgeLink),
		Linguistics:            NewSkill("Linguistics", int, true, false, LinguisticsLink),
		Perception:             NewSkill("Perception", wis, true, false, PerceptionLink),
		Ride:                   NewSkill("Ride", dex, true, true, RideLink),
		SenseMotive:            NewSkill("SenseMotive", wis, true, false, SenseMotiveLink),
		SleightOfHand:          NewSkill("SleightOfHand", dex, true, true, SleightOfHandLink),
		Spellcraft:             NewSkill("Spellcraft", int, false, false, SpellcraftLink),
		Stealth:                NewSkill("Stealth", dex, false, true, StealthLink),
		Survival:               NewSkill("Survival", wis, true, false, SurvivalLink),
		Swim:                   NewSkill("Swim", str, true, true, SwimLink),
		UseMagicDevice:         NewSkill("UseMagicDevice", cha, false, false, UseMagicDeviceLink),
		Craft:                  []*Skill{},
		Perform:                []*Skill{},
		Profession:             []*Skill{},
	}
}

func (c *characterSkills) NewCraftSkill(name CraftType, int *Ability) {
	c.Craft = append(c.Craft, NewSkill(fmt.Sprintf("Craft (%s)", name), int, false, false, CraftLink))
}

func (c *characterSkills) NewProfessionSkill(name string, int *Ability) {
	c.Craft = append(c.Profession, NewSkill(fmt.Sprintf("Profession (%s)", name), int, false, false, ProfessionLink))
}

func (c *characterSkills) NewPerformSkill(name string, int *Ability) {
	c.Craft = append(c.Perform, NewSkill(fmt.Sprintf("Perform (%s)", name), int, false, false, PerformLink))
}

func getSkill(character *Character, skill *Skill) (int, int, int) {
	abilityBonus := skill.Ability.GetBonus()
	ranksBonus := skill.Ranks
	miscBonus := getSkillMiscBonus(character, skill)

	return abilityBonus, ranksBonus, miscBonus
}

func getSkillMiscBonus(character *Character, skill *Skill) int {
	var miscBonus int

	miscBonus += getClassSkillBonus(character.Classes, skill)
	miscBonus += getFeatBonus(character.Feats, skill)
	miscBonus += getTraitBons(character.Traits, skill)
	miscBonus += getRaceBonus(character.Race, skill)
	for _, class := range character.Classes {
		miscBonus += getClassFeatureBonus(class, skill)
	}

	if character.EquippedArmor != nil && skill.ArmorCheckPenalty {
		miscBonus -= character.EquippedArmor.ArmorCheckPenalty
	}

	if character.EquippedShield != nil && skill.ArmorCheckPenalty {
		miscBonus -= character.EquippedShield.ArmorCheckPenalty
	}

	return miscBonus
}

func getClassSkillBonus(classes []c.Class, skill *Skill) int {
	skillName := skill.Name

	var isClassSkill bool
	for _, class := range classes {
		for _, classSkill := range class.GetClassSkills() {
			if classSkill == skillName {
				isClassSkill = true
				break
			}
		}
		if isClassSkill {
			break
		}
	}
	if isClassSkill && skill.Ranks > 0 {
		return 3
	}
	return 0
}

func getFeatBonus(charFeats []*feats.Feat, skill *Skill) int {
	var featBonus int
	skillName := skill.Name
	for _, feat := range charFeats {
		switch feat.Name {
		case feats.Acrobatic:
			if skillName == skills.AcrobaticsName || skillName == skills.FlyName {
				featBonus += 2
			}
		case feats.Alertness:
			if skillName == skills.PerceptionName || skillName == skills.SenseMotiveName {
				featBonus += 2
			}
		case feats.AnimalAffinity:
			if skillName == skills.HandleAnimalName || skillName == skills.RideName {
				featBonus += 2
			}
		case feats.Athletic:
			if skillName == skills.ClimbName || skillName == skills.SwimName {
				featBonus += 2
			}
		case feats.BreadthofExperience:
			if isKnowledgeSkill(skillName) || skillName == skills.ProfessionName {
				featBonus += 2
			}
		case feats.Deceitful:
			if skillName == skills.BluffName || skillName == skills.DisguiseName {
				featBonus += 2
			}
		case feats.ExoticHeritage:
			if core.IsStringInTheArray(skillName, feat.Extra) {
				featBonus += 2
			}
		case feats.MagicalAptitude:
			if skillName == skills.SpellcraftName || skillName == skills.UseMagicDeviceName {
				featBonus += 2
			}
		case feats.MasterAlchemist:
			if skillName == skills.CraftName && string(Alchemy) == skill.Special {
				featBonus += 2
			}
		case feats.Persuasive:
			if skillName == skills.DiplomacyName || skillName == skills.IntimidateName {
				featBonus += 2
			}
		case feats.Prodigy:
			if (skillName == skills.CraftName || skillName == skills.ProfessionName || skillName == skills.PerformName) && (core.IsStringInTheArray(skill.Special, feat.Extra)) {
				featBonus += 2
				if skill.Ranks >= 10 {
					featBonus += 2
				}
			}
		case feats.SelfSufficient:
			if skillName == skills.HealName || skillName == skills.SurvivalName {
				featBonus += 2
			}
		case feats.SharpSenses:
			if skillName == skills.PerceptionName {
				featBonus += 4
			}
		case feats.SkillFocus:
			if core.IsStringInTheArray(skillName, feat.Extra) {
				featBonus += 3
				if skill.Ranks >= 10 {
					featBonus += 3
				}
			}
		case feats.Stealthy:
			if skillName == skills.EscapeArtistName || skillName == skills.StealthName {
				featBonus += 2
			}
		}
	}
	return featBonus
}

// TODO implement get trait bonus
func getTraitBons(traits []*t.Trait, skill *Skill) int {
	return 0
}

// TODO implement get race bonus
func getRaceBonus(race *r.Race, skill *Skill) int {
	return 0
}

// TODO implement get class feature bonus
func getClassFeatureBonus(class c.Class, skill *Skill) int {
	return 0
}

func isKnowledgeSkill(skillName string) bool {
	return skillName == skills.KnowledgeDungeoneeringName ||
		skillName == skills.KnowledgeEngineeringName ||
		skillName == skills.KnowledgeGeographyName ||
		skillName == skills.KnowledgeHistoryName ||
		skillName == skills.KnowledgeLocalName ||
		skillName == skills.KnowledgeNatureName ||
		skillName == skills.KnowledgeNobilityName ||
		skillName == skills.KnowledgePlanesName ||
		skillName == skills.KnowledgeReligionName
}
