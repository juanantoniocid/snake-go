package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/snake/internal/geometry"
	"juanantoniocid/snake/internal/play"
)

const (
	ScreenWidth          = 640
	ScreenHeight         = 480
	gridSize     float32 = 10
)

type Game struct {
	play *play.Play

	bestScore int
}

func (g *Game) Update() error {
	err := g.iterate()
	if err != nil {
		return err
	}

	if g.bestScore < g.play.GetScore() {
		g.bestScore = g.play.GetScore()
	}

	return nil
}

func (g *Game) iterate() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.play.Reset()
		return nil
	}

	dir := geometry.DirNone
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dir = geometry.DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dir = geometry.DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir = geometry.DirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir = geometry.DirRight
	}
	return g.play.MoveSnake(dir)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSnake(screen)
	g.drawApple(screen)

	if g.play.GetStatus() == play.StatusInitial {
		ebitenutil.DebugPrint(screen, "Press up/down/left/right to start")
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Level: %d Score: %d Best Score: %d", ebiten.ActualFPS(), g.play.GetLevel(), g.play.GetScore(), g.bestScore))
	}
}

func (g *Game) drawApple(screen *ebiten.Image) {
	apple := g.play.GetAppleShape()
	for _, a := range apple {
		vector.DrawFilledRect(screen, float32(a.X)*gridSize, float32(a.Y)*gridSize, gridSize, gridSize,
			color.RGBA{R: 0xFF, A: 0xff}, false)
	}
}

func (g *Game) drawSnake(screen *ebiten.Image) {
	snake := g.play.GetSnakeShape()
	for _, v := range snake {
		vector.DrawFilledRect(
			screen, float32(v.X)*gridSize, float32(v.Y)*gridSize, gridSize, gridSize,
			color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff}, false)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		play: play.NewPlay(ScreenWidth/int(gridSize), ScreenHeight/int(gridSize)),
	}

	return g
}
