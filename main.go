package main

import (
	"github.com/moizalicious/go-snake/internal"
)

func main() {
	game := internal.NewGame()
	if err := game.Run(); err != nil {
		panic(err)
	}
}
