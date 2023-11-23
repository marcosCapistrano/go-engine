package systems

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
	"engine/util"
)

// Movement ...
type Movement struct {
	err       error
	deltaTime *float64
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

		// Vector2.Scale() return Vector2
		weight := gravity.Scale(mass.Mass)
		drag := vector.GenerateDragForce(velocity.Vector2, 0.2)

		forces := vector.Vector2{}

		forces = forces.Add(weight).Scale(1 / mass.Mass)
		forces = forces.Add(drag)
		acceleration.Vector2 = forces.Scale(*a.deltaTime)
		velocity.Vector2 = velocity.Add(acceleration.Vector2.Scale(*a.deltaTime))
		position.Vector2 = position.Add(velocity.Vector2.Scale(*a.deltaTime))
	}
}

func (a *Movement) Teardown() {

}

func (m *Movement) WithData(deltaTime *float64) *Movement {
	m.deltaTime = deltaTime
	return m
}

// NewMovement ...
func NewMovement() ecs.System {
	return &Movement{}
}
