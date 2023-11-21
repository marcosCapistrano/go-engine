package systems

import (
	"engine/ecs"
	"engine/game/components"
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
	for _, e := range registry.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		position.Add(velocity.Mult(*a.deltaTime))
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
