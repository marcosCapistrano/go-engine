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
	deltaTime *float64
}

func (a *Movement) Error() (err error) {
	return a.err
}

func (a *Movement) Setup() {}

func (a *Movement) Process(registry ecs.Registry) {
	gravity := vector.Vector2{X: 0, Y: 9.8 * util.PIXELS_PER_METER}

	mask := components.MaskPosition | components.MaskVelocity | components.MaskAcceleration | components.MaskMass
	for _, e := range registry.FilterByMask(mask) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		acceleration := e.Get(components.MaskAcceleration).(*components.Acceleration)
		mass := e.Get(components.MaskMass).(*components.Mass)
		rotation := e.Get(components.MaskRotation).(*components.Rotation)
		angularVel := e.Get(components.MaskAngularVelocity).(*components.AngularVelocity)
		angularAcc := e.Get(components.MaskAngularAcceleration).(*components.AngularAcceleration)
		shape := e.Get(components.MaskShape).(*components.Shape)

		// Vector2.Scale() return Vector2
		weight := gravity.Scale(mass.Mass)
		drag := vector.GenerateDragForce(velocity.Vector2, 0.2)

		forces := vector.Vector2{}

		forces = forces.Add(weight)
		forces = forces.Add(drag)

		torque := 20000.0

		//moment of inertia = 0.5 * (radius * radius)
		I := 0.5 * (shape.Radius * shape.Radius) * mass.Mass
		var invI float64

		if I != 0 {
			invI = 1.0 / I
		} else {
			invI = 0
		}

		angularAcc.AngularAcceleration = torque * invI
		fmt.Println(angularAcc.AngularAcceleration)
		angularVel.AngularVelocity += angularAcc.AngularAcceleration * (*a.deltaTime)
		fmt.Println(angularVel.AngularVelocity)
		rotation.Angle += angularVel.AngularVelocity * (*a.deltaTime)
		acceleration.Vector2 = forces.Scale(1.0 / mass.Mass)
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
