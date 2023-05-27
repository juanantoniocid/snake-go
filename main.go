package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"juanantoniocid/snake/internal/game"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Snake [Zippo Games]")

	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
