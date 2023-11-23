package components

import (
	"engine/math/vector"
)

type LinearMotion struct {
	ID           string `json:"id"`
	Velocity     vector.Vector2
	Acceleration vector.Vector2
	Forces       vector.Vector2
}

func (lm *LinearMotion) Mask() uint64 {
	return MaskLinearMotion
}

// Integrate causes the integration between acceleration, velocity, and position according to deltaTime(dt) and mass
func (lm *LinearMotion) Integrate(position *vector.Vector2, invMass float64, dt float64) {
	sumOfForces := lm.Forces.Scale(invMass)
	lm.Acceleration = sumOfForces
	lm.Velocity = lm.Velocity.Add(lm.Acceleration.Scale(dt))
	*position = (*position).Add(lm.Velocity.Scale(dt))

	// clear forces
	lm.Forces = vector.Vector2{}
}

func (lm *LinearMotion) AddForce(vec vector.Vector2) *LinearMotion {
	lm.Forces = lm.Forces.Add(vec)

	return lm
}

/*
func (a *BoxCollider) WithWidth(w int) *BoxCollider {
	a.Width = w
	return a
}

func (a *BoxCollider) WithHeight(h int) *BoxCollider {
	a.Height = h
	return a
}

// NewPosition ...
func NewBoxCollider() ecs.Component {
	return &BoxCollider{
		ID: "box-collider",
	}
}
*/
