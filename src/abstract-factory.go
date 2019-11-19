package main

import "fmt"

// 具体的不同家具的 struct
type Chair struct {   // Chair
	class string
	name  string
}
type Sofa struct {    // Sofa
	class string
	name  string
}
type Table struct {   // Table
	class string
	name  string
}

// 圆形家具工厂
type CircleFurnitureFactory struct {
}

// 矩形家具工厂
type RectangleFurnitureFactory struct {
}

// 三角形家具工厂
type TriangleFurnitureFactory struct {
}

// 抽象工厂类
type AbstractFurnitureFactory interface {
	MakeChair() Chair
	MakeSofa() Sofa
	MakeTable() Table
}

// 圆形工厂的家具生产过程
func (a CircleFurnitureFactory) MakeChair() Chair {
	return Chair{
		class: "Circle",
		name:  "Chair",
	}
}

func (a CircleFurnitureFactory) MakeSofa() Sofa {
	return Sofa{
		class: "Circle",
		name:  "Sofa",
	}
}

func (a CircleFurnitureFactory) MakeTable() Table {
	return Table{
		class: "Circle",
		name:  "Table",
	}
}

// 矩形工厂的家具生产过程
func (a RectangleFurnitureFactory) MakeChair() Chair {
	return Chair{
		class: "Rectangle",
		name:  "Chair",
	}
}

func (a RectangleFurnitureFactory) MakeSofa() Sofa {
	return Sofa{
		class: "Rectangle",
		name:  "Sofa",
	}
}

func (a RectangleFurnitureFactory) MakeTable() Table {
	return Table{
		class: "Rectangle",
		name:  "Table",
	}
}

// 三角形工厂的家具生产过程
func (a TriangleFurnitureFactory) MakeChair() Chair {
	return Chair{
		class: "Triangle",
		name:  "Chair",
	}
}

func (a TriangleFurnitureFactory) MakeSofa() Sofa {
	return Sofa{
		class: "Triangle",
		name:  "Sofa",
	}
}

func (a TriangleFurnitureFactory) MakeTable() Table {
	return Table{
		class: "Triangle",
		name:  "Table",
	}
}

// 选择自己需要的系列
func getFurnitureFactor(class string) AbstractFurnitureFactory {

	var res AbstractFurnitureFactory
	switch class {
	case "Circle":
		res = CircleFurnitureFactory{}
		return res
	case "Rectangle":
		res = RectangleFurnitureFactory{}
		return res
	case "Triangle":
		res = TriangleFurnitureFactory{}
		return res
	}
	return nil
}

// 主函数的逻辑
func main() {
	a := getFurnitureFactor("Circle")
	b := getFurnitureFactor("Rectangle")
	c := getFurnitureFactor("Triangle")
	fmt.Println(a.MakeChair())
	fmt.Println(a.MakeSofa())
	fmt.Println(a.MakeTable())
	fmt.Println(b.MakeChair())
	fmt.Println(b.MakeSofa())
	fmt.Println(b.MakeTable())
	fmt.Println(c.MakeChair())
	fmt.Println(c.MakeSofa())
	fmt.Println(c.MakeTable())
}

//output：
//{Circle Chair}
//{Circle Sofa}
//{Circle Table}
//{Rectangle Chair}
//{Rectangle Sofa}
//{Rectangle Table}
//{Triangle Chair}
//{Triangle Sofa}
//{Triangle Table}
