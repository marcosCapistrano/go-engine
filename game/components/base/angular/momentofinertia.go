package angular

import "engine/ecs"

type MomentOfInertia struct {
	ID              string  `json:"id"`
	MomentOfInertia float64 `json:"mass"`
}

func (m *MomentOfInertia) Mask() uint64 {
	return MaskMomentOfInertia
}

func (m *MomentOfInertia) WithMomentOfInertia(value float64) *MomentOfInertia {
	m.MomentOfInertia = value

	return m
}

// NewPosition ...
func NewMomentOfInertia() ecs.Component {
	return &MomentOfInertia{
		ID: "moment-of-inertia",
	}
}
