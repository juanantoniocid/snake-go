package characters

import (
	"juanantoniocid/snake/internal/geometry"
)

// Apple represents an apple
type Apple struct {
	position geometry.Position
}

// GetPosition returns the apple position
func (a *Apple) GetPosition() geometry.Position {
	return a.position
}

// NewApple creates a new apple
func NewApple(posX, posY int) *Apple {
	return &Apple{
		position: geometry.Position{X: posX, Y: posY},
	}
}
