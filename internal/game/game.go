package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth     = 800
	ScreenHeight    = 800
	TileSize        = 4   // Размер элемента
	FoodCount       = 200 // Начальное количество еды
	FoodRespawn     = 14  // Скорость восполнения еды, меньше - быстрее
	MaxEntities     = 500 // Максимум сущностей в мире
	StartPopulation = 40  // Начальное количество популяции особей
)

const (
	MaxEnergy       = 400 // Максимальная энергия особей
	MinEnergy       = 0   // Минимальная энергия
	MaxVerminEnergy = 200 // Максимальная энергия вредителя
)

const (
	MaxFoodEnergy = 250 // Энергия, даваемая особи от еды
)

type Game struct {
	population *Population
	world      *World
}

func NewGame() *Game {
	world := NewWorld(ScreenWidth/TileSize, ScreenHeight/TileSize)
	population := NewPopulation(StartPopulation, world)

	return &Game{
		population: population,
		world:      world,
	}
}

// Обновление кадра
func (g *Game) Update() error {
	g.population.Update()

	return nil
}

// Отрисовка объектов
func (g *Game) Draw(screen *ebiten.Image) {
	// Отрисовка еды
	g.world.DrawFood(screen)

	// Отрисовка популяции
	g.population.Draw(screen)
}

// Размеры экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
