package bridge

import "fmt"

type Shape interface {
	CreateColorShape()
}

type Color interface {
	AddColor(shape string)
}

type RedColor struct{}

func NewRedColor() Color {
	return &RedColor{}
}

func (*RedColor) AddColor(shape string) {
	fmt.Println("Red + " + shape)
}

type BlueColor struct{}

func NewBlueColor() Color {
	return &BlueColor{}
}

func (*BlueColor) AddColor(shape string) {
	fmt.Println("Blue + " + shape)
}

type CircleShape struct {
	color_ Color
}

func NewCircleShape(color_ Color) *CircleShape {
	return &CircleShape{
		color_: color_,
	}
}

func (m *CircleShape) CreateColorShape() {
	m.color_.AddColor("CircleShape")
}

type RectangleShape struct {
	color_ Color
}

func NewRectangleShape(color_ Color) *RectangleShape {
	return &RectangleShape{
		color_: color_,
	}
}

func (m *RectangleShape) CreateColorShape() {
	m.color_.AddColor("RectangleShape")
}

func main() {
	s1 := NewRedColor()
	s2 := NewBlueColor()
	a := NewCircleShape(s1)
	b := NewRectangleShape(s2)
	c := NewCircleShape(s2)
	d := NewRectangleShape(s1)
	a.CreateColorShape()
	b.CreateColorShape()
	c.CreateColorShape()
	d.CreateColorShape()
}

//output:
//Red + CircleShape
//Blue + RectangleShape
//Blue + CircleShape
//Red + RectangleShape
