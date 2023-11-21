package game

import "github.com/veandco/go-sdl2/sdl"

func (g *Game) processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			g.isRunning = false

		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				g.isRunning = false
			}
		}
	}
}
