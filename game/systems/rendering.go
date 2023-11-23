package systems

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
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

	mask := components.MaskShape & components.MaskPosition & components.MaskRotation
	for _, e := range registry.FilterByMask(mask) {
		shape := e.Get(components.MaskShape).(*components.Shape)
		position := e.Get(components.MaskPosition).(*components.Position)
		rotation := e.Get(components.MaskRotation).(*components.Rotation)

		a.renderer.SetDrawColor(255, 0, 0, 255)

		switch shape.Type {
		case components.CircleType:
			drawCircle(a.renderer, position.Vector, shape.Radius, rotation.Value)
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

func drawCircle(renderer *sdl.Renderer, position vector.Vector2, radius, rotationAngle float64) {
	var angle float64 = 0

	for i := 0; i < 360; i++ {
		angle = float64(i) * (math.Pi / 180) // converting degrees to radians
		dx := radius * math.Cos(angle)
		dy := radius * math.Sin(angle)

		renderer.DrawPoint(int32(position.X+dx), int32(position.Y+dy))
	}

	// Draw a line in the middle for debugging rotation
	rotationAngle = rotationAngle * (math.Pi / 180) // converting degrees to radians
	dx := radius * math.Cos(rotationAngle)
	dy := radius * math.Sin(rotationAngle)

	renderer.DrawLine(int32(position.X), int32(position.Y), int32(position.X+dx), int32(position.Y+dy))
}
