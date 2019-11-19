package adapter

import "fmt"

// 定义一个新的接口
type Target interface {
	NewAPI()
}

// 定义一个老的接口
type Adaptee interface {
	OldAPIA()
	OldAPIB()
}

// 定一个Adapter接口
type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) NewAPI(){
	a.adaptee.OldAPIA()
	a.adaptee.OldAPIB()
	fmt.Println("NewAPI = OldAPIA + OLDAPIB")
}

// 实现旧接口struct
type oldAdaptee struct {
}

func (o *oldAdaptee) OldAPIA(){
	fmt.Println("Old API: A")
}

func (o *oldAdaptee) OldAPIB(){
	fmt.Println("Old API: B")
}

func main(){
	oa := &oldAdaptee{}
	// 定一个适配器对象
	adapter := Adapter{
		adaptee:oa,
	}
	adapter.NewAPI()
}

//output：
//Old API: A
//Old API: B
//NewAPI = OldAPIA + OLDAPIB
