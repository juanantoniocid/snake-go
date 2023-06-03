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

	moveDirection geometry.Direction
	timer         int
	moveTime      int
}

func newBoard(width, height int) *board {
	b := &board{
		width:         width,
		height:        height,
		moveDirection: geometry.DirNone,
		timer:         0,
		moveTime:      4,
	}

	b.initSnake()
	b.initApple()

	return b
}

func (b *board) MoveSnake(dir geometry.Direction) Status {
	b.timer++
	b.setSnakeDirection(dir)
	return b.moveSnake()
}

func (b *board) setSnakeDirection(dir geometry.Direction) {
	if dir == geometry.DirLeft {
		if b.moveDirection != geometry.DirRight {
			b.moveDirection = geometry.DirLeft
		}
	} else if dir == geometry.DirRight {
		if b.moveDirection != geometry.DirLeft {
			b.moveDirection = geometry.DirRight
		}
	} else if dir == geometry.DirDown {
		if b.moveDirection != geometry.DirUp {
			b.moveDirection = geometry.DirDown
		}
	} else if dir == geometry.DirUp {
		if b.moveDirection != geometry.DirDown {
			b.moveDirection = geometry.DirUp
		}
	}
}

func (b *board) moveSnake() Status {
	if b.needsToMoveSnake() {
		b.snake.Move(b.moveDirection)

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
	return b.timer%b.moveTime == 0
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

func (b *board) SetSpeed(speed int) {
	b.moveTime = speed
}

func (b *board) GetSnake() *characters.Snake {
	return b.snake
}

func (b *board) GetApple() *characters.Apple {
	return b.apple
}
