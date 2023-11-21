package components

import (
	"engine/ecs"
	"engine/math/vector"
)

// Velocity ...
type Velocity struct {
	ID            string `json:"id"`
	vector.Vector `json:"velocity"`
}

func (a *Velocity) Mask() uint64 {
	return MaskVelocity
}

func (a *Velocity) WithX(x float32) *Velocity {
	a.X = x
	return a
}

func (a *Velocity) WithY(y float32) *Velocity {
	a.Y = y
	return a
}

// NewVelocity ...
func NewVelocity() ecs.Component {
	return &Velocity{
		ID: "velocity",
	}
}
