package arcaneDiscovery

import (
	"github.com/koodeex/pathgogen/internal/models/feats"
	"github.com/koodeex/pathgogen/internal/models/feats/arcaneBond"
	"github.com/koodeex/pathgogen/internal/models/prerequisites"
)

const ArcaneDiscovery feats.FeatType = "Arcane Discovery"

var (
	AlchemicalAffinity = &feats.Feat{
		Name:        "Alchemical Affinity",
		Type:        ArcaneDiscovery,
		Description: "Having studied alongside alchemists, you’ve learned to use their methodologies to enhance your spellcraft.",
		Benefit:     `Whenever you cast a spell that appears on both the <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> and <a href="https://www.d20pfsrd.com/classes/base-classes/alchemist">alchemist</a> spell lists, you treat your <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> as 1 higher than normal and the save DC of such spells increases by 1. Additionally, you may copy spells from an <a href="https://www.d20pfsrd.com/classes/base-classes/alchemist">alchemist’s</a> formula book into your spellbook just as you could with another <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard’s</a> spellbook.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/alchemical-affinity/",
		Source:      "Pathfinder Player Companion: Magical Marketplace © 2013, Paizo Publishing, LLC; Authors: John Ling, Ron Lundeen, Patrick Renie, David Schwartz, and Jerome Virnich.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256000/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256000&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	ArcaneBuilder = &feats.Feat{
		Name:        "Arcane Builder",
		Type:        ArcaneDiscovery,
		Description: "You have an exceptional understanding of the theory behind creating magical items.",
		Benefit:     `Select one type of magic item (<a href="https://www.d20pfsrd.com/magic-items/potions">potions</a>, <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items">wondrous items</a>, and so on). You create items of this type 25% faster than normal, and gain a +4 bonus on <a href="https://www.d20pfsrd.com/skills/spellcraft">Spellcraft</a> checks (or other checks, as appropriate) to craft items of this type.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-builder/",
		Source:      "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink:  "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Special:     `You may select this <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo">discovery</a> multiple times; its effects do not stack. Each time you select this discovery, it applies to a different type of magic item.`,
	}
	BalancedSummoning = &feats.Feat{
		Name:        "Balanced Summoning",
		Type:        ArcaneDiscovery,
		Description: "You have an exceptional understanding of the theory behind creating magical items.",
		Benefit:     `You maintain balance by calling on opposing forces when <a href="https://www.d20pfsrd.com/magic#TOC-Conjuration-Summoning">summoning</a>. Whenever you cast a <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/s/summon-monster">summon monster</a> spell, you can <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/universal-monster-rules#TOC-Summon-Sp-">summon</a> two creatures from a single list 1 or more levels lower than the level of the spell. The two creatures must have alignments that are opposite along at least one axis (<a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Chaotic">chaotic</a> and <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Lawful">lawful</a> or evil and good). For example, if you cast <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/s/summon-monster">summon monster III</a>, you could <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/universal-monster-rules#TOC-Summon-Sp-">summon</a> a <a href="https://www.d20pfsrd.com/bestiary/monster-listings/templates/celestial-creature-cr-special">celestial</a> <a href="https://www.d20pfsrd.com/bestiary/monster-listings/undead/ghoul-wolf-1">wolf</a> and a <a href="https://www.d20pfsrd.com/bestiary/monster-listings/templates/fiendish">fiendish</a> <a href="https://www.d20pfsrd.com/bestiary/monster-listings/animals/hyena">hyena</a> from the 2nd-level list.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/balanced-summoning/",
		Source:      "Pathfinder Player Companion: Champions of Balance © 2014, Paizo Publishing, LLC; Authors: Matt Goodall, Ron Lundeen, Philip Minchin, Patrick Renie, Jason Ridler, and David Schwartz.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256035/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256035&linkCode=as2&tag=httpwwwd20pfs-20",
	}
	BeyondMorality = &feats.Feat{
		Name:       "Beyond Morality (Su)",
		Type:       ArcaneDiscovery,
		Benefit:    `As long as you are neutral, you may choose to be treated as the most favorable <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Alignment">alignment</a> when affected by spells whose effects vary based on <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Alignment">alignment</a> (such as <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/h/holy-word">holy word</a>). If you are neutral in relation to evil and good, you may choose to be treated as good or evil. If you are neutral in relation to chaos and law, you may choose to be treated as lawful or chaotic. You may only choose to be treated as one <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Alignment">alignment</a> type along a single axis at a time (for instance, if you were within the area of both a <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/m/magic-circle-against-evil">magic circle against evil</a> spell and an <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/u/unholy-blight">unholy blight</a> spell, you would have to choose to be either evil, good, or neutral for the purpose of determining the spells’ effects).`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/beyond-morality-su/",
		Source:     "Pathfinder Player Companion: Champions of Balance © 2014, Paizo Publishing, LLC; Authors: Matt Goodall, Ron Lundeen, Philip Minchin, Patrick Renie, Jason Ridler, and David Schwartz.",
		SourceLink: "http://www.amazon.com/gp/product/1601256035/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256035&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 9,
			},
		},
	}
	CreativeDestruction = &feats.Feat{
		Name:        "Creative Destruction (Su)",
		Type:        ArcaneDiscovery,
		Description: "You have learned how to use destructive energy to empower yourself.",
		Benefit:     `When you cast an <a href="https://www.d20pfsrd.com/magic#TOC-Evocation">evocation</a> spell that deals damage, you gain a number of <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Temporary-Hit-Points">temporary hit points</a> equal to the total number of <a href="https://www.d20pfsrd.com/equipment/goods-and-services/toys-games-puzzles#TOC-Dice">dice</a> used to determine the damage caused by the spell. Temporary <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Hit-Points">hit points</a> gained from this discovery do not stack and disappear after 1 hour.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/creative-destruction-su/",
		Source:      "Pathfinder Player Companion: Champions of Balance © 2014, Paizo Publishing, LLC; Authors: Matt Goodall, Ron Lundeen, Philip Minchin, Patrick Renie, Jason Ridler, and David Schwartz.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256035/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256035&linkCode=as2&tag=httpwwwd20pfs-20",
	}
	DefensiveFeedback = &feats.Feat{
		Name:        "Defensive Feedback (Su)",
		Type:        ArcaneDiscovery,
		Description: "Rather than dissipate damaging energy, you can redirect some of it back to its source.",
		Benefit:     `When an <a href="https://www.d20pfsrd.com/magic#TOC-Abjuration">abjuration</a> spell you cast prevents damage (with <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Damage-Reduction">damage reduction</a> or <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Energy-Resistance">energy resistance</a>), if the attacking creature is within 30 feet of the protected creature, the foe takes 1d6 points of damage for every 10 points of damage prevented.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/defensive-feedback-su/",
		Source:      "Pathfinder Player Companion: Champions of Balance © 2014, Paizo Publishing, LLC; Authors: Matt Goodall, Ron Lundeen, Philip Minchin, Patrick Renie, Jason Ridler, and David Schwartz.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256035/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256035&linkCode=as2&tag=httpwwwd20pfs-20",
	}
	FaithMagic = &feats.Feat{
		Name:       "Faith Magic",
		Type:       ArcaneDiscovery,
		Benefit:    `Select one spell granted by a domain belonging to the god you worship. This spell must be at least 2 levels lower than the highest-level <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> spell you can cast. When you first prepare your spells for the day, you can prepare this spell once, using a spell slot 1 level higher than the spell’s actual level. This is cast as a divine spell.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/faith-magic/",
		Source:     "Pathfinder Player Companion: Magic Tactics Toolbox © 2016, Paizo Inc.; Authors: Alexander Augunas, Steven T. Helt, Thurston Hillman, and Ron Lundeen.",
		SourceLink: "https://www.amazon.com/gp/product/1601258380/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&tag=httpwwwd20pfs-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=1601258380&linkId=ee3470c7ae99ce605c107092f42710fa",
	}
	FastStudy = &feats.Feat{
		Name:       "Fast Study",
		Type:       ArcaneDiscovery,
		Benefit:    `Normally, a wizard spends 1 hour preparing all of his spells for the day, or proportionately less if he only prepares some spells, with a minimum of 15 minutes of preparation. Thanks to mental discipline and clever mnemonics, you can prepare all of your spells in only 15 minutes, and your minimum preparation time is only 1 minute.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/fast-study/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	FeralSpeech = &feats.Feat{
		Name:       "Feral Speech",
		Type:       ArcaneDiscovery,
		Benefit:    `You gain the <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Supernatural-Abilities-Su-">supernatural ability</a> to speak with and understand the response of any <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Animal">animal</a> as if using <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/s/speak-with-animals">speak with animals</a>, though each time you speak to <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Animal">animals</a>, you must decide to communicate with either amphibians, birds, fish, mammals, or reptiles, and can only speak to and understand animals of that type. You can make yourself understood as far as your voice carries. This <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo">discovery</a> does not predispose any animal addressed toward you in any way. When you reach 12th level, you can also use this ability to communicate with <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Vermin">vermin</a>.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/feral-speech/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	ForestsBlessing = &feats.Feat{
		Name:       "Forest’s Blessing",
		Type:       ArcaneDiscovery,
		Benefit:    `You cast any spells that appear on both the <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> and <a href="https://www.d20pfsrd.com/classes/core-classes/druid">druid</a> spell lists at +1 <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> and with +1 to the save DC. In addition, you may replace the material component of any arcane spell with gems of the same value.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/forest-s-blessing/",
		Source:     "Pathfinder Player Companion: Champions of Purity © 2013, Paizo Publishing, LLC; Authors: Jessica Blomstrom, Adam Daigle, Shaun Hocking, Daniel Marthaler, Tork Shaw, and Christina Stiles.",
		SourceLink: "http://www.amazon.com/gp/product/160125511X/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=160125511X&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	GolemConstructor = &feats.Feat{
		Name:       "Golem Constructor",
		Type:       ArcaneDiscovery,
		Benefit:    `You have learned the art and craft of creating a single type of <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem">golem</a> (such as <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem/golem-stone">stone golems</a> or <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem/golem-iron">iron golems</a>). When creating a <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem">golem</a> of this type, you count as having the <a href="https://www.d20pfsrd.com/feats/item-creation-feats/craft-wondrous-item-item-creation">Craft Wondrous Item</a>, <a href="https://www.d20pfsrd.com/feats/item-creation-feats/craft-magic-arms-and-armor-item-creation">Craft Magic Arms and Armor</a>, and <a href="https://www.d20pfsrd.com/feats/item-creation-feats/craft-construct-item-creation">Craft Construct</a> feats. You must meet all other construction requirements for the <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem">golem</a> as normal.`,
		Special:    `You may select this <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo">discovery</a> multiple times. Each time you select this <a href="https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo">discovery</a>, it applies to a different kind of <a href="https://www.d20pfsrd.com/bestiary/monster-listings/constructs/golem">golem</a>.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/golem-constructor/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 9,
			},
		},
	}
	Idealize = &feats.Feat{
		Name:        "Idealize (Su)",
		Type:        ArcaneDiscovery,
		Description: "In your quest for self-perfection, you have discovered a way to further enhance yourself and others.",
		Benefit:     `When a <a href="https://www.d20pfsrd.com/magic#TOC-Transmutation">transmutation</a> spell you cast grants an <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Enhancement-Bonus">enhancement bonus</a> to an ability score, that bonus increases by 2. At 20th level, the bonus increases by 4.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/idealize-su/",
		Source:      "Pathfinder Player Companion: Champions of Balance © 2014, Paizo Publishing, LLC; Authors: Matt Goodall, Ron Lundeen, Philip Minchin, Patrick Renie, Jason Ridler, and David Schwartz.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256035/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256035&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 10,
			},
		},
	}
	Immortality = &feats.Feat{
		Name:       "Immortality",
		Type:       ArcaneDiscovery,
		Benefit:    `You discover a cure for aging, and from this point forward you take no penalty to your physical ability scores from advanced age. If you are already taking such penalties, they are removed at this time. This is an <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Extraordinary-Abilities-Ex-">extraordinary ability</a>.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/immortality/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 20,
			},
		},
	}
	InfectiousCharms = &feats.Feat{
		Name:        "Infectious Charms",
		Type:        ArcaneDiscovery,
		Description: "Your charms are so smooth that they’re contagious.",
		Benefit:     `Anytime you target and successfully affect a single creature with a <a href="https://www.d20pfsrd.com/magic#TOC-Enchantment-Charm">charm</a> or <a href="https://www.d20pfsrd.com/magic#TOC-Enchantment-Compulsion">compulsion</a> spell and that creature is within 30 feet of another opponent, your spell has a chance of affecting the second creature as well. As a <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Swift-Actions">swift action</a> immediately after affecting a creature with a <a href="https://www.d20pfsrd.com/magic#TOC-Enchantment-Charm">charm</a> or <a href="https://www.d20pfsrd.com/magic#TOC-Enchantment-Compulsion">compulsion</a> spell, you can cause the spell to carry over to the nearest creature within 30 feet. The spell behaves in all ways as though its new target were the original target of the spell.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/infectious-charms/",
		Source:      "Pathfinder Player Companion: Magical Marketplace © 2013, Paizo Publishing, LLC; Authors: John Ling, Ron Lundeen, Patrick Renie, David Schwartz, and Jerome Virnich.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256000/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256000&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 11,
			},
		},
	}
	IounBond = &feats.Feat{
		Name:       "Ioun Bond",
		Type:       ArcaneDiscovery,
		Benefit:    `You can form an arcane bond with an <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a>. If you choose this arcane discovery at 1st level, you gain a dull gray <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a> as a bonded object at no cost. A bonded <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a> must be orbiting your head to have effect. At 12th level, you can turn a bonded dull gray <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a> into another kind of <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a> as if you possessed the <a href="https://www.d20pfsrd.com/feats/item-creation-feats/craft-wondrous-item-item-creation">Craft Wondrous Item</a> feat; if you die or replace a bonded <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a> that has been transformed in this way, the stone reverts to a dull gray <a href="https://www.d20pfsrd.com/magic-items/wondrous-items/wondrous-items/h-l/ioun-stones" class="spell">ioun stone</a>.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/ioun-bond/",
		Source:     "Pathfinder Player Companion: People of the River © 2014, Paizo Inc.; Authors: Tim Akers, Jason Brick, Ethan Day-Jones, James Jacobs, Nick Salestrom, David Schwartz, and William Thrasher.",
		SourceLink: "http://www.amazon.com/gp/product/1601256663/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256663&linkCode=as2&tag=httpwwwd20pfs-20&linkId=GVEQEMVO7TRZKFQA",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.ArcaneBondPrerequisite,
				Value: arcaneBond.BondedObject,
			},
		},
	}
	KnowledgeIsPower = &feats.Feat{
		Name:       "Knowledge is Power (Ex)",
		Type:       ArcaneDiscovery,
		Benefit:    `Your understanding of physical forces gives you power over them. You add your <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Intelligence-Int-">Intelligence</a> modifier on <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Combat-Maneuvers">combat maneuver</a> checks and to your <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Combat-Maneuver-Defense">CMD</a>. You also add your <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Intelligence-Int-">Intelligence</a> modifier on <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Strength-Str-">Strength</a> checks to break or lift objects.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/ioun-bond/",
		Source:     "Pathfinder Player Companion: People of the River © 2014, Paizo Inc.; Authors: Tim Akers, Jason Brick, Ethan Day-Jones, James Jacobs, Nick Salestrom, David Schwartz, and William Thrasher.",
		SourceLink: "http://www.amazon.com/gp/product/1601256663/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256663&linkCode=as2&tag=httpwwwd20pfs-20&linkId=GVEQEMVO7TRZKFQA",
	}
	Multimorph = &feats.Feat{
		Name:        "Multimorph",
		Type:        ArcaneDiscovery,
		Description: "Your studies in transmogrification have increased your control over shapechanging spells.",
		Benefit:     `When you cast a spell of the <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/p/polymorph">polymorph</a> subschool on yourself, you may expend 1 minute of the spell’s duration as a <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Standard-Actions">standard action</a> to assume another form allowed by the spell. You can do this as often as you like, subject to the duration of the spell. This is a <a href="https://www.d20pfsrd.com/gamemastering/special-abilities#TOC-Supernatural-Abilities-Su-">supernatural ability</a>.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/multimorph/",
		Source:      "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink:  "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	ObservantIllusion = &feats.Feat{
		Name:        "Observant Illusion (Su)",
		Type:        ArcaneDiscovery,
		Description: "The following wizard arcane discovery is available to elven wizards and wizard servants of a specific organization in place of a feat.",
		Benefit:     `You can project your senses into any ongoing <a href="https://www.d20pfsrd.com/magic#TOC-lllusion-Figment">figment</a> or shadow <a href="https://www.d20pfsrd.com/magic#TOC-Illusion">illusion</a> you create with a spell of at least 3rd level. You can see through its eyes and hear through its ears as if you were standing where it is, and during your turn you can switch from using its senses to using your own, or back again, as a swift or <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Move-Actions">move action</a>. While you are using its senses, your body is considered <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Blinded">blinded</a> and <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Deafened">deafened</a>.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/observant-illusion-su/",
		Source:      "Pathfinder Player Companion: Spymaster’s Handbook © 2016, Paizo Inc.; Authors: Alexander Augunas, David N. Ross, and Owen K.C. Stephens.",
		SourceLink:  "https://www.amazon.com/gp/product/1601258445/ref=as_li_tl?ie=UTF8&tag=httpwwwd20pfs-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=1601258445&linkId=6ed72bcd2ace3e794050b798e0c8346e",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 9,
			},
		},
	}
	OppositionResearch = &feats.Feat{
		Name:       "Opposition Research",
		Type:       ArcaneDiscovery,
		Benefit:    `Select one Wizard opposition school; preparing spells of this school now only requires one spell slot of the appropriate level instead of two, and you no longer have the –4 <a href="https://www.d20pfsrd.com/skills/spellcraft">Spellcraft</a> penalty for crafting items from that school.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/opposition-research/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 9,
			},
		},
	}
	PsychicPreparation = &feats.Feat{
		Name:        "Psychic Preparation",
		Type:        ArcaneDiscovery,
		Description: "You have learned a limited way to access psychic magic.",
		Benefit:     `When you first prepare your spells for the day, you can prepare one spell as a <a href="https://www.d20pfsrd.com/alternative-rule-systems/occult-adventures/occult-classes/psychic">psychic</a> spell. This spell must be at least 2 levels lower than the highest-level <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> spell you can cast, and takes a slot 1 level higher than the spell’s actual level. When you cast this spell, it operates as a <a href="https://www.d20pfsrd.com/alternative-rule-systems/occult-adventures/occult-classes/psychic">psychic</a> spell, including using emotional and <a href="https://www.d20pfsrd.com/alternative-rule-systems/occult-adventures/psychic-magic#thought-components">thought components</a>&nbsp;in place of somatic and verbal <a href="https://www.d20pfsrd.com/magic#TOC-Components">components</a>, and only having expensive material <a href="https://www.d20pfsrd.com/magic#TOC-Components">components</a>.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/psychic-preparation/",
		Source:      "Pathfinder Player Companion: Magic Tactics Toolbox © 2016, Paizo Inc.; Authors: Alexander Augunas, Steven T. Helt, Thurston Hillman, and Ron Lundeen.",
		SourceLink:  "https://www.amazon.com/gp/product/1601258380/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&tag=httpwwwd20pfs-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=1601258380&linkId=ee3470c7ae99ce605c107092f42710fa",
	}
	ResilientIllusions = &feats.Feat{
		Name:        "Resilient Illusions",
		Type:        ArcaneDiscovery,
		Description: "You are able to conjure illusions so lifelike that they defy disbelief.",
		Benefit:     `Anytime a creature tries to disbelieve one of your <a href="https://www.d20pfsrd.com/magic#TOC-Illusion">illusion</a> effects, make a <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> check. Treat the illusion’s save DC as its normal DC or the result of the <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> check, whichever is higher.`,
		Link:        "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/resilient-illusions/",
		Source:      "Pathfinder Player Companion: Magical Marketplace © 2013, Paizo Publishing, LLC; Authors: John Ling, Ron Lundeen, Patrick Renie, David Schwartz, and Jerome Virnich.",
		SourceLink:  "http://www.amazon.com/gp/product/1601256000/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256000&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 8,
			},
		},
	}
	SplitSlot = &feats.Feat{
		Name:       "Split Slot",
		Type:       ArcaneDiscovery,
		Benefit:    `Once per day when you prepare spells, you may treat any one of your open spell slots as if it were two spell slots that were two spell levels lower. For example, a 9th-level Wizard can split a 5th-level slot into two 3rd-level slots, preparing <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/f/fireball">fireball</a> and <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/l/lightning-bolt">lightning bolt</a> in those 3rd-level slots. For all purposes, the two lower-level slots are treated as that lower level (so the split 5th-level slot used for a <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/f/fireball">fireball</a> has a DC as if it were in a normal 3rd-level slot). Splitting a 2nd-level slot lets you prepare two additional cantrips (which you can cast over and over, just like normally prepared cantrips). This discovery has no effect on cantrips or 1st–level spells.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/split-slot/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Special:    "You may select this discovery multiple times; each time you select it, you may split another spell slot when you prepare spells. You cannot split a slot that you created by splitting a higher-level slot.",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 5,
			},
		},
	}
	StaffLikeWand = &feats.Feat{
		Name:       "Staff-Like Wand",
		Type:       ArcaneDiscovery,
		Benefit:    `Similar to using a magic staff, you use your own <a href="https://www.d20pfsrd.com/basics-ability-scores/ability-scores#TOC-Intelligence-Int-">Intelligence</a> score and relevant feats to set the DC for saves against spells you cast from a wand, and you can use your <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> when activating the power of a wand if it’s higher than the <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> of the wand.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/staff-like-wand/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "http://www.amazon.com/gp/product/1601252994/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601252994&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 11,
			},
			{
				Key:   prerequisites.FeatPrerequisite,
				Value: feats.CraftStaff,
			},
		},
	}
	StewardOfTheGreatBeyond = &feats.Feat{
		Name:       "Steward of the Great Beyond",
		Type:       ArcaneDiscovery,
		Benefit:    `Whenever a creature attempts to use a <a href="https://www.d20pfsrd.com/magic#TOC-Conjuration-Teleportation">teleportation</a> effect or <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/universal-monster-rules#TOC-Summon-Sp-">summon</a> a creature within 30 feet of you, you may attempt to block the effect. Make an opposed <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a> check (1d20 + <a href="https://www.d20pfsrd.com/magic#TOC-Caster-Level">caster level</a>) as an <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Immediate-Actions">immediate action</a>. If the check succeeds, the spell or effect fails and is wasted; otherwise, it is unaffected. You can use this ability once per day plus one additional time for every 5 <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> levels you possess beyond 10th.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/resilient-illusions/",
		Source:     "Pathfinder Player Companion: Champions of Purity © 2013, Paizo Publishing, LLC; Authors: Jessica Blomstrom, Adam Daigle, Shaun Hocking, Daniel Marthaler, Tork Shaw, and Christina Stiles.",
		SourceLink: "http://www.amazon.com/gp/product/160125511X/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=160125511X&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 9,
			},
		},
	}
	TimeStutter = &feats.Feat{
		Name:       "Time Stutter (Sp)",
		Type:       ArcaneDiscovery,
		Benefit:    `You can briefly step out of time, pausing the world around you. This ability acts as the <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/t/time-stop">time stop</a> spell, except that you gain only 1 round of apparent time. You can use this ability once per day plus one additional time for every 5 <a href="https://www.d20pfsrd.com/classes/core-classes/wizard">wizard</a> levels you possess beyond 10th.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/time-stutter-sp/",
		Source:     "Pathfinder Player Companion: People of the River © 2014, Paizo Inc.; Authors: Tim Akers, Jason Brick, Ethan Day-Jones, James Jacobs, Nick Salestrom, David Schwartz, and William Thrasher.",
		SourceLink: "http://www.amazon.com/gp/product/1601256663/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601256663&linkCode=as2&tag=httpwwwd20pfs-20&linkId=GVEQEMVO7TRZKFQA",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 10,
			},
		},
	}
	TrueName = &feats.Feat{
		Name:       "True Name (Sp)",
		Type:       ArcaneDiscovery,
		Benefit:    `Your researches into ancient tomes and your inquisitions of bound spirits have led you to one of the best-hidden secrets of the multiverse: the true name of an <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsider</a>—the name that defines the very essence of the creature and that gives the speaker control over the being. This <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsider</a> can have no more than 12 <a href="https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Hit-Dice-HD-">Hit Dice</a>. Once per day, you can speak the common name by which the <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsider</a> is known, and the <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsider</a> travels to you as if you had cast <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/p/planar-binding">planar binding</a> upon it. It must obey you to the best of its ability, without pay or bargaining for its services, for its fear that you might release its true name to the wider world is enough to bring even the most recalcitrant of <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsiders</a> to bear. If the creature is within 100 feet, as a <a href="https://www.d20pfsrd.com/gamemastering/combat#TOC-Move-Actions">move action</a>, you may punish it by deliberately mispronouncing its name, wracking its very essence and giving it the <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Sickened">sickened</a> and <a href="https://www.d20pfsrd.com/gamemastering/conditions#TOC-Staggered">staggered</a> conditions for 1 round (even if the creature is normally immune to these conditions). You cannot use true name in an area of <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/s/silence">silence</a>, but the creature does not have to be able to hear you for it to be harmed by the ability. It is in your best interest to call this creature only sparingly, and occasionally reward it in some fashion to mollify its wrath. If you repeatedly fail to offer it a reward appropriate to its type and ethos, the creature may begin plotting ways to destroy the bond between you, whether by creating an accident that will destroy your memory of the name, by plaguing you with nuisances or dangers until you vow never to call on it again, or by actively seeking to destroy you through its own devices or those of an underling. If this creature is of a <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Lawful">lawful</a> type and you are violating its ethos, its superiors may even destroy it or you rather than allow you to contaminate their servant further. Worse, they may establish situations where it is necessary for you to <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/universal-monster-rules#TOC-Summon-Sp-">summon</a> this <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider">outsider</a>, opening gateways to infernal or angelic interference, in order to gain a foothold on the Material Plane.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/true-name/",
		Source:     "Pathfinder Roleplaying Game: Ultimate Magic. © 2011, Paizo Publishing, LLC; Authors: Jason Bulmahn, Tim Hitchcock, Colin McComb, Rob McCreary, Jason Nelson, Stephen Radney-MacFarland, Sean K Reynolds, Owen K.C. Stephens, and Russ Taylor.",
		SourceLink: "https://www.opengamingstore.com/products/ultimate-magic-pathfinder-roleplaying-game",
		Special:    "You may select this <a href=\"https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo\">discovery</a> multiple times. Each time you select this <a href=\"https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/arcane-discoveries-paizo\">discovery</a>, it applies to a different, specific <a href=\"https://www.d20pfsrd.com/bestiary/rules-for-monsters/creature-types#TOC-Outsider\">outsider</a>. If you select this discovery at 15th level or higher, the creature may have up to 18 <a href=\"https://www.d20pfsrd.com/basics-ability-scores/glossary#TOC-Hit-Dice-HD-\">Hit Dice</a> and the call acts as <a class=\"spell\" href=\"https://www.d20pfsrd.com/magic/all-spells/p/planar-binding\">greater planar binding</a> instead of <a class=\"spell\" href=\"https://www.d20pfsrd.com/magic/all-spells/p/planar-binding\">planar binding</a>.",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 11,
			},
		},
	}
	WerewolfShape = &feats.Feat{
		Name:       "Werewolf Shape",
		Type:       ArcaneDiscovery,
		Benefit:    `When you cast <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/b/beast-shape">beast shape IV</a> or <a class="spell" href="https://www.d20pfsrd.com/magic/all-spells/s/shapechange">shapechange</a>, you can choose to take the shape of a <a href="https://www.d20pfsrd.com/bestiary/monster-listings/humanoids/lycanthrope/werewolf">werewolf</a> in addition to the other shapes available. While in <a href="https://www.d20pfsrd.com/bestiary/monster-listings/humanoids/lycanthrope/werewolf">werewolf</a> shape, you gain all the effects of the <a href="https://www.d20pfsrd.com/bestiary/monster-listings/humanoids/lycanthrope/werewolf">werewolf</a> template instead of the spell’s normal benefits. You act in all respects as a natural <a href="https://www.d20pfsrd.com/bestiary/monster-listings/templates/lycanthrope">lycanthrope</a> for the duration of the spell, including the ability to inflict the <a href="https://www.d20pfsrd.com/bestiary/rules-for-monsters/universal-monster-rules#TOC-Curse-of-Lycanthropy-Su-">curse of lycanthropy</a> using the spell’s save DC.`,
		Link:       "https://www.d20pfsrd.com/classes/core-classes/wizard/arcane-discoveries/arcane-discoveries-paizo/werewolf-shape/",
		Source:     "Pathfinder Player Companion: Blood of the Moon © 2013, Paizo Publishing, LLC; Authors: Tim Akers, Neal Litherland, David N. Ross, and Tork Shaw.",
		SourceLink: "http://www.amazon.com/gp/product/1601255780/ref=as_li_qf_sp_asin_il_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1601255780&linkCode=as2&tag=httpwwwd20pfs-20",
		Prerequisites: []*prerequisites.Prerequisite{
			{
				Key:   prerequisites.WizardLevelPrerequisite,
				Value: 15,
			},
		},
	}
)

var (
	ArcaneDiscoveriesList = []*feats.Feat{
		AlchemicalAffinity,
		ArcaneBuilder,
		BalancedSummoning,
		BeyondMorality,
		CreativeDestruction,
		DefensiveFeedback,
		FaithMagic,
		FastStudy,
		FeralSpeech,
		ForestsBlessing,
		GolemConstructor,
		Idealize,
		Immortality,
		InfectiousCharms,
		IounBond,
		KnowledgeIsPower,
		Multimorph,
		ObservantIllusion,
		OppositionResearch,
		PsychicPreparation,
		ResilientIllusions,
		SplitSlot,
		StaffLikeWand,
		StewardOfTheGreatBeyond,
		TimeStutter,
		TrueName,
		WerewolfShape,
	}
)
