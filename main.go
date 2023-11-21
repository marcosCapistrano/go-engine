package main

import (
	"engine/game"
)

func main() {
	game := game.New()

	game.Initialize()
	game.Run()
	game.Destroy()
}
