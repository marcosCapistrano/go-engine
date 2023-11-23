package linear

import (
	"engine/math/vector"
)

type LinearMotion struct {
	Position     vector.Vector2
	Velocity     vector.Vector2
	Acceleration vector.Vector2
}

/*
func (a *Lin) Mask() uint64 {
	return MaskBoxCollider
}

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
