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
	board *board

	status Status
	score  int
	level  int
}

// NewPlay creates a new game
func NewPlay(width, height int) *Play {
	g := &Play{
		board: newBoard(width, height),

		status: StatusInitial,
		score:  0,
		level:  1,
	}

	return g
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
	return p.board.GetSnakeShape()
}

// GetAppleShape returns the current apple shape
func (p *Play) GetAppleShape() geometry.Shape {
	return p.board.GetAppleShape()
}

// MoveSnake moves the snake in the given direction
func (p *Play) MoveSnake(dir geometry.Direction) {
	if p.status == StatusGameOver {
		return
	}

	p.status = p.board.MoveSnake(dir)
	if p.status == StatusSnakeEating {
		p.increaseScore()
		p.status = StatusPlaying
		return
	}
}

func (p *Play) increaseScore() {
	var speed int
	p.score++

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
