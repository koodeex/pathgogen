package traits

type Category string

const (
	Campaign = "Campaign"
	Race     = "Race"
)

type Trait struct {
	Name       string
	Category   string
	Benefit    string
	Additional string
	Special    interface{}
}
