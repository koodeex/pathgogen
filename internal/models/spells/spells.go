package spells

type SpellComponents string

const (
	Verbal   SpellComponents = "V"
	Somatic  SpellComponents = "S"
	Material SpellComponents = "M"
	DF       SpellComponents = "DF"
)

type SpellSchool string

const (
	Abjuration SpellSchool = "abjuration"
)

type RequiredLevel struct {
	Name  string
	Level int
}

type Spell struct {
	Name            string
	School          SpellSchool
	CastingTime     string
	Circle          int
	Components      []SpellComponents
	Description     string
	Range           int
	Target          string
	Duration        string
	SavingThrow     string
	SpellResistance bool
	Link            string
	Source          string
	SourceLink      string
}
