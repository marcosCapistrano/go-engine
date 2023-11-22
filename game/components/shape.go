package components

import (
	"engine/ecs"
)

// Shape ...
type Shape struct {
	ID string `json:"id"`
}

func (a *Shape) Mask() uint64 {
	return MaskShape
}

func (a *Rotation) WithAngle(r float64) *Rotation {
	a.angle = r
	return a
}

// NewRotation ...
func NewShape() ecs.Component {
	return &Shape{
		ID: "shape",
	}
}
