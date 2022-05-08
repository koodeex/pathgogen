package character

import (
	"errors"
	"fmt"
	"github.com/koodeex/pathgogen/internal/core"
	"github.com/koodeex/pathgogen/internal/models/class"
	"github.com/koodeex/pathgogen/internal/models/equipment/armor"
	w "github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
	i "github.com/koodeex/pathgogen/internal/models/items"
	"github.com/koodeex/pathgogen/internal/models/prerequisites"
	r "github.com/koodeex/pathgogen/internal/models/races"
	t "github.com/koodeex/pathgogen/internal/models/traits"
)

type Character struct {
	*characterSkills
	*characterAbilities

	HD        int
	Name      string
	Race      *r.Race
	Alignment class.Alignment

	EquippedWeapon *w.Weapon
	EquippedArmor  *armor.Armor
	EquippedShield *armor.Shield

	Inventory     []*i.Item
	WondrousItems *WondrousSlots

	Classes []class.Class
	Feats   []*feats.Feat
	Traits  []*t.Trait
}

// NewCharacter ...
func NewCharacter(name string, alignment class.Alignment, hd int, traits []*t.Trait, race *r.Race, favoredClass class.Class, str, dex, con, int, wis, cha int) *Character {
	abilities := newCharacterAbilities(Ability(str), Ability(dex), Ability(con), Ability(int), Ability(wis), Ability(cha))
	skills := newCharacterSkills(&abilities.Str, &abilities.Dex, &abilities.Int, &abilities.Wis, &abilities.Cha)

	return &Character{
		HD:                 hd,
		Name:               name,
		Race:               race,
		Alignment:          alignment,
		characterSkills:    skills,
		characterAbilities: abilities,
		Traits:             traits,
		Feats:              []*feats.Feat{},
		Classes:            []class.Class{favoredClass},
		WondrousItems:      &WondrousSlots{},
		Inventory:          []*i.Item{},
	}
}

// GetDamage ...
func (c *Character) GetDamage() *DamageBonus {
	var dice = &core.Dice{
		Multiple: 1,
		Dice:     core.D3,
	}

	if c.EquippedWeapon != nil {
		dice = c.EquippedWeapon.DmgM
	}
	abilityBonus := getDamageAbilityBonus(c)
	miscBonus := getDamageMiscBonus(c)
	return NewDamageBonus(dice, abilityBonus, miscBonus)
}

// GetAttack ...
func (c *Character) GetAttack() []*AttackBonus {
	ability := getAttackAbilityBonus(c)
	bab := c.GetBAB()
	size := getAttackSizeCategory(c)
	misc := getAttackMiscBonus(c)
	return NewAttackBonus(ability, bab, size, misc)
}

// GetAC ...
func (c *Character) GetAC() *ArmorClassBonus {
	var armorBonus int
	var shieldBonus int

	abilityModifier := getACAbilityModifier(c.Dex.GetBonus(), c.EquippedArmor, c.EquippedShield)

	if c.EquippedShield != nil {
		shieldBonus = c.EquippedShield.ArmorBonus
	}
	armorBonus = getArmorBonus(c)
	dodgeBonus := getDodgeBonus(c)
	deflectionBonus := getACDeflectionBonus(c)
	naturalBonus := getACNaturalBonus(c)
	sizeBonus := getACSizeBonus(c)
	miscBonus := getACMiscBonus(c)
	return NewArmorClassBonus(abilityModifier, shieldBonus, armorBonus, dodgeBonus, deflectionBonus, naturalBonus, sizeBonus, miscBonus)
}

// GetCMB ...
func (c *Character) GetCMB() *CMBBonus {
	bab := c.GetBAB()
	str := c.Str.GetBonus()
	size := getCmbSizeModifier(c)
	misc := getCmbMiscBonus(c)

	return NewCmbBonus(bab, str, size, misc)
}

// GetCMD ...
func (c *Character) GetCMD() *CMDBonus {
	bab := c.GetBAB()
	str := c.Str.GetBonus()
	dex := c.Dex.GetBonus()
	size := getCmdSizeModifier(c)
	misc := getCmdMiscBonus(c)

	return NewCmdBonus(bab, str, dex, size, misc)
}

// GetFortSave ...
func (c *Character) GetFortSave() *SaveBonus {
	var base int
	var ability int
	var magic int
	var misc int
	var temporary int

	for _, class := range c.Classes {
		base += class.GetFortSave()
	}
	ability = c.Con.GetBonus()
	magic = getSaveMagicBonus(c, "Fort")
	misc = getSaveMiscBonus(c, "Fort")
	temporary = getTemporaryBonus(c, "Fort")

	return NewSaveBonus(base, ability, magic, misc, temporary)
}

// GetRefSave ...
func (c *Character) GetRefSave() *SaveBonus {
	var base int
	var ability int
	var magic int
	var misc int
	var temporary int

	for _, class := range c.Classes {
		base += class.GetRefSave()
	}
	ability = c.Dex.GetBonus()
	magic = getSaveMagicBonus(c, "Ref")
	misc = getSaveMiscBonus(c, "Ref")
	temporary = getTemporaryBonus(c, "Ref")

	return NewSaveBonus(base, ability, magic, misc, temporary)
}

// GetWillSave ...
func (c *Character) GetWillSave() *SaveBonus {
	var base int
	var ability int
	var magic int
	var misc int
	var temporary int

	for _, class := range c.Classes {
		base += class.GetWillSave()
	}
	ability = c.Wis.GetBonus()
	magic = getSaveMagicBonus(c, "Will")
	misc = getSaveMiscBonus(c, "Will")
	temporary = getTemporaryBonus(c, "Will")

	return NewSaveBonus(base, ability, magic, misc, temporary)
}

// AddFeat ...
// TODO: multi feats fix
func (c *Character) AddFeat(newFeat *feats.Feat) error {
	for _, p := range newFeat.Prerequisites {
		if !c.VerifyPrerequisite(p) {
			return errors.New(fmt.Sprintf("Do not meet feat prerequisite (%s:%+v): %s", p.Key, p.Value, newFeat.Name))
		}
	}
	if core.IsStringInTheArray(newFeat.Name, []string{
		feats.SkillFocus,
		feats.ExoticHeritage,
		feats.WeaponFocus,
		feats.WeaponSpecialization,
		feats.ArcaneTalent,
	}) {
		for _, feat := range c.Feats {
			if feat.Name == newFeat.Name {
				return errors.New(fmt.Sprintf("Feat already taken: %s", newFeat.Name))
			}
		}
	}
	c.Feats = append(c.Feats, newFeat)
	return nil
}

// RemoveFeat ...
func (c *Character) RemoveFeat(featToRemove *feats.Feat) {
	for idx, feat := range c.Feats {
		if featToRemove.Name == feat.Name {
			c.Feats = append(c.Feats[:idx], c.Feats[idx+1:]...)
			return
		}
	}
}

// GetBAB ...
func (c *Character) GetBAB() int {
	var bab int
	for _, class := range c.Classes {
		bab += class.GetBAB()
	}
	return bab
}

// GetSkills ...
func (c *Character) GetSkills() map[string]*SkillBonus {
	result := make(map[string]*SkillBonus)

	result[AcrobaticsName] = NewSkillBonus(getSkill(c, c.Acrobatics))
	result[AppraiseName] = NewSkillBonus(getSkill(c, c.Appraise))
	result[BluffName] = NewSkillBonus(getSkill(c, c.Bluff))
	result[ClimbName] = NewSkillBonus(getSkill(c, c.Climb))
	for _, craft := range c.Craft {
		result[fmt.Sprintf("%s (%s)", CraftName, craft.Special)] = NewSkillBonus(getSkill(c, craft))
	}
	result[DiplomacyName] = NewSkillBonus(getSkill(c, c.Diplomacy))
	result[DisableDeviceName] = NewSkillBonus(getSkill(c, c.DisableDevice))
	result[DisguiseName] = NewSkillBonus(getSkill(c, c.Disguise))
	result[EscapeArtistName] = NewSkillBonus(getSkill(c, c.EscapeArtist))
	result[FlyName] = NewSkillBonus(getSkill(c, c.Fly))
	result[HandleAnimalName] = NewSkillBonus(getSkill(c, c.HandleAnimal))
	result[HealName] = NewSkillBonus(getSkill(c, c.Heal))
	result[IntimidateName] = NewSkillBonus(getSkill(c, c.Intimidate))
	result[KnowledgeArcanaName] = NewSkillBonus(getSkill(c, c.KnowledgeArcana))
	result[KnowledgeDungeoneeringName] = NewSkillBonus(getSkill(c, c.KnowledgeDungeoneering))
	result[KnowledgeEngineeringName] = NewSkillBonus(getSkill(c, c.KnowledgeEngineering))
	result[KnowledgeGeographyName] = NewSkillBonus(getSkill(c, c.KnowledgeGeography))
	result[KnowledgeHistoryName] = NewSkillBonus(getSkill(c, c.KnowledgeHistory))
	result[KnowledgeLocalName] = NewSkillBonus(getSkill(c, c.KnowledgeLocal))
	result[KnowledgeNatureName] = NewSkillBonus(getSkill(c, c.KnowledgeNature))
	result[KnowledgeNobilityName] = NewSkillBonus(getSkill(c, c.KnowledgeNobility))
	result[KnowledgePlanesName] = NewSkillBonus(getSkill(c, c.KnowledgePlanes))
	result[KnowledgeReligionName] = NewSkillBonus(getSkill(c, c.KnowledgeReligion))
	result[LinguisticsName] = NewSkillBonus(getSkill(c, c.Linguistics))
	result[PerceptionName] = NewSkillBonus(getSkill(c, c.Perception))
	for _, perform := range c.Perform {
		result[fmt.Sprintf("%s (%s)", CraftName, perform.Special)] = NewSkillBonus(getSkill(c, perform))
	}
	for _, profession := range c.Profession {
		result[fmt.Sprintf("%s (%s)", CraftName, profession.Special)] = NewSkillBonus(getSkill(c, profession))
	}
	result[RideName] = NewSkillBonus(getSkill(c, c.Ride))
	result[SenseMotiveName] = NewSkillBonus(getSkill(c, c.SenseMotive))
	result[SleightOfHandName] = NewSkillBonus(getSkill(c, c.SleightOfHand))
	result[SpellcraftName] = NewSkillBonus(getSkill(c, c.Spellcraft))
	result[StealthName] = NewSkillBonus(getSkill(c, c.Stealth))
	result[SurvivalName] = NewSkillBonus(getSkill(c, c.Survival))
	result[SwimName] = NewSkillBonus(getSkill(c, c.Swim))
	result[UseMagicDeviceName] = NewSkillBonus(getSkill(c, c.UseMagicDevice))

	return result
}

func (c *Character) VerifyPrerequisite(prerequisite *prerequisites.Prerequisite) bool {
	switch prerequisite.Key {
	case prerequisites.FeatPrerequisite:
		for _, feat := range c.Feats {
			if feat.Name == prerequisite.Value.(string) {
				return true
			}
		}

	case prerequisites.StrPrerequisite:
		if c.Str.GetValue() >= prerequisite.Value.(int) {
			return true
		}
	case prerequisites.ConPrerequisite:
		if c.Dex.GetValue() >= prerequisite.Value.(int) {
			return true
		}
	case prerequisites.IntPrerequisite:
		if c.Con.GetValue() >= prerequisite.Value.(int) {
			return true
		}
	case prerequisites.DexPrerequisite:
		if c.Int.GetValue() >= prerequisite.Value.(int) {
			return true
		}
	case prerequisites.WisPrerequisite:
		if c.Wis.GetValue() >= prerequisite.Value.(int) {
			return true
		}
	case prerequisites.ChaPrerequisite:
		if c.Cha.GetValue() >= prerequisite.Value.(int) {
			return true
		}

	case prerequisites.BabPrerequisite:
		if c.GetBAB() >= prerequisite.Value.(int) {
			return true
		}

	case prerequisites.ClassFeaturesPrerequisites:
		for _, class := range c.Classes {
			for _, classFeature := range class.GetClassFeatures() {
				if classFeature.Name == prerequisite.Value.(string) {
					return true
				}
			}
		}
	case prerequisites.TraitPrerequisites:
		for _, trait := range c.Traits {
			if trait.Name == prerequisite.Value.(string) {
				return true
			}
		}
	case prerequisites.MultiplePrerequisites:
		for _, p := range prerequisite.Value.([]*prerequisites.Prerequisite) {
			if !c.VerifyPrerequisite(p) {
				return false
			}
		}
		return true

	case prerequisites.AcrobaticsPrerequisite:
		if c.Acrobatics.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.AppraisePrerequisite:
		if c.Appraise.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.BluffPrerequisite:
		if c.Bluff.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.ClimbPrerequisite:
		if c.Climb.Ranks == prerequisite.Value.(int) {
			return true
		}
	//case CraftPrerequisite:
	//	if c.Craft.Ranks == prerequisite.Value.(int) {
	//		return true
	//	}
	case prerequisites.DiplomacyPrerequisite:
		if c.Diplomacy.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.DisableDevicePrerequisite:
		if c.DisableDevice.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.DisguisePrerequisite:
		if c.Disguise.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.EscapeArtistPrerequisite:
		if c.EscapeArtist.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.FlyPrerequisite:
		if c.Fly.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.HandleAnimalPrerequisite:
		if c.HandleAnimal.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.HealPrerequisite:
		if c.Heal.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.IntimidatePrerequisite:
		if c.Intimidate.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeArcanaPrerequisite:
		if c.KnowledgeArcana.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeDungeoneeringPrerequisite:
		if c.KnowledgeDungeoneering.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeEngineeringPrerequisite:
		if c.KnowledgeEngineering.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeGeographyPrerequisite:
		if c.KnowledgeGeography.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeHistoryPrerequisite:
		if c.KnowledgeHistory.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeLocalPrerequisite:
		if c.KnowledgeLocal.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeNaturePrerequisite:
		if c.KnowledgeNature.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeNobilityPrerequisite:
		if c.KnowledgeNobility.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgePlanesPrerequisite:
		if c.KnowledgePlanes.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.KnowledgeReligionPrerequisite:
		if c.KnowledgeReligion.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.LinguisticsPrerequisite:
		if c.Linguistics.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.PerceptionPrerequisite:
		if c.Perception.Ranks == prerequisite.Value.(int) {
			return true
		}
	//case f.PerformPrerequisite:
	//	if c.Perform.Ranks == prerequisite.Value.(int) {
	//		return true
	//	}
	//case f.ProfessionPrerequisite:
	//	if c.Profession.Ranks == prerequisite.Value.(int) {
	//		return true
	//	}
	case prerequisites.RidePrerequisite:
		if c.Ride.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.SenseMotivePrerequisite:
		if c.SenseMotive.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.SleightOfHandPrerequisite:
		if c.SleightOfHand.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.SpellcraftPrerequisite:
		if c.Spellcraft.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.StealthPrerequisite:
		if c.Stealth.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.SurvivalPrerequisite:
		if c.Survival.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.SwimPrerequisite:
		if c.Swim.Ranks == prerequisite.Value.(int) {
			return true
		}
	case prerequisites.UseMagicDevicePrerequisite:
		if c.UseMagicDevice.Ranks == prerequisite.Value.(int) {
			return true
		}
	}

	return false
}
