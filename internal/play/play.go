package play

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"juanantoniocid/snake/internal/direction"
	"juanantoniocid/snake/internal/play/characters"
)

type Play struct {
	boardWidth  int
	boardHeight int

	MoveDirection int
	Snake         *characters.Snake
	Apple         *characters.Apple

	timer    int
	moveTime int

	Score     int
	BestScore int
	Level     int
}

func (p *Play) collidesWithApple() bool {
	applePos := p.Apple.GetPosition()
	snakePos := p.Snake.GetHead()
	return snakePos.X == applePos.X &&
		snakePos.Y == applePos.Y
}

func (p *Play) collidesWithSelf() bool {
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

func (p *Play) collidesWithWall() bool {
	head := p.Snake.GetHead()
	return head.X < 0 ||
		head.Y < 0 ||
		head.X >= p.boardWidth ||
		head.Y >= p.boardHeight
}

func (p *Play) needsToMoveSnake() bool {
	return p.timer%p.moveTime == 0
}

func (p *Play) reset() {
	p.initApple()
	p.initSnake()

	p.moveTime = 4
	p.Score = 0
	p.Level = 1
	p.MoveDirection = direction.DirNone
}

func (p *Play) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if p.MoveDirection != direction.DirRight {
			p.MoveDirection = direction.DirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if p.MoveDirection != direction.DirLeft {
			p.MoveDirection = direction.DirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if p.MoveDirection != direction.DirUp {
			p.MoveDirection = direction.DirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if p.MoveDirection != direction.DirDown {
			p.MoveDirection = direction.DirUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.reset()
	}

	if p.needsToMoveSnake() {
		if p.collidesWithWall() || p.collidesWithSelf() {
			p.reset()
		}

		if p.collidesWithApple() {
			p.initApple()
			p.Snake.Grow()
			if p.Snake.Len() > 10 && p.Snake.Len() < 20 {
				p.Level = 2
				p.moveTime = 3
			} else if p.Snake.Len() > 20 {
				p.Level = 3
				p.moveTime = 2
			} else {
				p.Level = 1
			}
			p.Score++
			if p.BestScore < p.Score {
				p.BestScore = p.Score
			}
		}

		p.Snake.Move(p.MoveDirection)
	}

	p.timer++

	return nil
}

func (p *Play) initApple() {
	p.Apple = characters.NewApple(rand.Intn(p.boardWidth-1), rand.Intn(p.boardHeight-1))
}

func (p *Play) initSnake() {
	p.Snake = characters.NewSnake(p.boardWidth/2, p.boardHeight/2)
}

func NewPlay(width, height int) *Play {
	g := &Play{
		boardWidth:  width,
		boardHeight: height,
		moveTime:    4,
	}
	g.initApple()
	g.initSnake()
	return g
}
