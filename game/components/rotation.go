package components

import (
	"engine/ecs"
)

// Position ...
type Rotation struct {
	ID    string `json:"id"`
	angle float64
}

func (a *Rotation) Mask() uint64 {
	return MaskRotation
}

func (a *Rotation) WithAngle(r float64) *Rotation {
	a.angle = r
	return a
}

// NewRotation ...
func NewRotation() ecs.Component {
	return &Rotation{
		ID: "rotation",
	}
}
