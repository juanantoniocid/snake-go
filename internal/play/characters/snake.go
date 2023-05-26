package characters

import (
	"juanantoniocid/snake/internal/geometry"
)

// Snake represents a snake
type Snake struct {
	body geometry.Shape
}

// Move moves the snake
func (s *Snake) Move(dir geometry.Direction) {
	for i := int64(len(s.body)) - 1; i > 0; i-- {
		s.body[i].X = s.body[i-1].X
		s.body[i].Y = s.body[i-1].Y
	}

	head := &s.body[0]
	switch dir {
	case geometry.DirLeft:
		head.X--
	case geometry.DirRight:
		head.X++
	case geometry.DirDown:
		head.Y++
	case geometry.DirUp:
		head.Y--
	}
}

// Grow grows the snake
func (s *Snake) Grow() {
	tail := s.body[len(s.body)-1]
	s.body = append(s.body, tail)
}

// GetHead returns the snake head
func (s *Snake) GetHead() geometry.Position {
	return s.body[0]
}

// GetTail returns the snake tail
func (s *Snake) GetTail() []geometry.Position {
	return s.body[1:]
}

func (s *Snake) GetShape() geometry.Shape {
	return s.body
}

func (s *Snake) Len() int {
	return len(s.body)
}

// NewSnake creates a new snake
func NewSnake(posX, posY int) *Snake {
	return &Snake{
		body: []geometry.Position{
			{X: posX, Y: posY},
		},
	}
}
