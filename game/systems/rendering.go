package systems

import (
	"engine/ecs"
	"engine/game/components"
	"math"

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

	for _, e := range registry.FilterByMask(components.MaskPosition | components.MaskShape | components.MaskRotation) {
		position := e.Get(components.MaskPosition).(*components.Position)
		shape := e.Get(components.MaskShape).(*components.Shape)
		rotation := e.Get(components.MaskRotation).(*components.Rotation)

		a.renderer.SetDrawColor(255, 0, 0, 255)

		switch shape.Type {
		case "circle":
			drawCircle(a.renderer, position.X, position.Y, shape.Radius, rotation.Angle)
		}
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

func drawCircle(renderer *sdl.Renderer, x, y, radius, rotationAngle float64) {
	var angle float64 = 0

	for i := 0; i < 360; i++ {
		angle = float64(i) * (math.Pi / 180) // converting degrees to radians
		dx := radius * math.Cos(angle)
		dy := radius * math.Sin(angle)

		renderer.DrawPoint(int32(x+dx), int32(y+dy))
	}

	// Draw a line in the middle for debugging rotation
	rotationAngle = rotationAngle * (math.Pi / 180) // converting degrees to radians
	dx := radius * math.Cos(rotationAngle)
	dy := radius * math.Sin(rotationAngle)

	renderer.DrawLine(int32(x), int32(y), int32(x+dx), int32(y+dy))
}
