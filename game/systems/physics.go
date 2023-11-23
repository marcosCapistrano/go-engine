package systems

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/forces"
	"engine/util"
	"math"
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

func (system *Movement) Process(registry ecs.Registry) {

	mask := components.MaskLinearMotion & components.MaskAngularAcceleration & components.MaskRotation
	for _, e := range registry.FilterByMask(mask) {
		mass := e.Get(components.MaskMass).(*components.Mass)
		momentOfInertia := e.Get(components.MaskMomentOfInertia).(*components.MomentOfInertia)
		position := e.Get(components.MaskPosition).(*components.Position)
		linearMotion := e.Get(components.MaskLinearMotion).(*components.LinearMotion)
		rotation := e.Get(components.MaskRotation).(*components.Rotation)
		angularMotion := e.Get(components.MaskAngularMotion).(*components.AngularMotion)
		shape := e.Get(components.MaskShape).(*components.Shape)

		linearMotion.AddForce(forces.Weight(mass.Value))
		linearMotion.AddForce(forces.Drag(linearMotion.Velocity, util.DRAG_COEFF))
		linearMotion.Integrate(&position.Vector, mass.Inverse, *system.deltaTime)

		angularMotion.AddTorque(100)
		angularMotion.Integrate(&rotation.Value, momentOfInertia.Inverse, *system.deltaTime)

		if rotation.Value > 360 {
			rotation.Value = math.Mod(rotation.Value, 360)
		}

		shape.UpdateVertices(position.Vector, rotation.Value)
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
