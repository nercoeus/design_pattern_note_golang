package composite

import "fmt"

// Component 接口，其中定义了一系列容器和叶子的操作
type Component interface {
	Parent() Component    // 基本操作
	SetParent(Component)  // 基本操作
	Name() string         // 基本操作
	SetName(string)       // 基本操作
	AddChild(Component)   // 容器操作
	Print(string)         // 容器、叶子操作
}

const (
	LeafNode = iota
	CompositeNode
)

// 创建一个新的节点
func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}
// 通用节点对象
type component struct {
	parent Component
	name   string
}
// 通用节点基本方法的实现
func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}
// 叶子类型
type Leaf struct {
	component
}
// 创建叶子节点的工厂函数
func NewLeaf() *Leaf {
	return &Leaf{}
}
// 叶子类型接口的实现
func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
}
// 容器类型
type Composite struct {
	component
	childs []Component
}
// 创建容器节点的工厂函数
func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}
// 容器类型接口的实现
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}

// 测试用例，构建了一棵树并进行打印操作
func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print(" ")
}

//output:
//+root
//  +c1
//    +c3
//    -l1
//  +c2
//    -l2
//    -l3
