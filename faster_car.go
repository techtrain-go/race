package race

type FasterCar struct {
	maps string
}

func NewFasterCar(m string) *FasterCar {
	return &FasterCar{
		maps: m,
	}
}

func (c *FasterCar) Go(start, finish string) []string {
	// ваш код должен быть тут

	return []string{}
}
