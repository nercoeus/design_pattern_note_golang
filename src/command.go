package command

import "fmt"

type Command interface {
	Execute()
}

type CommandA struct {
	mb *Parents
}

func NewCommandA(mb *Parents) *CommandA {
	return &CommandA{
		mb: mb,
	}
}

func (c *CommandA) Execute() {
	c.mb.Acommand()
}

type CommandB struct {
	mb *Parents
}

func NewCommandB(mb *Parents) *CommandB {
	return &CommandB{
		mb: mb,
	}
}

func (c *CommandB) Execute() {
	c.mb.Bcommand()
}

type Parents struct{}

func (*Parents) Acommand() {
	fmt.Print("push A button execute command A\n")
}

func (*Parents) Bcommand() {
	fmt.Print("push B button execute command B\n")
}

type Box struct {
	buttion1 Command
	buttion2 Command
}

func NewBox(buttion1, buttion2 Command) *Box {
	return &Box{
		buttion1: buttion1,
		buttion2: buttion2,
	}
}

func (b *Box) PressButtion1() {
	b.buttion1.Execute()
}

func (b *Box) PressButtion2() {
	b.buttion2.Execute()
}

func main() {
	mb := &Parents{}
	commandA := NewCommandA(mb)
	commandB := NewCommandB(mb)

	box1 := NewBox(commandA, commandB)
	box1.PressButtion1()
	box1.PressButtion2()

	box2 := NewBox(commandB, commandA)
	box2.PressButtion1()
	box2.PressButtion2()
}

//output:
//push A button execute command A
//push B button execute command B
//push B button execute command B
//push A button execute command A
