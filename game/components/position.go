package components

import (
	"engine/ecs"
	"engine/math/vector"
)

// Position ...
type Position struct {
	ID             string `json:"id"`
	vector.Vector2 `json:"position"`
}

func (a *Position) Mask() uint64 {
	return MaskPosition
}

func (a *Position) WithX(x float64) *Position {
	a.X = x
	return a
}

func (a *Position) WithY(y float64) *Position {
	a.Y = y
	return a
}

// NewPosition ...
func NewPosition() ecs.Component {
	return &Position{
		ID: "position",
	}
}
