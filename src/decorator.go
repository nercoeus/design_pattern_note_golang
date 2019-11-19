package decorator

import (
	"fmt"
)

type Component interface {
	Output() string
}

type ConcreteComponent struct {
	base string
}

func (c *ConcreteComponent) Output() string {
	return c.base
}

type BaseDecoratorA struct {
	Component
	S string
}

func AddDecoratorA(c Component, s string) Component {
	return &BaseDecoratorA{
		Component: c,
		S:         s,
	}
}

func (d *BaseDecoratorA) Output() string {
	return d.Component.Output() + "\n" + d.S
}

type BaseDecoratorB struct {
	Component
	S string
}

func AddDecoratorB(c Component, s string) Component {
	return &BaseDecoratorB{
		Component: c,
		S:         s,
	}
}

func (d *BaseDecoratorB) Output() string {
	return d.Component.Output() + "\n" + d.S
}

func main() {
	var c Component = &ConcreteComponent{
		base: "base part",
	}
	c = AddDecoratorA(c, "Add Decorator A")
	c = AddDecoratorB(c, "Add Decorator B")
	s := c.Output()
	fmt.Println(s)
}

//output:
//base part
//Add Decorator A
//Add Decorator B
