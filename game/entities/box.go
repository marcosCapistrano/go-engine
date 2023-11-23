package entities

import (
	"engine/ecs"
	"engine/game/components"
	"engine/math/vector"
)

func NewBox(x, y, rotation, mass, width, height float64) []ecs.Component {
	return []ecs.Component{
		&components.Position{
			Vector: vector.Vector2{X: x, Y: y},
		},
		&components.Rotation{
			ID:    "rotation",
			Value: rotation,
		},
		&components.Mass{
			Value:   mass,
			Inverse: 1 / mass,
		},
		&components.MomentOfInertia{
			Value:   (1 / 12.0) * (width*width + height*height),
			Inverse: 1 / ((1 / 12.0) * (width*width + height*height)),
		},
		&components.LinearMotion{
			Velocity:     vector.NewVector2(0, 0),
			Acceleration: vector.NewVector2(0, 0),
			Forces:       vector.NewVector2(0, 0),
		},
		&components.AngularMotion{
			Velocity:     0,
			Acceleration: 0,
			Torque:       0,
		},
		&components.Shape{
			Type:          components.BoxType,
			Width:         width,
			Height:        height,
			LocalVertices: createVerticesForBox(width, height),
			WorldVertices: createVerticesForBox(width, height),
		},
	}
}

func createVerticesForBox(width, height float64) []vector.Vector2 {
	result := []vector.Vector2{}

	result = append(result, vector.NewVector2(-width/2, -height/2))
	result = append(result, vector.NewVector2(width/2, -height/2))
	result = append(result, vector.NewVector2(width/2, height/2))
	result = append(result, vector.NewVector2(-width/2, height/2))

	return result
}
