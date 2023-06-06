package play

import (
	"math/rand"

	"juanantoniocid/snake/internal/geometry"
	"juanantoniocid/snake/internal/play/characters"
)

type board struct {
	width  int
	height int

	snake *characters.Snake
	apple *characters.Apple

	snakeDirection geometry.Direction
	snakeSpeed     int
	timer          int
}

func newBoard(width, height int) *board {
	b := &board{
		width:          width,
		height:         height,
		snakeDirection: geometry.DirNone,
		snakeSpeed:     4,
		timer:          0,
	}

	b.initSnake()
	b.initApple()

	return b
}

// GetSnakeShape returns the snake shape
func (b *board) GetSnakeShape() geometry.Shape {
	return b.snake.GetShape()
}

// GetAppleShape returns the apple shape
func (b *board) GetAppleShape() geometry.Shape {
	return b.apple.GetShape()
}

// SetSpeed sets the speed of the snake
func (b *board) SetSpeed(speed int) {
	b.snakeSpeed = speed
}

// MoveSnake moves the snake in the given direction
func (b *board) MoveSnake(dir geometry.Direction) Status {
	b.timer++
	return b.moveSnake(dir)
}

func (b *board) moveSnake(dir geometry.Direction) Status {
	var forceAdvance bool
	if dir != b.snakeDirection && b.turnAllowed(dir) {
		forceAdvance = true
		b.snakeDirection = dir
	}

	return b.advanceSnake(forceAdvance)
}

func (b *board) turnAllowed(dir geometry.Direction) bool {
	if dir == geometry.DirLeft {
		return b.snakeDirection != geometry.DirRight
	} else if dir == geometry.DirRight {
		return b.snakeDirection != geometry.DirLeft
	} else if dir == geometry.DirDown {
		return b.snakeDirection != geometry.DirUp
	} else if dir == geometry.DirUp {
		return b.snakeDirection != geometry.DirDown
	}
	return false
}

func (b *board) advanceSnake(force bool) Status {
	if b.needsToMoveSnake() || force {
		b.snake.Move(b.snakeDirection)

		if b.snakeCollidesWithWall() || b.snakeCollidesWithSelf() {
			return StatusGameOver
		}

		if b.snakeEatsApple() {
			b.initApple()
			b.snake.Grow()
			return StatusSnakeEating
		}

	}
	return StatusPlaying
}

func (b *board) initApple() {
	b.apple = characters.NewApple(rand.Intn(b.width-1), rand.Intn(b.height-1))
}

func (b *board) initSnake() {
	b.snake = characters.NewSnake(b.width/2, b.height/2)
}

func (b *board) needsToMoveSnake() bool {
	return b.timer%b.snakeSpeed == 0
}

func (b *board) snakeCollidesWithWall() bool {
	snakeHead := b.snake.GetShape()[0]
	return snakeHead.X < 0 || snakeHead.Y < 0 ||
		snakeHead.X >= b.width || snakeHead.Y >= b.height
}

func (b *board) snakeCollidesWithSelf() bool {
	snakeShape := b.snake.GetShape()
	snakeHead := snakeShape[0]
	snakeTail := snakeShape[1:]

	for _, v := range snakeTail {
		if snakeHead.X == v.X && snakeHead.Y == v.Y {
			return true
		}
	}
	return false
}

func (b *board) snakeEatsApple() bool {
	appleShape := b.apple.GetShape()[0]
	snakeHead := b.snake.GetShape()[0]
	return snakeHead.X == appleShape.X && snakeHead.Y == appleShape.Y
}
