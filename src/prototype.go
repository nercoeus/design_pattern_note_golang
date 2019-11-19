package prototype

import "fmt"

// 产品的克隆接口
type Prototype interface {
	clone() Prototype
	getName() string
}

// 具体的产品类
type Product struct {
	name string
}

// 克隆自身
func (p *Product) clone() Prototype {
	return &Product{name: p.name}
}
func (p *Product) getName() string {
	return p.name
}

// Prototype Manager，原型管理器
type PrototypeManager struct {
	m map[string]Prototype
}

// 获取新的克隆对象
func (p *PrototypeManager) getProduct(s string) Prototype {
	return p.m[s].clone()
}

func (p *PrototypeManager) setProduct(s string, proto Prototype) {
	p.m[s] = proto
}

// 客户逻辑
func main() {
	a := Product{name: "Product A"}
	b := Product{name: "Product B"}
	c := Product{name: "Product C"}
	manager := PrototypeManager{m: make(map[string]Prototype)}
	manager.setProduct("A", &a)
	manager.setProduct("B", &b)
	manager.setProduct("C", &c)
	cloneA := manager.getProduct("A")
	cloneB := manager.getProduct("B")
	cloneC := manager.getProduct("C")
	fmt.Println(cloneA.getName())
	fmt.Println(cloneB.getName())
	fmt.Println(cloneC.getName())
}

//output:
//Product A
//Product B
//Product C
