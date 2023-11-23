package entities

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
)

func NewCircle(x, y, rotation, mass, radius float64) []ecs.Component {
	return []ecs.Component{
		&components.Position{
			Vector: vector.Vector2{X: x, Y: y},
		},
		&components.Rotation{
			Value: rotation,
		},
		&components.Mass{
			Value:   mass,
			Inverse: 1 / mass,
		},
		&components.MomentOfInertia{
			Value:   0.5 * (radius * radius),
			Inverse: 1 / (0.5 * (radius * radius)),
		},
		&components.LinearMotion{
			Velocity:     vector.NewVector2(0, 0),
			Acceleration: vector.NewVector2(0, 0),
			Forces:       vector.NewVector2(0, 0),
		},
		&components.AngularMotion{
			Velocity:     10,
			Acceleration: 0,
			Torque:       0,
		},
		&components.Shape{
			Type:   components.CircleType,
			Radius: radius,
		},
	}
}
