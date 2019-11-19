package visitor

import "fmt"

type Hole interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Hole)
}

type Orchard struct {
	holes []Hole
}

func (o *Orchard) Add(hole Hole) {
	o.holes = append(o.holes, hole)
}

func (o *Orchard) Accept(visitor Visitor) {
	for _, hole := range o.holes {
		hole.Accept(visitor)
	}
}

type OrangeHole struct {
	name string
}

func NewOrangeHole(name string) *OrangeHole {
	return &OrangeHole{
		name: name,
	}
}

func (c *OrangeHole) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type AppleHole struct {
	name string
}

func NewAppleHole(name string) *AppleHole {
	return &AppleHole{
		name: name,
	}
}

func (c *AppleHole) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type Farmer struct {
	name string
}

func (f *Farmer) Visit(hole Hole) {
	switch c := hole.(type) {
	case *AppleHole:
		fmt.Println(f.name, " planted ", c.name)
	case *OrangeHole:
		fmt.Println(f.name, " planted ", c.name)
	}
}

func main() {
	c := &Orchard{}
	c.Add(NewAppleHole("apple tree A"))
	c.Add(NewAppleHole("apple tree B"))
	c.Add(NewAppleHole("apple tree C"))
	c.Add(NewAppleHole("apple tree D"))
	c.Add(NewOrangeHole("orange tree A"))
	c.Add(NewOrangeHole("orange tree B"))
	c.Add(NewOrangeHole("orange tree C"))
	c.Add(NewOrangeHole("orange tree D"))
	c.Accept(&Farmer{name: "nercoeus"})
}

//output:
//nercoeus  planted  apple tree A
//nercoeus  planted  apple tree B
//nercoeus  planted  apple tree C
//nercoeus  planted  apple tree D
//nercoeus  planted  orange tree A
//nercoeus  planted  orange tree B
//nercoeus  planted  orange tree C
//nercoeus  planted  orange tree D
