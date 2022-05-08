package character

import (
	"fmt"

	"github.com/koodeex/pathgogen/core"
	c "github.com/koodeex/pathgogen/models/class"
	f "github.com/koodeex/pathgogen/models/feats"
	r "github.com/koodeex/pathgogen/models/races"
	t "github.com/koodeex/pathgogen/models/traits"
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

const (
	AcrobaticsName             string = "Acrobatics"
	AppraiseName               string = "Appraise"
	BluffName                  string = "Bluff"
	ClimbName                  string = "Climb"
	CraftName                  string = "Craft"
	DiplomacyName              string = "Diplomacy"
	DisableDeviceName          string = "DisableDevice"
	DisguiseName               string = "Disguise"
	EscapeArtistName           string = "EscapeArtist"
	FlyName                    string = "Fly"
	HandleAnimalName           string = "HandleAnimal"
	HealName                   string = "Heal"
	IntimidateName             string = "Intimidate"
	KnowledgeArcanaName        string = "KnowledgeArcana"
	KnowledgeDungeoneeringName string = "KnowledgeDungeoneering"
	KnowledgeEngineeringName   string = "KnowledgeEngineering"
	KnowledgeGeographyName     string = "KnowledgeGeography"
	KnowledgeHistoryName       string = "KnowledgeHistory"
	KnowledgeLocalName         string = "KnowledgeLocal"
	KnowledgeNatureName        string = "KnowledgeNature"
	KnowledgeNobilityName      string = "KnowledgeNobility"
	KnowledgePlanesName        string = "KnowledgePlanes"
	KnowledgeReligionName      string = "KnowledgeReligion"
	LinguisticsName            string = "Linguistics"
	PerceptionName             string = "Perception"
	PerformName                string = "Perform"
	ProfessionName             string = "Profession"
	RideName                   string = "Ride"
	SenseMotiveName            string = "SenseMotive"
	SleightOfHandName          string = "SleightOfHand"
	SpellcraftName             string = "Spellcraft"
	StealthName                string = "Stealth"
	SurvivalName               string = "Survival"
	SwimName                   string = "Swim"
	UseMagicDeviceName         string = "UseMagicDevice"
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

func getFeatBonus(charFeats []*f.Feat, skill *Skill) int {
	var featBonus int
	skillName := skill.Name
	for _, feat := range charFeats {
		switch feat.Name {
		case f.Acrobatic:
			if skillName == AcrobaticsName || skillName == FlyName {
				featBonus += 2
			}
		case f.Alertness:
			if skillName == PerceptionName || skillName == SenseMotiveName {
				featBonus += 2
			}
		case f.AnimalAffinity:
			if skillName == HandleAnimalName || skillName == RideName {
				featBonus += 2
			}
		case f.Athletic:
			if skillName == ClimbName || skillName == SwimName {
				featBonus += 2
			}
		case f.BreadthofExperience:
			if isKnowledgeSkill(skillName) || skillName == ProfessionName {
				featBonus += 2
			}
		case f.Deceitful:
			if skillName == BluffName || skillName == DisguiseName {
				featBonus += 2
			}
		case f.ExoticHeritage:
			if core.IsStringInTheArray(skillName, feat.Extra) {
				featBonus += 2
			}
		case f.MagicalAptitude:
			if skillName == SpellcraftName || skillName == UseMagicDeviceName {
				featBonus += 2
			}
		case f.MasterAlchemist:
			if skillName == CraftName && string(Alchemy) == skill.Special {
				featBonus += 2
			}
		case f.Persuasive:
			if skillName == DiplomacyName || skillName == IntimidateName {
				featBonus += 2
			}
		case f.Prodigy:
			if (skillName == CraftName || skillName == ProfessionName || skillName == PerformName) && (core.IsStringInTheArray(skill.Special, feat.Extra)) {
				featBonus += 2
				if skill.Ranks >= 10 {
					featBonus += 2
				}
			}
		case f.SelfSufficient:
			if skillName == HealName || skillName == SurvivalName {
				featBonus += 2
			}
		case f.SharpSenses:
			if skillName == PerceptionName {
				featBonus += 4
			}
		case f.SkillFocus:
			if core.IsStringInTheArray(skillName, feat.Extra) {
				featBonus += 3
				if skill.Ranks >= 10 {
					featBonus += 3
				}
			}
		case f.Stealthy:
			if skillName == EscapeArtistName || skillName == StealthName {
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
	return skillName == KnowledgeDungeoneeringName ||
		skillName == KnowledgeEngineeringName ||
		skillName == KnowledgeGeographyName ||
		skillName == KnowledgeHistoryName ||
		skillName == KnowledgeLocalName ||
		skillName == KnowledgeNatureName ||
		skillName == KnowledgeNobilityName ||
		skillName == KnowledgePlanesName ||
		skillName == KnowledgeReligionName
}
