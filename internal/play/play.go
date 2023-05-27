package play

import (
	"math/rand"

	"juanantoniocid/snake/internal/geometry"
	"juanantoniocid/snake/internal/play/characters"
)

type Status int

const (
	StatusInitial Status = iota
	StatusPlaying
	StatusGameOver
)

// Play is the main game struct
type Play struct {
	boardWidth  int
	boardHeight int

	snake *characters.Snake
	apple *characters.Apple

	status        Status
	timer         int
	moveDirection geometry.Direction
	moveTime      int
	score         int
	level         int
}

// NewPlay creates a new game
func NewPlay(width, height int) *Play {
	g := &Play{
		boardWidth:  width,
		boardHeight: height,
	}
	g.Reset()

	return g
}

// Reset resets the game
func (p *Play) Reset() {
	p.initApple()
	p.initSnake()

	p.status = StatusInitial
	p.timer = 0
	p.moveDirection = geometry.DirNone
	p.moveTime = 4
	p.score = 0
	p.level = 1
}

func (p *Play) initApple() {
	p.apple = characters.NewApple(rand.Intn(p.boardWidth-1), rand.Intn(p.boardHeight-1))
}

func (p *Play) initSnake() {
	p.snake = characters.NewSnake(p.boardWidth/2, p.boardHeight/2)
}

// GetStatus returns the current game status
func (p *Play) GetStatus() Status {
	return p.status
}

// GetScore returns the current game score
func (p *Play) GetScore() int {
	return p.score
}

// GetLevel returns the current game level
func (p *Play) GetLevel() int {
	return p.level
}

// GetSnakeShape returns the current snake shape
func (p *Play) GetSnakeShape() geometry.Shape {
	return p.snake.GetShape()
}

// GetAppleShape returns the current apple shape
func (p *Play) GetAppleShape() geometry.Shape {
	return p.apple.GetShape()
}
