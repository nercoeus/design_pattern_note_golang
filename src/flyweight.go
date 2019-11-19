package flyweight

import (
	"fmt"
)

// 享元工厂类
type FlyweightFactory struct {
	maps map[string]*Flyweight
}

var flyweightFactory *FlyweightFactory

// 单例模式实现的 FlyweightFactory
func GetFlyweightFactory() *FlyweightFactory {
	if flyweightFactory == nil {
		flyweightFactory = &FlyweightFactory{
			maps: make(map[string]*Flyweight),
		}
	}
	return flyweightFactory
}

// 获取一个享元的指针
func (f *FlyweightFactory) Get(filename string) *Flyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}

// 一个享元对象
type Flyweight struct {
	data string
}

// 新的享元
func NewFlyweight(filename string) *Flyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &Flyweight{
		data: data,
	}
}

func (i *Flyweight) Data() string {
	return i.data
}

// 一个视图
type Viewer struct {
	*Flyweight
}

func NewViewer(filename string) *Viewer {
	image := GetFlyweightFactory().Get(filename)
	return &Viewer{
		Flyweight: image,
	}
}

func (i *Viewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}

// 用户逻辑
func main() {
	viewer1 := NewViewer("image1.png")
	viewer2 := NewViewer("image1.png")
	if viewer1.Flyweight == viewer2.Flyweight {
		fmt.Println("使用同一个享元")
	}

}

// output：使用同一个享元
