package game

import (
	"github.com/t3mnikov/genetical/internal/utils"
	"image/color"
	"math/rand"
)

// Разные скрещиватели и их логика
type MediumCrossover struct {
}

type TemperatureCrossover struct {
}

func NewMediumCrossover() *MediumCrossover {
	return &MediumCrossover{}
}

func NewTemperatureCrossover() *TemperatureCrossover {
	return &TemperatureCrossover{}
}

// Умеренное скрещивание
func (c *MediumCrossover) Crossover(p *Population, parent1, parent2 Entity) *Entity {
	x, y := 0, 0

	if rand.Float64() < 0.25 {
		x = parent2.X
		y = parent2.Y
	} else {
		x = rand.Intn(p.world.width)
		y = rand.Intn(p.world.height)
	}

	child := NewEntity(x, y, MaxEnergy/2, color.White, NewAdultMover())
	child.SetGreed((parent1.greed + parent2.greed) / 2)
	child.SetSpeed((parent1.speed + parent2.speed) / 2)

	// мутации
	if rand.Float64() < 0.1 {
		child.speed += rand.Intn(5) - 1 // Изменение скорости
		if child.speed < 1 {
			child.speed = 1
		}
	}

	if rand.Float64() < 0.1 {
		child.greed += rand.Intn(5) - 1 // Изменение жадности
		if child.greed < 1 {
			child.greed = 1
		}
	}

	// Случайное движение ребенка в сторону
	child.X = utils.FixVal(child.X+rand.Intn(9)-4, 0, p.world.width-1)
	child.Y = utils.FixVal(child.Y+rand.Intn(9)-4, 0, p.world.height-1)

	return &child
}

// Скрещивание в зависмости от температурных условий
func (c *TemperatureCrossover) Crossover(p *Population, parent1, parent2 Entity) *Entity {
	// Проверка зоны
	temperature := p.world.temperature[parent1.X][parent1.Y]
	var reproductionChance float64

	switch temperature {
	case Hot:
		reproductionChance = 0.9 // высокий шанс
	case Neutral:
		reproductionChance = 0.5 // обычный шанс
	case Cold:
		reproductionChance = 0.0 // нулевой шанс
	}

	if rand.Float64() < reproductionChance {
		x := rand.Intn(p.world.width)
		y := rand.Intn(p.world.height)

		child := NewEntity(x, y, MaxEnergy, color.White, NewAdultMover())
		child.SetSpeed((parent1.speed + parent2.speed) / 2)
		child.SetGreed((parent1.greed + parent2.greed) / 2)

		return &child
	}

	return nil
}
