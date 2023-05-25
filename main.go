package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"juanantoniocid/snake/internal/game"
	"log"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Snake (Ebitengine Demo)")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
