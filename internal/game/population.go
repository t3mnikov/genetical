package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

// Популяция особей
type Population struct {
	entities     []Entity
	world        *World
	crossbreader Сrossbreeder // скрещиватель
}

// Конструктор новой популяции
func NewPopulation(size int, world *World) *Population {
	entities := make([]Entity, size)

	for i := range entities {
		entities[i] = NewEntity(rand.Intn(world.width), rand.Intn(world.height), MaxEnergy, color.White, NewAdultMover())
	}

	return &Population{
		entities:     entities,
		world:        world,
		crossbreader: NewMediumCrossover(),
	}
}

// Обновление кадра
func (p *Population) Update() {
	p.world.Update()

	p.Move()
	p.Crossing()
	p.Harming()
}

// Движение популяции
func (p *Population) Move() {
	for i := 0; i < len(p.entities); i++ {
		p.entities[i].Move(p.world)
		// Если энергия объекта падает ниже минимума, он умирает
		if p.entities[i].energy <= MinEnergy {
			p.entities = append(p.entities[:i], p.entities[i+1:]...)
			i--
		}
	}
}

// Процесс скрещивания
func (p *Population) Crossing() {
	newEntities := []Entity{}

	if len(p.entities) < MaxEntities {
		for i := 0; i < len(p.entities); i++ {
			for j := i + 1; j < len(p.entities); j++ {
				if p.entities[i].DistanceTo(&p.entities[j]) <= 4 {
					child := p.Crossover(p.entities[i], p.entities[j])
					if child != nil {
						newEntities = append(newEntities, *child)
					}

					// Ограничиваем количество потомков за один цикл
					if len(newEntities) > 2 { // например, не более 2 новых особей за цикл
						break
					}
				}
			}
			// Если достигнут лимит потомков за цикл, прекращаем создание новых особей
			if len(newEntities) > 2 {
				break
			}
		}
	}

	if len(p.entities) == MaxEntities-1 {
		child := p.BornVermin()
		newEntities = append(newEntities, *child)
	}

	p.entities = append(p.entities, newEntities...)
}

// Отрисовка популяции
func (p *Population) Draw(screen *ebiten.Image) {
	for _, entity := range p.entities {
		// Цвет для сущностей
		vector.DrawFilledRect(screen, float32(entity.X*TileSize), float32(entity.Y*TileSize), TileSize, TileSize, entity.color, true)
	}
}

// Скрещивание
func (p *Population) Crossover(parent1, parent2 Entity) *Entity {
	return p.crossbreader.Crossover(p, parent1, parent2)
}

// Процесс причинения вреда
func (p *Population) Harming() {
	for i := 0; i < len(p.entities); i++ {
		for j := i + 1; j < len(p.entities); j++ {
			if p.entities[i].DistanceTo(&p.entities[j]) <= 4 {

				if !p.entities[i].vermin {
					return
				}

				p.entities[j].Hit(100)
			}
		}
	}
}

// Добавление новых особей
func (p *Population) AddEntities(size int) {
	entities := make([]Entity, size)

	for i := range entities {
		entities[i] = NewEntity(rand.Intn(p.world.width), rand.Intn(p.world.height), MaxEnergy, color.White, NewAdultMover())
	}

	p.entities = append(p.entities, entities...)
}

// BornVermin порождает вредителя
func (p *Population) BornVermin() *Entity {
	x := rand.Intn(p.world.width)
	y := rand.Intn(p.world.height)

	child := NewEntity(x, y, MaxVerminEnergy, color.RGBA{255, 0, 0, 255}, NewAdultMover())
	child.SetGreed(1)
	child.SetSpeed(1)
	child.SetVermin(true)

	return &child
}
