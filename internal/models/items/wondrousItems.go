package items

type WondrousItem struct {
	Name                     string
	Description              string
	Cost                     int
	Weight                   int
	ConstructionRequirements string
	Special                  interface{}
}
