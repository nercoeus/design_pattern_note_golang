package facade

import "fmt"

func FacadeAPI() API {
	return &FacadeAPIImpl{
		a: NewAClass(),
		b: NewBClass(),
	}
}

type API interface {
	FacadeOutput() string
}

type FacadeAPIImpl struct {
	a APIclassA
	b APIclassB
}

func (f *FacadeAPIImpl) FacadeOutput() string {
	res1 := f.a.AOutput()
	res2 := f.b.BOutput()
	return fmt.Sprintf("%s\n%s", res1, res2)
}

func NewAClass() APIclassA {
	return &classA{}
}

type APIclassA interface {
	AOutput() string
}

type classA struct{}

func (*classA) AOutput() string {
	return "class A output"
}

func NewBClass() APIclassB {
	return &classB{}
}

type APIclassB interface {
	BOutput() string
}

type classB struct{}

func (*classB) BOutput() string {
	return "class B output"
}

func main() {
	facadeAPI := FacadeAPI()
	facadeRes := facadeAPI.FacadeOutput()
	a := NewAClass()
	b := NewBClass()
	normalRes := a.AOutput() + "\n" + b.BOutput()
	fmt.Println("--- facade output ---")
	fmt.Println(facadeRes)
	fmt.Println("--- normal output ---")
	fmt.Println(normalRes)
}

//output；
//--- facade output ---
//class A output
//class B output
//--- normal output ---
//class A output
//class B output
