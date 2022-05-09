package feats

import (
	"github.com/koodeex/pathgogen/internal/models/prerequisites"
)

type FeatType string

const (
	CombatFeats          FeatType = "Combat Feats"
	BloodHexFeats        FeatType = "Blood Hex Feats"
	GeneralFeats         FeatType = "General Feats"
	ConduitFeats         FeatType = "Conduit Feats"
	CriticalFeats        FeatType = "Critical Feats"
	AnimalCompanionFeats FeatType = "Animal Companion Feats"
	DamnationFeats       FeatType = "Damnation Feats"
	FactionFeats         FeatType = "Faction Feats"
	GritandPanacheFeats  FeatType = "Grit and Panache Feats"
	HeroPointFeats       FeatType = "Hero Point Feats"
	ItemCreationFeats    FeatType = "Item Creation Feats"
	ItemMasteryFeats     FeatType = "Item Mastery Feats"
	MeditationFeats      FeatType = "Meditation Feats"
	MetamagicFeats       FeatType = "Metamagic Feats"
	MythicFeats          FeatType = "Mythic Feats"
	PerformanceFeats     FeatType = "Performance Feats"
	RacialFeats          FeatType = "Racial Feats"
	StareFeats           FeatType = "Stare Feats"
	StoryFeats           FeatType = "Story Feats"
	StyleFeats           FeatType = "Style Feats"
	AchievementFeats     FeatType = "Achievement Feats"
	TargetingFeats       FeatType = "Targeting Feats"
	AnimalFamiliarFeats  FeatType = "Animal/Familiar Feats"
	TeamworkFeats        FeatType = "Teamwork Feats"
	MonsterFeats         FeatType = "Monster Feats"
)

type Feat struct {
	Name          string
	Type          FeatType
	Benefit       string
	Description   string
	Normal        string
	Special       string
	Source        string
	SourceLink    string
	Link          string
	Extra         []string
	Prerequisites []*prerequisites.Prerequisite
}
