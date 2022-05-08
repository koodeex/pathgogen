package feats

import "pathgogen/models/prerequisites"

type AbilityType string

const (
	Ex AbilityType = "Ex"
)

type ClassFeature struct {
	Name          string
	AbilityType   AbilityType
	Description   string
	Link          string
	LevelRequired int
	Special       interface{}
}

type RagePower struct {
	Name          string
	Benefit       string
	Prerequisites []*prerequisites.Prerequisite
}
