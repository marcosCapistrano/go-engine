package systems

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
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
	gravity := vector.Vector{X: 0, Y: 2}
	for _, e := range registry.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		acceleration := e.Get(components.MaskAcceleration).(*components.Acceleration)

		acceleration.Add(gravity)
		velocity.Add(acceleration.Vector)
		position.Add(velocity.Mult(*a.deltaTime))

		acceleration.Vector = acceleration.Mult(0)
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
