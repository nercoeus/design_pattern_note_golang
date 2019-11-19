package memento

import "fmt"

// 复杂接口，这里并没有实现
type Memento interface{}

// 简单接口，给管理器使用的接口
type MementoSimple interface {
	Copy() MementoSimple
}

// 原始对象
type User struct {
	id    int
	state string
}

// 修改原始对象状态
func (u *User) ChangeState(s string) {
	u.state = s
}

// 产生备忘录
func (u *User) SaveMemento() *UserMemento {
	return &UserMemento{
		state: u.state,
	}
}

// 加载备忘录
func (u *User) LoadMemento(m *UserMemento) {
	u.state = m.state
}

func (u *User) PrintState() {
	fmt.Printf("User: %d state : %s\n", u.id, u.state)
}

// 备忘录对象
type UserMemento struct {
	state string
}

// mementosimple 接口实现
func (obj *UserMemento) Copy() MementoSimple {
	return &UserMemento{
		state: obj.state,
	}
}

// 备忘录管理类
type Caretaker struct {
	mementoMap map[int]MementoSimple
}

// 保存备忘录
func (c *Caretaker) SaveMemento(i int, m MementoSimple) {
	c.mementoMap[i] = m
}

// 获取备忘录
func (c *Caretaker) GetMemento(i int) MementoSimple {
	return c.mementoMap[i]
}

// 测试逻辑
func main() {
	// 原始对象
	user := User{
		id:    1,
		state: "old state",
	}
	// 备忘录管理器
	caretaker := Caretaker{
		mementoMap: map[int]MementoSimple{},
	}
	fmt.Println("user id: ", user.id, " state: ", user.state)
	// 生成一个备忘录
	oldState := user.SaveMemento()
	caretaker.SaveMemento(user.id, oldState)
	// 修改原始对象
	user.ChangeState("new state")
	fmt.Println("user id: ", user.id, " state: ", user.state)
	// 从 Caretaker 获取旧的对象
	oldState = caretaker.GetMemento(user.id).(*UserMemento)
	// 恢复旧的状态
	user.LoadMemento(oldState)
	fmt.Println("user id: ", user.id, " state: ", user.state)
}

//output:
//user id:  1  state:  old state
//user id:  1  state:  new state
//user id:  1  state:  old state
