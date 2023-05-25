package game

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/snake/internal/apple"
	"juanantoniocid/snake/internal/position"
)

const (
	ScreenWidth        = 640
	ScreenHeight       = 480
	gridSize           = 10
	xGridCountInScreen = ScreenWidth / gridSize
	yGridCountInScreen = ScreenHeight / gridSize
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

type Game struct {
	moveDirection int
	snakeBody     []position.Position
	apple         *apple.Apple
	timer         int
	moveTime      int
	score         int
	bestScore     int
	level         int
}

func (g *Game) collidesWithApple() bool {
	applePos := g.apple.GetPosition()
	return g.snakeBody[0].X == applePos.X &&
		g.snakeBody[0].Y == applePos.Y
}

func (g *Game) collidesWithSelf() bool {
	for _, v := range g.snakeBody[1:] {
		if g.snakeBody[0].X == v.X &&
			g.snakeBody[0].Y == v.Y {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithWall() bool {
	return g.snakeBody[0].X < 0 ||
		g.snakeBody[0].Y < 0 ||
		g.snakeBody[0].X >= xGridCountInScreen ||
		g.snakeBody[0].Y >= yGridCountInScreen
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) reset() {
	g.setNewApple()
	g.moveTime = 4
	g.snakeBody = g.snakeBody[:1]
	g.snakeBody[0].X = xGridCountInScreen / 2
	g.snakeBody[0].Y = yGridCountInScreen / 2
	g.score = 0
	g.level = 1
	g.moveDirection = dirNone
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.moveDirection != dirRight {
			g.moveDirection = dirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.moveDirection != dirLeft {
			g.moveDirection = dirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.moveDirection != dirUp {
			g.moveDirection = dirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.moveDirection != dirDown {
			g.moveDirection = dirUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.reset()
	}

	if g.needsToMoveSnake() {
		if g.collidesWithWall() || g.collidesWithSelf() {
			g.reset()
		}

		if g.collidesWithApple() {
			g.setNewApple()
			g.snakeBody = append(g.snakeBody, position.Position{
				X: g.snakeBody[len(g.snakeBody)-1].X,
				Y: g.snakeBody[len(g.snakeBody)-1].Y,
			})
			if len(g.snakeBody) > 10 && len(g.snakeBody) < 20 {
				g.level = 2
				g.moveTime = 3
			} else if len(g.snakeBody) > 20 {
				g.level = 3
				g.moveTime = 2
			} else {
				g.level = 1
			}
			g.score++
			if g.bestScore < g.score {
				g.bestScore = g.score
			}
		}

		for i := int64(len(g.snakeBody)) - 1; i > 0; i-- {
			g.snakeBody[i].X = g.snakeBody[i-1].X
			g.snakeBody[i].Y = g.snakeBody[i-1].Y
		}
		switch g.moveDirection {
		case dirLeft:
			g.snakeBody[0].X--
		case dirRight:
			g.snakeBody[0].X++
		case dirDown:
			g.snakeBody[0].Y++
		case dirUp:
			g.snakeBody[0].Y--
		}
	}

	g.timer++

	return nil
}

func (g *Game) setNewApple() {
	g.apple = apple.NewApple(rand.Intn(xGridCountInScreen-1), rand.Intn(yGridCountInScreen-1), gridSize)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.snakeBody {
		vector.DrawFilledRect(screen, float32(v.X*gridSize), float32(v.Y*gridSize), gridSize, gridSize, color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff}, false)
	}

	g.apple.Draw(screen)

	if g.moveDirection == dirNone {
		ebitenutil.DebugPrint(screen, "Press up/down/left/right to start")
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Level: %d Score: %d Best Score: %d", ebiten.ActualFPS(), g.level, g.score, g.bestScore))
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		moveTime:  4,
		snakeBody: make([]position.Position, 1),
	}
	g.snakeBody[0].X = xGridCountInScreen / 2
	g.snakeBody[0].Y = yGridCountInScreen / 2
	g.setNewApple()
	return g
}
