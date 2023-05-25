package game

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"juanantoniocid/snake/internal/apple"
	"juanantoniocid/snake/internal/direction"
	"juanantoniocid/snake/internal/snake"
)

const (
	ScreenWidth        = 640
	ScreenHeight       = 480
	gridSize           = 10
	xGridCountInScreen = ScreenWidth / gridSize
	yGridCountInScreen = ScreenHeight / gridSize
)

type Game struct {
	moveDirection int
	snake         *snake.Snake
	apple         *apple.Apple
	timer         int
	moveTime      int
	score         int
	bestScore     int
	level         int
}

func (g *Game) collidesWithApple() bool {
	applePos := g.apple.GetPosition()
	snakePos := g.snake.GetHead()
	return snakePos.X == applePos.X &&
		snakePos.Y == applePos.Y
}

func (g *Game) collidesWithSelf() bool {
	head := g.snake.GetHead()
	tail := g.snake.GetTail()
	for _, v := range tail {
		if head.X == v.X &&
			head.Y == v.Y {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithWall() bool {
	head := g.snake.GetHead()
	return head.X < 0 ||
		head.Y < 0 ||
		head.X >= xGridCountInScreen ||
		head.Y >= yGridCountInScreen
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) reset() {
	g.initApple()
	g.initSnake()

	g.moveTime = 4
	g.score = 0
	g.level = 1
	g.moveDirection = direction.DirNone
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.moveDirection != direction.DirRight {
			g.moveDirection = direction.DirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.moveDirection != direction.DirLeft {
			g.moveDirection = direction.DirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.moveDirection != direction.DirUp {
			g.moveDirection = direction.DirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.moveDirection != direction.DirDown {
			g.moveDirection = direction.DirUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.reset()
	}

	if g.needsToMoveSnake() {
		if g.collidesWithWall() || g.collidesWithSelf() {
			g.reset()
		}

		if g.collidesWithApple() {
			g.initApple()
			g.snake.Grow()
			if g.snake.Len() > 10 && g.snake.Len() < 20 {
				g.level = 2
				g.moveTime = 3
			} else if g.snake.Len() > 20 {
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

		g.snake.Move(g.moveDirection)
	}

	g.timer++

	return nil
}

func (g *Game) initApple() {
	g.apple = apple.NewApple(rand.Intn(xGridCountInScreen-1), rand.Intn(yGridCountInScreen-1), gridSize)
}

func (g *Game) initSnake() {
	g.snake = snake.NewSnake(xGridCountInScreen/2, yGridCountInScreen/2, gridSize)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)
	g.apple.Draw(screen)

	if g.moveDirection == direction.DirNone {
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
		moveTime: 4,
	}
	g.initApple()
	g.initSnake()
	return g
}
