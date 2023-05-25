package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/snake/internal/direction"
	"juanantoniocid/snake/internal/position"
)

// Snake represents a snake
type Snake struct {
	body []position.Position
	size int
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

func (s *Snake) Len() int {
	return len(s.body)
}

// Draw draws the snake
func (s *Snake) Draw(screen *ebiten.Image) {
	for _, v := range s.body {
		vector.DrawFilledRect(
			screen,
			float32(v.X*s.size),
			float32(v.Y*s.size),
			float32(s.size),
			float32(s.size),
			color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff},
			false,
		)
	}
}

// NewSnake creates a new snake
func NewSnake(posX, posY, size int) *Snake {
	return &Snake{
		body: []position.Position{
			{X: posX, Y: posY},
		},
		size: size,
	}
}
