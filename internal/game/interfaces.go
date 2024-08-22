package game

// Интерфейс для передвижения
type Mover interface {
	MoveEntity(e *Entity, w *World)
}

// Интерфейс для скрещивания
type Сrossbreeder interface {
	Crossover(p *Population, parent1, parent2 Entity) *Entity
}
