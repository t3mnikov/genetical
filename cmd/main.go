package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t3mnikov/genetical/internal/game"
	"log"
)

func main() {
	// Запуск симулятора
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("GeneticAl")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
