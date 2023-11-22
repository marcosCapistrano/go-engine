package components

import "engine/ecs"

type Mass struct {
	ID   string  `json:"id"`
	Mass float64 `json:"mass"`
}

func (m *Mass) Mask() uint64 {
	return MaskMass
}

func (m *Mass) WithValue(value float64) *Mass {
	m.Mass = value

	return m
}

// NewPosition ...
func NewMass() ecs.Component {
	return &Mass{
		ID: "mass",
	}
}
