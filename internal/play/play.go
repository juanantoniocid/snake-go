package play

import (
	"juanantoniocid/snake/internal/geometry"
)

type Status int

const (
	StatusInitial Status = iota
	StatusPlaying
	StatusSnakeEating
	StatusGameOver
)

// Play is the main game struct
type Play struct {
	board       *board
	boardWidth  int
	boardHeight int

	status Status

	score int
	level int
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
	p.board = newBoard(p.boardWidth, p.boardHeight)
	p.setLevel()

	p.status = StatusInitial
	p.score = 0
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
	return p.board.GetSnake().GetShape()
}

// GetAppleShape returns the current apple shape
func (p *Play) GetAppleShape() geometry.Shape {
	return p.board.GetApple().GetShape()
}

// MoveSnake moves the snake in the given direction
func (p *Play) MoveSnake(dir geometry.Direction) {
	if p.status == StatusGameOver {
		return
	}

	p.status = p.board.MoveSnake(dir)
	p.setLevel()

	if p.status == StatusSnakeEating {
		p.score++
		p.status = StatusPlaying
		return
	}

	if p.status == StatusInitial && dir != geometry.DirNone {
		p.status = StatusPlaying
	}
}

func (p *Play) setLevel() {
	var speed int
	if p.score < 10 {
		p.level = 1
		speed = 4
	} else if p.score < 20 {
		p.level = 2
		speed = 3
	} else if p.score < 30 {
		p.level = 3
		speed = 2
	} else {
		p.level = 4
		speed = 1
	}

	p.board.SetSpeed(speed)
}
