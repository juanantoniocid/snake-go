package apple

import (
	"juanantoniocid/snake/internal/position"
)

// Apple represents an apple
type Apple struct {
	position position.Position
}

// GetPosition returns the apple position
func (a *Apple) GetPosition() position.Position {
	return a.position
}

// NewApple creates a new apple
func NewApple(posX, posY int) *Apple {
	return &Apple{
		position: position.Position{X: posX, Y: posY},
	}
}
