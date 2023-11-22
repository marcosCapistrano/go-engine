package systems

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
	"engine/util"
	"fmt"
)

// Movement ...
type Movement struct {
	err       error
	deltaTime *float32
}

func (a *Movement) Error() (err error) {
	return a.err
}

func (a *Movement) Setup() {}

func (a *Movement) Process(registry ecs.Registry) {
	gravity := vector.Vector2{X: 0, Y: 9.8 * util.PIXELS_PER_METER}

	for _, e := range registry.FilterByMask(components.MaskPosition | components.MaskVelocity | components.MaskMass) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		acceleration := e.Get(components.MaskAcceleration).(*components.Acceleration)
		mass := e.Get(components.MaskMass).(*components.Mass)

		fmt.Println("------")
		forces := vector.Vector2{}
		fmt.Println("forces", forces)

		weight := gravity
		weight.Scale(mass.Mass)
		fmt.Println("weight", weight)

		forces.Add(weight)
		forces.Scale(1 / mass.Mass)
		fmt.Println("canceled", forces)

		acceleration.Add(forces)

		acceleration.Scale(float64(*a.deltaTime))
		velocity.Add(acceleration.Vector2)
		velocity.Scale(float64(*a.deltaTime))
		position.Add(velocity.Vector2)

		forces.Scale(0)
	}
}

func (a *Movement) Teardown() {

}

func (m *Movement) WithData(deltaTime *float32) *Movement {
	m.deltaTime = deltaTime
	return m
}

// NewMovement ...
func NewMovement() ecs.System {
	return &Movement{}
}
