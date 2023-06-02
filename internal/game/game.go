package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/snake/internal/geometry"
	"juanantoniocid/snake/internal/play"
)

const (
	width                  = 320
	height                 = 200
	gridSize       float32 = 4
	sizeMultiplier         = 3
)

type Game struct {
	play *play.Play

	bestScore int
}

func (g *Game) Update() error {
	g.iterate()
	if g.bestScore < g.play.GetScore() {
		g.bestScore = g.play.GetScore()
	}
	return nil
}

func (g *Game) iterate() {
	dir := geometry.DirNone

	if g.play.GetStatus() == play.StatusGameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.play.Reset()
			dir = geometry.DirUp
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.play.Reset()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dir = geometry.DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dir = geometry.DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir = geometry.DirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir = geometry.DirRight
	}
	g.play.MoveSnake(dir)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSnake(screen)
	g.drawApple(screen)

	switch g.play.GetStatus() {
	case play.StatusInitial:
		ebitenutil.DebugPrint(screen, "Press any arrow key to start.")
	case play.StatusPlaying:
		ebitenutil.DebugPrint(screen, fmt.Sprintf(
			"Level %d : Score %d : Best Score: %d", g.play.GetLevel(), g.play.GetScore(), g.bestScore))
	case play.StatusGameOver:
		ebitenutil.DebugPrint(screen, "Game Over. Enter to restart or Q to quit.")
	}
}

func (g *Game) drawApple(screen *ebiten.Image) {
	apple := g.play.GetAppleShape()
	for _, a := range apple {
		vector.DrawFilledRect(screen, float32(a.X)*gridSize, float32(a.Y)*gridSize, gridSize, gridSize,
			color.RGBA{R: 0xaa, A: 0xff}, false)
	}
}

func (g *Game) drawSnake(screen *ebiten.Image) {
	snake := g.play.GetSnakeShape()
	for _, v := range snake {
		vector.DrawFilledRect(
			screen, float32(v.X)*gridSize, float32(v.Y)*gridSize, gridSize, gridSize,
			color.RGBA{G: 0xaa, A: 0xff}, false)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return width, height
}

func NewGame() *Game {
	g := &Game{
		play: play.NewPlay(width/int(gridSize), height/int(gridSize)),
	}

	return g
}

func Run() {
	ebiten.SetWindowSize(width*sizeMultiplier, height*sizeMultiplier)
	ebiten.SetWindowTitle("Snake (Zippo Studios)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
