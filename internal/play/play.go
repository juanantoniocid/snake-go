package play

import (
	"juanantoniocid/snake/internal/geometry"
	"math/rand"

	"juanantoniocid/snake/internal/play/characters"
)

type Status int

const (
	StatusInitial Status = iota
	StatusPlaying
	StatusGameOver
)

type Play struct {
	boardWidth  int
	boardHeight int

	snake *characters.Snake
	Apple *characters.Apple

	status        Status
	timer         int
	moveDirection geometry.Direction
	moveTime      int
	score         int
	level         int
}

func NewPlay(width, height int) *Play {
	g := &Play{
		boardWidth:  width,
		boardHeight: height,
	}
	g.Reset()

	return g
}

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
	p.Apple = characters.NewApple(rand.Intn(p.boardWidth-1), rand.Intn(p.boardHeight-1))
}

func (p *Play) initSnake() {
	p.snake = characters.NewSnake(p.boardWidth/2, p.boardHeight/2)
}

func (p *Play) GetStatus() Status {
	return p.status
}

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

func (p *Play) snakeCollidesWithApple() bool {
	applePos := p.Apple.GetPosition()
	snakePos := p.snake.GetHead()
	return snakePos.X == applePos.X &&
		snakePos.Y == applePos.Y
}

func (p *Play) snakeCollidesWithSelf() bool {
	head := p.snake.GetHead()
	tail := p.snake.GetTail()
	for _, v := range tail {
		if head.X == v.X &&
			head.Y == v.Y {
			return true
		}
	}
	return false
}

func (p *Play) snakeCollidesWithWall() bool {
	head := p.snake.GetHead()
	return head.X < 0 ||
		head.Y < 0 ||
		head.X >= p.boardWidth ||
		head.Y >= p.boardHeight
}

func (p *Play) needsToMoveSnake() bool {
	return p.timer%p.moveTime == 0
}

func (p *Play) GetScore() int {
	return p.score
}

func (p *Play) GetLevel() int {
	return p.level
}

func (p *Play) GetSnakeShape() geometry.Shape {
	return p.snake.GetShape()
}
