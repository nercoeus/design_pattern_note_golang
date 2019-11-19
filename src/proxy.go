package proxy

import "fmt"

// 代理和被代理对象均实现的接口，用来给用户进行使用
type ServiceInterface interface {
	Do()
}

// 被代理对象
type Service struct{}

func (Service) Do() {
	fmt.Println("被代理对象中执行操作")
}

// 代理对象，和用户以及被代理对象进行交互
type Proxy struct {
	ser Service
}

// 通过在代理类中使用被代理对象的接口
func (p Proxy) Do() {
	fmt.Println("调用代理之前可以执行一些操作")
	// 调用真实对象
	p.ser.Do()
	fmt.Println("调用代理之后可以执行一些操作")
}

// 用户逻辑
func main() {
	var service ServiceInterface
	service = &Proxy{}
	service.Do()
}

//output;
//调用代理之前可以执行一些操作
//被代理对象中执行操作
//调用代理之后可以执行一些操作
