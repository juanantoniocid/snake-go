package geometry

// Direction is a direction in a 2D space.
type Direction int

const (
	DirNone Direction = iota
	DirLeft
	DirRight
	DirDown
	DirUp
)
