package game

import (
	"image/color"
	"math/rand"
)

// Особь
type Entity struct {
	X, Y   int
	energy int
	mover  Mover
	color  color.Color
	speed  int // Ген скорости
	greed  int // Ген жадности
	vermin bool
}

// Конструктор особей
func NewEntity(x, y int, energy int, color color.Color, mover Mover) Entity {
	return Entity{
		X:      x,
		Y:      y,
		energy: energy,
		mover:  mover,
		color:  color,
		speed:  rand.Intn(3) + 1,
		greed:  rand.Intn(3) + 1,
	}
}

// Задает цвета
func (e *Entity) SetColor(color color.Color) {
	e.color = color
}

// Задает ген скорости
func (e *Entity) SetSpeed(speed int) {
	e.speed = speed
}

// Задает ген жадности
func (e *Entity) SetGreed(greed int) {
	e.greed = greed
}

// Если true, то делает особь вредителем
func (e *Entity) SetVermin(vermin bool) {
	e.vermin = vermin
}

// Движение особи
func (e *Entity) Move(world *World) {
	e.mover.MoveEntity(e, world)
}

// Нанесение удара
func (e *Entity) Hit(value int) {
	e.energy -= value
}

// Вычисляет дистанцию до другой особи (без квадратного корня)
func (e *Entity) DistanceTo(other *Entity) int {
	dx := e.X - other.X
	dy := e.Y - other.Y

	return dx*dx + dy*dy
}
