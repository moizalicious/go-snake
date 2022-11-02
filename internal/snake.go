package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions

	lastKeyPress ebiten.Key

	posX, posY int
}

const (
	speed = 10

	snakeWidth  = 100
	snakeHeight = 10
)

func NewSnake() *Snake {
	i := ebiten.NewImage(snakeWidth, snakeHeight)
	i.Fill(color.White)

	o := &ebiten.DrawImageOptions{}

	return &Snake{
		image:   i,
		options: o,
		// lastKeyPress: ebiten.KeyRight,
	}
}

func (s *Snake) UpdatePosition() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		s.lastKeyPress = ebiten.KeyUp
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		s.lastKeyPress = ebiten.KeyDown
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		s.lastKeyPress = ebiten.KeyLeft
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		s.lastKeyPress = ebiten.KeyRight
	}

	switch s.lastKeyPress {
	case ebiten.KeyUp:
		s.moveUp()
	case ebiten.KeyDown:
		s.moveDown()
	case ebiten.KeyLeft:
		s.moveLeft()
	case ebiten.KeyRight:
		s.moveRight()
	}
}

func (s *Snake) moveUp() {
	if s.posY > 0 {
		s.options.GeoM.Translate(0, -speed)
		s.posY += -speed
	}
}

func (s *Snake) moveDown() {
	if s.posY < GameHeight-snakeHeight {
		s.options.GeoM.Translate(0, speed)
		s.posY += speed
	}
}

func (s *Snake) moveLeft() {
	if s.posX > 0 {
		s.options.GeoM.Translate(-speed, 0)
		s.posX += -speed
	}
}

func (s *Snake) moveRight() {
	if s.posX < GameWidth-snakeWidth {
		s.options.GeoM.Translate(speed, 0)
		s.posX += speed
	}
}
