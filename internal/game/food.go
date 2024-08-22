package game

// Еда для особей
type Food struct {
	X      int
	Y      int
	energy int
	eaten  bool
}

// Конструтор еды
func NewFood(x int, y int, energy int) *Food {
	return &Food{X: x, Y: y, energy: energy, eaten: false}
}
