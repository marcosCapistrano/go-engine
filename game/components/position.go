package components

import "engine/math/vector"

type Position struct {
	ID     string
	Vector vector.Vector2
}

func (p *Position) Mask() uint64 {
	return MaskPosition
}
