package races

type SizeCategory struct {
	Name     string
	Modifier int
}

var (
	Colossal = &SizeCategory{
		Name:     "Colossal",
		Modifier: -8,
	}
	Gargantuan = &SizeCategory{
		Name:     "Gargantuan",
		Modifier: -4,
	}
	Huge = &SizeCategory{
		Name:     "Huge",
		Modifier: -2,
	}
	Large = &SizeCategory{
		Name:     "Large",
		Modifier: -1,
	}
	Medium = &SizeCategory{
		Name:     "Medium",
		Modifier: 0,
	}
	Small = &SizeCategory{
		Name:     "Small",
		Modifier: 1,
	}
	Tine = &SizeCategory{
		Name:     "Tine",
		Modifier: 2,
	}
	Diminitive = &SizeCategory{
		Name:     "Diminitive",
		Modifier: 4,
	}
	Fine = &SizeCategory{
		Name:     "Fine",
		Modifier: 8,
	}
)

type Race struct {
	Name         string
	SizeCategory *SizeCategory
}

func NewRace(name string, category *SizeCategory) *Race {
	return &Race{
		Name:         name,
		SizeCategory: category,
	}
}
