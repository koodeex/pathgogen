package feats

type AbilityType string

const (
	Ex   AbilityType = "Ex"
	Su   AbilityType = "Su"
	Sp   AbilityType = "Sp"
	ExSp AbilityType = "Ex or Sp"
)

type ClassFeature struct {
	Name          string
	AbilityType   AbilityType
	Description   string
	Link          string
	LevelRequired int
	Special       interface{}
}
