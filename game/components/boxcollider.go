package components

import (
	"engine/ecs"
)

// Position ...
type BoxCollider struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (a *BoxCollider) Mask() uint64 {
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
