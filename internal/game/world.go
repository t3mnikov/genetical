package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

// Компература и их значения
type Temperature int

const (
	Cold Temperature = iota
	Neutral
	Hot
)

// Мир
type World struct {
	width, height int
	food          [][]*Food
	foodTimer     int
	temperature   [][]Temperature
}

// Конструктор мира
func NewWorld(width, height int) *World {
	food := make([][]*Food, width)
	for i := range food {
		food[i] = make([]*Food, height)
	}

	// Разбрасываем еду по миру
	for i := 0; i < FoodCount; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		e := rand.Intn(MaxFoodEnergy)
		food[x][y] = NewFood(x, y, e)
	}

	return &World{
		width:       width,
		height:      height,
		food:        food,
		temperature: generateTemperature(width, height),
	}
}

// Отрисовка еды
func (w *World) DrawFood(screen *ebiten.Image) {
	for x := 0; x < w.width; x++ {
		for y := 0; y < w.height; y++ {
			if f := w.food[x][y]; f != nil && !f.eaten {
				// Green color for food
				vector.DrawFilledRect(screen, float32(x*TileSize), float32(y*TileSize), TileSize, TileSize, color.RGBA{0, 255, 0, 255}, true)
			}
		}
	}
}

// Обновление еды
func (w *World) Update() {
	w.foodTimer++

	if w.foodTimer >= FoodRespawn {
		w.RespawnFood()
		// Сброс таймера еды
		w.foodTimer = 0
	}
}

// Появление новой еды
func (w *World) RespawnFood() {
	x := rand.Intn(w.width)
	y := rand.Intn(w.height)
	e := rand.Intn(MaxFoodEnergy)
	w.food[x][y] = NewFood(x, y, e)
}

// Генерирование температуры в мире
func generateTemperature(width, height int) [][]Temperature {
	temperature := make([][]Temperature, width)

	for x := 0; x < width; x++ {
		temperature[x] = make([]Temperature, height)
		for y := 0; y < height; y++ {
			// распределение температурных зон
			r := rand.Float64()
			if r < 0.7 {
				temperature[x][y] = Cold
			} else if r < 0.8 {
				temperature[x][y] = Neutral
			} else {
				temperature[x][y] = Hot
			}
		}
	}

	return temperature
}
