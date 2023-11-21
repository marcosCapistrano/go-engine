package systems

import (
	"engine/ecs"
	"engine/game/components"

	"github.com/veandco/go-sdl2/sdl"
)

// Rendering ...
type Rendering struct {
	err      error
	width    int
	height   int
	renderer *sdl.Renderer
}

func (a *Rendering) Error() error {
	return a.err
}

func (a *Rendering) Setup() {

}

func (a *Rendering) Process(registry ecs.Registry) {
	a.renderer.SetDrawColor(0, 0, 0, 255)
	a.renderer.Clear()

	for _, e := range registry.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		a.renderer.SetDrawColor(255, 0, 0, 255)
		a.renderer.DrawRect(&sdl.Rect{
			X: int32(position.X),
			Y: int32(position.Y),
			W: int32(size.Width),
			H: int32(size.Height),
		})
	}

	a.renderer.Present()
}

func (a *Rendering) Teardown() {

}

func (a *Rendering) WithHeight(height int) *Rendering {
	a.height = height
	return a
}

func (a *Rendering) WithWidth(width int) *Rendering {
	a.width = width
	return a
}

// NewRendering ...
func NewRendering(renderer *sdl.Renderer) ecs.System {
	return &Rendering{
		renderer: renderer,
	}
}
