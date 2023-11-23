package angular

import (
	"engine/ecs"
)

// Position ...
type AngularVelocity struct {
	ID              string `json:"id"`
	AngularVelocity float64
}

func (a *AngularVelocity) Mask() uint64 {
	return MaskAngularVelocity
}

func (a *AngularVelocity) WithAngularVelocity(r float64) *AngularVelocity {
	a.AngularVelocity = r
	return a
}

// NewRotation ...
func NewAngularVelocity() ecs.Component {
	return &AngularVelocity{
		ID: "angular-velocity",
	}
}
