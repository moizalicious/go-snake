package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	snake *Snake
}

const (
	GameTitle = "Go Snake (Alpha)"

	GameWidth  = 800
	GameHeight = 600
)

func NewGame() *Game {
	ebiten.SetWindowSize(GameWidth, GameHeight)
	ebiten.SetWindowTitle(GameTitle)

	return &Game{
		snake: NewSnake(),
	}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.snake.UpdatePosition()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.snake == nil || g.snake.image == nil || g.snake.options == nil {
		panic("Snake has not been initialised!")
	}

	screen.DrawImage(g.snake.image, g.snake.options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}
