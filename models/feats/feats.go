package feats

import "pathgogen/models/prerequisites"

type Feat struct {
	Name          string
	Type          string
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