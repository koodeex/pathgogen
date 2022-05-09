package arcaneBond

type ArcaneBondType string

const (
	BondedObject = "Bonded Object"
	Familiar     = "Familiar"
)

type ArcaneBond struct {
	Name    string
	Type    ArcaneBondType
	Special interface{}
}
