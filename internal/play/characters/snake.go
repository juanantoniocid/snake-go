package characters

import (
	"juanantoniocid/snake/internal/direction"
	"juanantoniocid/snake/internal/position"
)

// Snake represents a snake
type Snake struct {
	body []position.Position
}

// Move moves the snake
func (s *Snake) Move(dir int) {
	for i := int64(len(s.body)) - 1; i > 0; i-- {
		s.body[i].X = s.body[i-1].X
		s.body[i].Y = s.body[i-1].Y
	}

	head := &s.body[0]
	switch dir {
	case direction.DirLeft:
		head.X--
	case direction.DirRight:
		head.X++
	case direction.DirDown:
		head.Y++
	case direction.DirUp:
		head.Y--
	}
}

// Grow grows the snake
func (s *Snake) Grow() {
	tail := s.body[len(s.body)-1]
	s.body = append(s.body, tail)
}

// GetHead returns the snake head
func (s *Snake) GetHead() position.Position {
	return s.body[0]
}

// GetTail returns the snake tail
func (s *Snake) GetTail() []position.Position {
	return s.body[1:]
}

func (s *Snake) GetBody() []position.Position {
	return s.body
}

func (s *Snake) Len() int {
	return len(s.body)
}

// NewSnake creates a new snake
func NewSnake(posX, posY int) *Snake {
	return &Snake{
		body: []position.Position{
			{X: posX, Y: posY},
		},
	}
}
