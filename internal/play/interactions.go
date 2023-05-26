package play

import "juanantoniocid/snake/internal/geometry"

func (p *Play) setSnakeDirection(dir geometry.Direction) {
	if dir == geometry.DirLeft {
		if p.moveDirection != geometry.DirRight {
			p.moveDirection = geometry.DirLeft
		}
	} else if dir == geometry.DirRight {
		if p.moveDirection != geometry.DirLeft {
			p.moveDirection = geometry.DirRight
		}
	} else if dir == geometry.DirDown {
		if p.moveDirection != geometry.DirUp {
			p.moveDirection = geometry.DirDown
		}
	} else if dir == geometry.DirUp {
		if p.moveDirection != geometry.DirDown {
			p.moveDirection = geometry.DirUp
		}
	}
}

func (p *Play) moveSnake() {
	if p.needsToMoveSnake() {
		if p.snakeCollidesWithWall() || p.snakeCollidesWithSelf() {
			p.Reset()
		}

		if p.snakeEatsApple() {
			p.initApple()
			p.snake.Grow()
			p.setLevel()
			p.score++
		}

		p.snake.Move(p.moveDirection)
	}
}

func (p *Play) needsToMoveSnake() bool {
	return p.timer%p.moveTime == 0
}

func (p *Play) snakeCollidesWithWall() bool {
	snakeHead := p.snake.GetShape()[0]
	return snakeHead.X < 0 || snakeHead.Y < 0 ||
		snakeHead.X >= p.boardWidth || snakeHead.Y >= p.boardHeight
}

func (p *Play) snakeCollidesWithSelf() bool {
	snakeShape := p.snake.GetShape()
	snakeHead := snakeShape[0]
	snakeTail := snakeShape[1:]

	for _, v := range snakeTail {
		if snakeHead.X == v.X && snakeHead.Y == v.Y {
			return true
		}
	}
	return false
}

func (p *Play) snakeEatsApple() bool {
	appleShape := p.apple.GetShape()[0]
	snakeHead := p.snake.GetShape()[0]
	return snakeHead.X == appleShape.X && snakeHead.Y == appleShape.Y
}

func (p *Play) setLevel() {
	if p.snake.Len() > 10 && p.snake.Len() < 20 {
		p.level = 2
		p.moveTime = 3
	} else if p.snake.Len() > 20 {
		p.level = 3
		p.moveTime = 2
	} else {
		p.level = 1
	}
}
