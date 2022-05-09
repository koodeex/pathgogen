package class

import (
	"github.com/koodeex/pathgogen/internal/core"
	"github.com/koodeex/pathgogen/internal/models/equipment/armor"
	"github.com/koodeex/pathgogen/internal/models/equipment/weapon"
	"github.com/koodeex/pathgogen/internal/models/feats"
	"github.com/koodeex/pathgogen/internal/models/feats/arcaneBond"
	"github.com/koodeex/pathgogen/internal/models/feats/arcaneDiscovery"
	"github.com/koodeex/pathgogen/internal/models/prerequisites"
	"github.com/koodeex/pathgogen/internal/models/skills"
	"github.com/koodeex/pathgogen/internal/models/spells"
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

var wizardBonusFeatAllowedCategories = []feats.FeatType{
	feats.ItemCreationFeats,
	feats.MetamagicFeats,
	arcaneDiscovery.ArcaneDiscovery,
}

var wizardBonusFeatAllowedFeats = []string{
	feats.SpellMastery,
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
			Casting:          InitFullCasterSlots(),
		},
	}
}

func (w *Wizard) GetCastStat() string {
	return "INT"
}

func (w *Wizard) AddSpell(castStatBonus int, spell *spells.Spell) error {
	slotsRequired := 1
	schools := w.GetArcaneSchools()
	for _, oppositeSchool := range schools.OppositeSchools {
		if oppositeSchool == spell.School {
			slotsRequired += 1
			break
		}
	}
	return w.Casting.AddSpell(w.Level, slotsRequired, castStatBonus, spell)
}

func (w *Wizard) SetArcaneSchools(favored []spells.SpellSchool, opposite []spells.SpellSchool) {
	for _, feature := range w.ClassFeatures {
		if feature.Name == "Arcane School" {
			schools := feature.Special.(*WizardArcaneSchools)
			schools.FavoredSchools = favored
			schools.OppositeSchools = opposite
		}
	}
}

func (w *Wizard) GetArcaneSchools() *WizardArcaneSchools {
	for _, feature := range w.ClassFeatures {
		if feature.Name == "Arcane School" {
			return feature.Special.(*WizardArcaneSchools)
		}
	}
	return nil
}

type WizardArcaneSchools struct {
	FavoredSchools  []spells.SpellSchool
	OppositeSchools []spells.SpellSchool
}

// TODO: wizard features init
func initWizardFeatures() []*feats.ClassFeature {
	return []*feats.ClassFeature{
		{
			Name:          "Arcane BondedObject",
			AbilityType:   feats.ExSp,
			Description:   `<p>At 1st level, wizards form a powerful bond with an object or a creature. This bond can take one of two forms: a <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/familiar"><i>familiar</i></a> or a <i>bonded object</i>. A <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/familiar"><i>familiar</i></a> is a magical pet that enhances the wizard’s skills and senses and can aid him in magic, while a <i>bonded object</i> is an item a wizard can use to cast additional spells or to serve as a magical item. Once a wizard makes this choice, it is permanent and cannot be changed.</p><p>Rules for bonded items are given below, while rules for <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/familiar">familiars</a> are located here: <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/familiar"><i>Familiars</i></a>.</p><p>Wizards who select a bonded object begin play with one at no cost. Objects that are the subject of an arcane bond must fall into one of the following categories: amulet, ring, staff, wand, or weapon. These objects are always masterwork quality. Weapons acquired at 1st level are not made of any special material. If the object is an amulet or ring, it must be worn to have effect, while staves, wands, and weapons must be held in one hand. If a wizard attempts to cast a spell without his bonded object worn or in hand, he must make a <a href="https://www.d20pfsrd.com/magic#TOC-Concentration" rel="nofollow">concentration</a> check or lose the spell. The DC for this check is equal to 20 + the spell’s level. If the object is a ring or amulet, it occupies the ring or neck slot accordingly.</p><p>A bonded object can be used once per day to cast any one spell that the wizard has in his spellbook and is capable of casting, even if the spell is not prepared. This spell is treated like any other spell cast by the wizard, including casting time, duration, and other effects dependent on the wizard’s level. This spell cannot be modified by <a href="https://www.d20pfsrd.com/feats/metamagic-feats">metamagic feats</a> or other abilities. The bonded object cannot be used to cast spells from the wizard’s opposition schools (see <a href="https://www.d20pfsrd.com/classes/core-classes/wizard#TOC-Arcane-School"><b>arcane school</b></a> below).</p><p>A wizard can add additional magic abilities to his bonded object as if he has the required <a href="https://www.d20pfsrd.com/feats/item-creation-feats">Item Creation Feats</a> and if he meets the level prerequisites of the feat. For example, a wizard with a bonded dagger must be at least 5th level to add magic abilities to the dagger (see <a href="https://www.d20pfsrd.com/feats/item-creation-feats/craft-magic-arms-and-armor-item-creation">Craft Magic Arms and Armor</a> feat). If the bonded object is a wand, it loses its wand abilities when its last charge is consumed, but it is not destroyed and it retains all of its bonded object properties and can be used to craft a new wand. The magic properties of a bonded object, including any magic abilities added to the object, only function for the wizard who owns it. If a bonded object’s owner dies, or the item is replaced, the object reverts to being an ordinary masterwork item of the appropriate type.</p><p>If a bonded object is damaged, it is restored to full hit points the next time the wizard prepares his spells. If the object of an arcane bond is lost or destroyed, it can be replaced after 1 week in a special ritual that costs 200 gp per wizard level plus the cost of the masterwork item. This ritual takes 8 hours to complete. Items replaced in this way do not possess any of the additional enchantments of the previous bonded item. A wizard can designate an existing magic item as his bonded item. This functions in the same way as replacing a lost or destroyed item except that the new magic item retains its abilities while gaining the benefits and drawbacks of becoming a bonded item.</p><p>Instead of specializing in a focused arcane school of magic, a wizard can choose to specialize in one of the <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/elemental-arcane-schools">elemental schools of magic</a>. Like a normal arcane school, an elemental school grants a number of school powers and one bonus spell slot of each level the wizard can cast, from 1st on up. This bonus spell slot must be used to prepare a spell from the elemental school’s spell list. Unlike a normal arcane school, each elemental school requires the wizard to select his opposed element as his opposition school (<a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/elemental-arcane-schools/air">air</a> opposes <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/elemental-arcane-schools/earth">earth</a>, <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/elemental-arcane-schools/fire">fire</a> opposes <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/elemental-arcane-schools/water">water</a>). He does not need to select a second opposition school. He must expend two spell baseSlots to prepare a spell from his opposed elemental school as normal.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Arcane-Bond-Ex-or-Sp-",
			LevelRequired: 1,
			Special:       arcaneBond.ArcaneBond{},
		},
		{
			Name:          "Arcane School",
			Description:   `<p>A wizard can choose to specialize in one school of magic, gaining additional spells and powers based on that school. This choice must be made at 1st level, and once made, it cannot be changed. A wizard that does not select a school receives the <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/classic-arcane-schools/universalist">universalist</a> school instead.</p><p>A wizard that chooses to specialize in one school of magic must select two other schools as his opposition schools, representing knowledge sacrificed in one area of arcane lore to gain mastery in another. A wizard who prepares spells from his opposition schools must use two spell baseSlots of that level to prepare the spell. For example, a wizard with <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/classic-arcane-schools/evocation">evocation</a> as an opposition school must expend two of his available 3rd-level spell baseSlots to prepare a <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/f/fireball">fireball</a>. In addition, a specialist takes a –4 penalty on any <a href="https://www.d20pfsrd.com/skills#TOC-Skill-Checks" rel="nofollow">skill checks</a> made when crafting a magic item that has a spell from one of his opposition schools as a prerequisite. A <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/classic-arcane-schools/universalist">universalist</a> wizard can prepare spells from any school without restriction.</p><p>Each arcane school gives the wizard a number of school powers. In addition, specialist wizards receive an additional spell slot of each spell level he can cast, from 1st on up. Each day, a wizard can prepare a spell from his specialty school in that slot. This spell must be in the wizard’s spellbook. A wizard can select a spell modified by a metamagic feat to prepare in his school slot, but it uses up a higher-level spell slot. Wizards with the <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-schools/paizo-arcane-schools/classic-arcane-schools/universalist">universalist</a> school do not receive a school slot.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Arcane-School",
			LevelRequired: 1,
			Special:       []*WizardArcaneSchools{},
		},
		{
			Name:          "Cantrips",
			Description:   `<p>Wizards can prepare a number of cantrips, or 0-level spells, each day, as noted on <b>Table: Wizard</b> under “Spells per Day.” These spells are cast like any other spell, but they are not expended when cast and may be used again. A wizard can prepare a cantrip from an opposition school, but it uses up two of his available baseSlots (see below).</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Cantrips",
			LevelRequired: 1,
		},
		{
			Name:          "Scribe Scroll",
			Description:   `<p>At 1st level, a wizard gains <a href="https://www.d20pfsrd.com/feats/item-creation-feats/scribe-scroll-item-creation" rel="nofollow">Scribe Scroll</a> as a bonus feat.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#Scribe_Scroll",
			LevelRequired: 1,
			Special: &feats.Feat{
				Name:        feats.ScribeScroll,
				Type:        feats.ItemCreationFeats,
				Benefit:     "<p><b>Benefit</b>: You can create a <a href=\"https://www.d20pfsrd.com/magic-items/scrolls\">scroll</a> of any spell that you know. Scribing a scroll takes 2 hours if its base price is 250 gp or less, otherwise scribing a scroll takes 1 day for each 1,000 gp in its base price. To scribe a scroll, you must use up raw materials costing half of this base price.</p><p>See <a href=\"https://www.d20pfsrd.com/magic-items#TOC-Magic-Item-Creation\">magic item creation rules</a> for more information.</p>",
				Description: "You can create magic scrolls.",
				Source:      "Pathfinder RPG Core Rulebook. Copyright 2009, Paizo Publishing, LLC; Author: Jason Bulmahn, based on material by Jonathan Tweet, Monte Cook, and Skip Williams.",
				SourceLink:  "http://www.amazon.com/gp/product/1601251505/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601251505&linkCode=as2&tag=httpwwwd20pfs-20",
				Link:        "https://www.d20pfsrd.com/feats/item-creation-feats/scribe-scroll-item-creation",
				Extra:       nil,
				Prerequisites: []*prerequisites.Prerequisite{
					{
						Key:   prerequisites.CasterLevelPrerequisite,
						Value: 1,
					},
				},
			},
		},
		{
			Name:          "Bonus Feat (5)",
			Description:   `<p>At 5th, 10th, 15th, and 20th level, a wizard gains a bonus feat. At each such opportunity, he can choose a <a href="https://www.d20pfsrd.com/feats/metamagic-feats">metamagic feat</a>, an <a href="https://www.d20pfsrd.com/feats/item-creation-feats">item creation</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a>. The wizard must still meet all prerequisites for a bonus feat, including caster level minimums. These bonus feats are in addition to the feats that a character of any class gets from advancing levels. The wizard is not limited to the categories of <a href="https://www.d20pfsrd.com/feats/item-creation-feats">Item Creation Feats</a>, <a href="https://www.d20pfsrd.com/feats/metamagic-feats">Metamagic Feats</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a> when choosing those feats.</p><p>A wizard may also choose an <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/">Arcane Discovery</a> that he qualifies for in place of a bonus feat at these levels.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Bonus-Feats",
			LevelRequired: 5,
		},
		{
			Name:          "Bonus Feat (10)",
			Description:   `<p>At 5th, 10th, 15th, and 20th level, a wizard gains a bonus feat. At each such opportunity, he can choose a <a href="https://www.d20pfsrd.com/feats/metamagic-feats">metamagic feat</a>, an <a href="https://www.d20pfsrd.com/feats/item-creation-feats">item creation</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a>. The wizard must still meet all prerequisites for a bonus feat, including caster level minimums. These bonus feats are in addition to the feats that a character of any class gets from advancing levels. The wizard is not limited to the categories of <a href="https://www.d20pfsrd.com/feats/item-creation-feats">Item Creation Feats</a>, <a href="https://www.d20pfsrd.com/feats/metamagic-feats">Metamagic Feats</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a> when choosing those feats.</p><p>A wizard may also choose an <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/">Arcane Discovery</a> that he qualifies for in place of a bonus feat at these levels.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Bonus-Feats",
			LevelRequired: 10,
		},
		{
			Name:          "Bonus Feat (15)",
			Description:   `<p>At 5th, 10th, 15th, and 20th level, a wizard gains a bonus feat. At each such opportunity, he can choose a <a href="https://www.d20pfsrd.com/feats/metamagic-feats">metamagic feat</a>, an <a href="https://www.d20pfsrd.com/feats/item-creation-feats">item creation</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a>. The wizard must still meet all prerequisites for a bonus feat, including caster level minimums. These bonus feats are in addition to the feats that a character of any class gets from advancing levels. The wizard is not limited to the categories of <a href="https://www.d20pfsrd.com/feats/item-creation-feats">Item Creation Feats</a>, <a href="https://www.d20pfsrd.com/feats/metamagic-feats">Metamagic Feats</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a> when choosing those feats.</p><p>A wizard may also choose an <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/">Arcane Discovery</a> that he qualifies for in place of a bonus feat at these levels.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Bonus-Feats",
			LevelRequired: 15,
		},
		{
			Name:          "Bonus Feat (20)",
			Description:   `<p>At 5th, 10th, 15th, and 20th level, a wizard gains a bonus feat. At each such opportunity, he can choose a <a href="https://www.d20pfsrd.com/feats/metamagic-feats">metamagic feat</a>, an <a href="https://www.d20pfsrd.com/feats/item-creation-feats">item creation</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a>. The wizard must still meet all prerequisites for a bonus feat, including caster level minimums. These bonus feats are in addition to the feats that a character of any class gets from advancing levels. The wizard is not limited to the categories of <a href="https://www.d20pfsrd.com/feats/item-creation-feats">Item Creation Feats</a>, <a href="https://www.d20pfsrd.com/feats/metamagic-feats">Metamagic Feats</a>, or <a href="https://www.d20pfsrd.com/feats/general-feats/spell-mastery" rel="nofollow">Spell Mastery</a> when choosing those feats.</p><p>A wizard may also choose an <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/">Arcane Discovery</a> that he qualifies for in place of a bonus feat at these levels.</p>`,
			Link:          "https://www.d20pfsrd.com/classes/core-classes/Wizard/#TOC-Bonus-Feats",
			LevelRequired: 20,
		},
	}
}
