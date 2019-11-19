package observer

import "fmt"

// 目标对象，里面包含对应的订阅列表
type Publisher struct {
	observers []Observer
	context   string
}

// 创建新目标对象的工厂函数
func NewPublisher() *Publisher {
	return &Publisher{
		observers: make([]Observer, 0),
	}
}

// 添加观察者
func (s *Publisher) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// 通知函数，通知所有注册列表中的观察者
func (s *Publisher) Notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// 目标对象进行修改会调用 Notify 用来通知订阅列表
func (s *Publisher) Update(context string) {
	s.context = context
	s.Notify()
}

// 观察者接口，实现了 Update 用来供目标对象进行调用
type Observer interface {
	Update(*Publisher)
}

// 观察者类
type Subscriber struct {
	name string
}

// 创建观察者类的工厂函数
func NewSubscriber(name string) *Subscriber {
	return &Subscriber{
		name: name,
	}
}

// 实现 Observer 接口
func (r *Subscriber) Update(s *Publisher) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

// 测试函数的逻辑如下：
func main() {
	publisher := NewPublisher()
	subscriberA := NewSubscriber("SubscriberA")
	subscriberB := NewSubscriber("SubscriberB")
	subscriberC := NewSubscriber("SubscriberC")
	// 将观察者注册到对应的通知列表中
	publisher.AddObserver(subscriberA)
	publisher.AddObserver(subscriberB)
	publisher.AddObserver(subscriberC)

	publisher.Update("Publisher change")
}

//output:
//SubscriberA receive Publisher change
//SubscriberB receive Publisher change
//SubscriberC receive Publisher change
