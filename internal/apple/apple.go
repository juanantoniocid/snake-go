package apple

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"juanantoniocid/snake/internal/position"
)

// Apple represents an apple
type Apple struct {
	position position.Position
	size     int
}

// Draw draws the apple
func (a *Apple) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(a.position.X*a.size),
		float32(a.position.Y*a.size),
		float32(a.size),
		float32(a.size),
		color.RGBA{R: 0xFF, A: 0xff},
		false,
	)
}

// GetPosition returns the apple position
func (a *Apple) GetPosition() position.Position {
	return a.position
}

// NewApple creates a new apple
func NewApple(posX, posY, size int) *Apple {
	return &Apple{
		position: position.Position{X: posX, Y: posY},
		size:     size,
	}
}
