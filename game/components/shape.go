package components

import (
	"engine/ecs"
)

// ShapeType represents the type of shape.
type ShapeType string

const (
	CircleType    ShapeType = "circle"
	RectangleType ShapeType = "rectangle"
	TriangleType  ShapeType = "triangle"
	PolygonType   ShapeType = "polygon"
)

// Shape ...
type Shape struct {
	ID     string    `json:"id"`
	Type   ShapeType `json:"type"`
	Radius float64   `json:"radius,omitempty"`
	Width  float64   `json:"width,omitempty"`
	Height float64   `json:"height,omitempty"`
	Sides  []float64 `json:"sides,omitempty"`
}

func (a *Shape) Mask() uint64 {
	return MaskShape
}

// WithRadius sets the radius of the circle.
func (s *Shape) WithRadius(radius float64) *Shape {
	s.Radius = radius
	return s
}

// WithWidth sets the width of the rectangle.
func (s *Shape) WithWidth(width float64) *Shape {
	s.Width = width
	return s
}

// WithHeight sets the height of the rectangle.
func (s *Shape) WithHeight(height float64) *Shape {
	s.Height = height
	return s
}

// WithSides sets the sides of the polygon.
func (s *Shape) WithSides(sides []float64) *Shape {
	s.Sides = sides
	return s
}

// WithType sets the type of the shape.
func (s *Shape) WithType(shapeType ShapeType) *Shape {
	s.Type = shapeType
	return s
}

// NewRotation ...
func NewShape() ecs.Component {
	return &Shape{
		ID: "shape",
	}
}
