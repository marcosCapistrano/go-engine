package components

import (
	"engine/ecs"
	"engine/math/vector"
)

// Velocity ...
type Acceleration struct {
	ID            string `json:"id"`
	vector.Vector `json:"acceleration"`
}

func (a *Acceleration) Mask() uint64 {
	return MaskAcceleration
}

func (a *Acceleration) WithX(x float32) *Acceleration {
	a.X = x
	return a
}

func (a *Acceleration) WithY(y float32) *Acceleration {
	a.Y = y
	return a
}

// NewVelocity ...
func NewAcceleration() ecs.Component {
	return &Acceleration{
		ID: "acceleration",
	}
}
