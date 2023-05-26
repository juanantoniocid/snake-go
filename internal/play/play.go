package play

import (
	"math/rand"

	"juanantoniocid/snake/internal/direction"
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

	Snake *characters.Snake
	Apple *characters.Apple

	status        Status
	timer         int
	moveDirection int
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
	p.moveDirection = direction.DirNone
	p.moveTime = 4
	p.score = 0
	p.level = 1
}

func (p *Play) initApple() {
	p.Apple = characters.NewApple(rand.Intn(p.boardWidth-1), rand.Intn(p.boardHeight-1))
}

func (p *Play) initSnake() {
	p.Snake = characters.NewSnake(p.boardWidth/2, p.boardHeight/2)
}

func (p *Play) GetStatus() Status {
	return p.status
}

func (p *Play) MoveSnake(dir int) error {
	p.setSnakeDirection(dir)
	p.moveSnake()
	p.timer++

	return nil
}

func (p *Play) setSnakeDirection(dir int) {
	if dir == direction.DirLeft {
		if p.moveDirection != direction.DirRight {
			p.moveDirection = direction.DirLeft
		}
	} else if dir == direction.DirRight {
		if p.moveDirection != direction.DirLeft {
			p.moveDirection = direction.DirRight
		}
	} else if dir == direction.DirDown {
		if p.moveDirection != direction.DirUp {
			p.moveDirection = direction.DirDown
		}
	} else if dir == direction.DirUp {
		if p.moveDirection != direction.DirDown {
			p.moveDirection = direction.DirUp
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
			p.Snake.Grow()
			if p.Snake.Len() > 10 && p.Snake.Len() < 20 {
				p.level = 2
				p.moveTime = 3
			} else if p.Snake.Len() > 20 {
				p.level = 3
				p.moveTime = 2
			} else {
				p.level = 1
			}
			p.score++
		}

		p.Snake.Move(p.moveDirection)
	}
}

func (p *Play) snakeCollidesWithApple() bool {
	applePos := p.Apple.GetPosition()
	snakePos := p.Snake.GetHead()
	return snakePos.X == applePos.X &&
		snakePos.Y == applePos.Y
}

func (p *Play) snakeCollidesWithSelf() bool {
	head := p.Snake.GetHead()
	tail := p.Snake.GetTail()
	for _, v := range tail {
		if head.X == v.X &&
			head.Y == v.Y {
			return true
		}
	}
	return false
}

func (p *Play) snakeCollidesWithWall() bool {
	head := p.Snake.GetHead()
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
