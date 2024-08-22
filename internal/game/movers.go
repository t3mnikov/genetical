package game

import (
	"github.com/t3mnikov/genetical/internal/utils"
	"math/rand"
)

// Передвигатель особей
type AdultMover struct {
}

func NewAdultMover() *AdultMover {
	return &AdultMover{}
}

// Передвигает особь
func (m *AdultMover) MoveEntity(e *Entity, world *World) {

	if e.greed > rand.Intn(300) {
		// Поиск ближайшей еды
		for dx := -2; dx <= 2; dx++ {
			for dy := -2; dy <= 2; dy++ {
				nx, ny := e.X+dx, e.Y+dy
				// Проверяем, что новые координаты внутри границ мира
				if nx >= 0 && nx < world.width && ny >= 0 && ny < world.height && world.food[nx][ny] != nil && !world.food[nx][ny].eaten {
					e.X, e.Y = nx, ny

					f := world.food[nx][ny]
					e.energy += f.energy
					if e.energy > MaxEnergy {
						e.energy = MaxEnergy
					}

					world.food[nx][ny].eaten = true
					return
				}
			}
		}
	}

	// Простейшее случайное движение
	for i := 0; i < e.speed; i++ {
		dir := rand.Intn(4)
		switch dir {
		case 0:
			if e.Y > 0 {
				e.Y -= rand.Intn(2)
			}
		case 1:
			if e.Y < world.height-1 {
				e.Y += rand.Intn(2)
			}
		case 2:
			if e.X > 0 {
				e.X -= rand.Intn(2)
			}
		case 3:
			if e.X < world.width-1 {
				e.X += rand.Intn(2)
			}
		}

		e.X = utils.FixVal(e.X, 0, world.width-1)
		e.Y = utils.FixVal(e.Y, 0, world.height-1)
	}

	// Когда особи двигаются, то они теряют энергию
	e.energy--
}
