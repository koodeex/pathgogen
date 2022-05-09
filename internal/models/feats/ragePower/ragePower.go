package ragePower

import "github.com/koodeex/pathgogen/internal/models/prerequisites"

type RagePower struct {
	Name          string
	Benefit       string
	Prerequisites []*prerequisites.Prerequisite
}
