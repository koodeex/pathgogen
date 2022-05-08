package prerequisites

type PrerequisiteType string

const (
	FeatPrerequisite PrerequisiteType = "Feat"

	StrPrerequisite PrerequisiteType = "STR"
	DexPrerequisite PrerequisiteType = "DEX"
	ConPrerequisite PrerequisiteType = "CON"
	IntPrerequisite PrerequisiteType = "INT"
	WisPrerequisite PrerequisiteType = "WIS"
	ChaPrerequisite PrerequisiteType = "CHA"

	BabPrerequisite PrerequisiteType = "BAB"

	TraitPrerequisites         PrerequisiteType = "Trait"
	MultiplePrerequisites      PrerequisiteType = "MultiPrerequisite"
	ClassFeaturesPrerequisites PrerequisiteType = "ClassFeature"

	AcrobaticsPrerequisite             PrerequisiteType = "Acrobatics"
	AppraisePrerequisite               PrerequisiteType = "Appraise"
	BluffPrerequisite                  PrerequisiteType = "Bluff"
	ClimbPrerequisite                  PrerequisiteType = "Climb"
	CraftPrerequisite                  PrerequisiteType = "Craft"
	DiplomacyPrerequisite              PrerequisiteType = "Diplomacy"
	DisableDevicePrerequisite          PrerequisiteType = "DisableDevice"
	DisguisePrerequisite               PrerequisiteType = "Disguise"
	EscapeArtistPrerequisite           PrerequisiteType = "EscapeArtist"
	FlyPrerequisite                    PrerequisiteType = "Fly"
	HandleAnimalPrerequisite           PrerequisiteType = "HandleAnimal"
	HealPrerequisite                   PrerequisiteType = "Heal"
	IntimidatePrerequisite             PrerequisiteType = "Intimidate"
	KnowledgeArcanaPrerequisite        PrerequisiteType = "KnowledgeArcana"
	KnowledgeDungeoneeringPrerequisite PrerequisiteType = "KnowledgeDungeoneering"
	KnowledgeEngineeringPrerequisite   PrerequisiteType = "KnowledgeEngineering"
	KnowledgeGeographyPrerequisite     PrerequisiteType = "KnowledgeGeography"
	KnowledgeHistoryPrerequisite       PrerequisiteType = "KnowledgeHistory"
	KnowledgeLocalPrerequisite         PrerequisiteType = "KnowledgeLocal"
	KnowledgeNaturePrerequisite        PrerequisiteType = "KnowledgeNature"
	KnowledgeNobilityPrerequisite      PrerequisiteType = "KnowledgeNobility"
	KnowledgePlanesPrerequisite        PrerequisiteType = "KnowledgePlanes"
	KnowledgeReligionPrerequisite      PrerequisiteType = "KnowledgeReligion"
	LinguisticsPrerequisite            PrerequisiteType = "Linguistics"
	PerceptionPrerequisite             PrerequisiteType = "Perception"
	PerformPrerequisite                PrerequisiteType = "Perform"
	ProfessionPrerequisite             PrerequisiteType = "Profession"
	RidePrerequisite                   PrerequisiteType = "Ride"
	SenseMotivePrerequisite            PrerequisiteType = "SenseMotive"
	SleightOfHandPrerequisite          PrerequisiteType = "SleightOfHand"
	SpellcraftPrerequisite             PrerequisiteType = "Spellcraft"
	StealthPrerequisite                PrerequisiteType = "Stealth"
	SurvivalPrerequisite               PrerequisiteType = "Survival"
	SwimPrerequisite                   PrerequisiteType = "Swim"
	UseMagicDevicePrerequisite         PrerequisiteType = "UseMagicDevice"
)

type Prerequisite struct {
	Key   PrerequisiteType
	Value interface{}
}
