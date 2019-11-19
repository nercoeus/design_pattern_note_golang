package factory_method

import "fmt"

// 所有产品共同的功能
type Product interface {
	getName() string
}
// 卡车类
type Truck struct {
	name string
}
// 实现该类对象的相同功能
func (t *Truck)getName()string{
	return t.name
}
// 轮船类
type Ship struct {
	name string
}

func (s *Ship)getName()string{
	return s.name
}
// 工厂实现接口
type Factory interface {
	createProduct() Product
}
// 不同对象对应的工厂
type TruckFactory struct {
}
// 对工厂接口的实现,用来创建对应的对象
func(tf *TruckFactory)createProduct()Product{
	return &Truck{
		name:"This is a Truck",
	}
}

type ShipFactory struct{
}

func(sf *ShipFactory)createProduct()Product{
	return &Ship{
		name:"This is a Ship",
	}
}

func main(){
	tFactory := TruckFactory{}
	sFactory := ShipFactory{}
	truck := tFactory.createProduct()
	ship := sFactory.createProduct()
	fmt.Println()
	fmt.Println(truck.getName())
	fmt.Println(ship.getName())
}

//output：
//This is a Truck
//This is a Ship
