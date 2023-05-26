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

// MoveSnake moves the snake in the given direction
func (p *Play) MoveSnake(dir geometry.Direction) error {
	p.setSnakeDirection(dir)
	p.moveSnake()
	p.timer++

	return nil
}

func (p *Play) setSnakeDirection(dir geometry.Direction) {
	if dir == geometry.DirLeft {
		if p.moveDirection != geometry.DirRight {
			p.moveDirection = geometry.DirLeft
		}
	} else if dir == geometry.DirRight {
		if p.moveDirection != geometry.DirLeft {
			p.moveDirection = geometry.DirRight
		}
	} else if dir == geometry.DirDown {
		if p.moveDirection != geometry.DirUp {
			p.moveDirection = geometry.DirDown
		}
	} else if dir == geometry.DirUp {
		if p.moveDirection != geometry.DirDown {
			p.moveDirection = geometry.DirUp
		}
	}
}

func (p *Play) moveSnake() {
	if p.needsToMoveSnake() {
		if p.snakeCollidesWithWall() || p.snakeCollidesWithSelf() {
			p.Reset()
		}

		if p.snakeCollidesWithApple() {
			p.initApple()
			p.snake.Grow()
			if p.snake.Len() > 10 && p.snake.Len() < 20 {
				p.level = 2
				p.moveTime = 3
			} else if p.snake.Len() > 20 {
				p.level = 3
				p.moveTime = 2
			} else {
				p.level = 1
			}
			p.score++
		}

		p.snake.Move(p.moveDirection)
	}
}

func (p *Play) snakeCollidesWithWall() bool {
	snakeHead := p.snake.GetShape()[0]
	return snakeHead.X < 0 ||
		snakeHead.Y < 0 ||
		snakeHead.X >= p.boardWidth ||
		snakeHead.Y >= p.boardHeight
}

func (p *Play) snakeCollidesWithSelf() bool {
	snakeShape := p.snake.GetShape()
	snakeHead := snakeShape[0]
	snakeTail := snakeShape[1:]
	for _, v := range snakeTail {
		if snakeHead.X == v.X &&
			snakeHead.Y == v.Y {
			return true
		}
	}
	return false
}

func (p *Play) snakeCollidesWithApple() bool {
	appleShape := p.apple.GetShape()[0]
	snakeHead := p.snake.GetShape()[0]
	return snakeHead.X == appleShape.X &&
		snakeHead.Y == appleShape.Y
}

func (p *Play) needsToMoveSnake() bool {
	return p.timer%p.moveTime == 0
}
