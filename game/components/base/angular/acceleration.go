package angular

import (
	"engine/ecs"
)

// Position ...
type AngularAcceleration struct {
	ID                  string `json:"id"`
	AngularAcceleration float64
}

func (a *AngularAcceleration) Mask() uint64 {
	return MaskAngularAcceleration
}

func (a *AngularAcceleration) WithAngularAcceleration(r float64) *AngularAcceleration {
	a.AngularAcceleration = r
	return a
}

// NewRotation ...
func NewAngularAcceleration() ecs.Component {
	return &AngularAcceleration{
		ID: "angular-acceleration",
	}
}
