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

	direction geometry.Direction
	speed     int
	timer     int
}

func newBoard(width, height int) *board {
	b := &board{
		width:     width,
		height:    height,
		direction: geometry.DirNone,
		speed:     4,
		timer:     0,
	}

	b.initSnake()
	b.initApple()

	return b
}

func (b *board) MoveSnake(dir geometry.Direction) Status {
	b.timer++
	return b.moveSnake(dir)
}

func (b *board) moveSnake(dir geometry.Direction) Status {
	if dir == geometry.DirLeft {
		if b.direction != geometry.DirRight {
			b.direction = geometry.DirLeft
		}
	} else if dir == geometry.DirRight {
		if b.direction != geometry.DirLeft {
			b.direction = geometry.DirRight
		}
	} else if dir == geometry.DirDown {
		if b.direction != geometry.DirUp {
			b.direction = geometry.DirDown
		}
	} else if dir == geometry.DirUp {
		if b.direction != geometry.DirDown {
			b.direction = geometry.DirUp
		}
	}

	return b.advanceSnake()
}

func (b *board) advanceSnake() Status {
	if b.needsToMoveSnake() {
		b.snake.Move(b.direction)

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
	return b.timer%b.speed == 0
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
	b.speed = speed
}

func (b *board) GetSnake() *characters.Snake {
	return b.snake
}

func (b *board) GetApple() *characters.Apple {
	return b.apple
}
