package components

import (
	"engine/ecs"
	"engine/math/vector"
)

// ShapeType represents the type of shape.
type ShapeType string

const (
	CircleType   ShapeType = "circle"
	BoxType      ShapeType = "box"
	TriangleType ShapeType = "triangle"
	PolygonType  ShapeType = "polygon"
)

// Shape ...
type Shape struct {
	ID       string           `json:"id"`
	Type     ShapeType        `json:"type"`
	Radius   float64          `json:"radius,omitempty"`
	Width    float64          `json:"width,omitempty"`
	Height   float64          `json:"height,omitempty"`
	Vertices []vector.Vector2 `json:"sides,omitempty"`

	LocalVertices []vector.Vector2
	WorldVertices []vector.Vector2
}

func (a *Shape) Mask() uint64 {
	return MaskShape
}

func (a *Shape) UpdateVertices(position vector.Vector2, rotation float64) {
	if a.Type != CircleType {
		for i := 0; i < len(a.LocalVertices); i++ {
			//a.WorldVertices[i] = a.LocalVertices[i].Rotate(rotation)
		}
	}

	for i := 0; i < len(a.LocalVertices); i++ {
		//a.WorldVertices[i] = a.LocalVertices[i].Add(position)
	}
}

// NewRotation ...
func NewShape() ecs.Component {
	return &Shape{
		ID: "shape",
	}
}
