package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"juanantoniocid/snake/internal/play"

	"juanantoniocid/snake/internal/direction"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	gridSize     = 10
)

type Game struct {
	play *play.Play
}

func (g *Game) Update() error {
	dir := direction.DirNone

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dir = direction.DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dir = direction.DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir = direction.DirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir = direction.DirRight
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.play.Reset()
	}

	return g.play.MoveSnake(dir)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSnake(screen)
	g.drawApple(screen)

	if g.play.MoveDirection == direction.DirNone {
		ebitenutil.DebugPrint(screen, "Press up/down/left/right to start")
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Level: %d Score: %d Best Score: %d", ebiten.ActualFPS(), g.play.Level, g.play.Score, g.play.BestScore))
	}
}

func (g *Game) drawApple(screen *ebiten.Image) {
	apple := g.play.Apple.GetPosition()

	vector.DrawFilledRect(
		screen,
		float32(apple.X*gridSize),
		float32(apple.Y*gridSize),
		float32(gridSize),
		float32(gridSize),
		color.RGBA{R: 0xFF, A: 0xff},
		false,
	)
}

func (g *Game) drawSnake(screen *ebiten.Image) {
	snake := g.play.Snake.GetBody()
	for _, v := range snake {
		vector.DrawFilledRect(
			screen,
			float32(v.X*gridSize),
			float32(v.Y*gridSize),
			float32(gridSize),
			float32(gridSize),
			color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff},
			false,
		)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		play: play.NewPlay(ScreenWidth/gridSize, ScreenHeight/gridSize),
	}

	return g
}
