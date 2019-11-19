package iterator

import "fmt"

// 容器为切片类型
type Ints []int

// 迭代器通用接口
type interator interface {
	HasNext() bool
	Next() int
	Index() int
	BehindNum() int
}

// 返回切片的迭代器
func (i Ints) Iterator() *Iterator {
	return &Iterator{
		data:  &i,
		index: 0,
	}
}

// 迭代器类型
type Iterator struct {
	data  *Ints // 保存目标容器指针
	index int
}

// 迭代器可执行操作
func (i *Iterator) HasNext() bool {
	return i.index < len(*i.data)
}

func (i *Iterator) Next() (v int) {
	v = (*i.data)[i.index]
	i.index++
	return v
}
func (i *Iterator) Index() int {
	return i.index
}

func (i *Iterator) BehindNum() int {
	return len(*i.data) - i.index - 1
}

// 用户对容器的遍历
func main() {
	ints := Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for it := ints.Iterator(); it.HasNext(); it.Next() {
		fmt.Println("data: ", (*it.data)[it.index], " index: ", it.Index(), " behind num: ", it.BehindNum())
	}
}

//output:
//data:  1  index:  0  behind num:  9
//data:  2  index:  1  behind num:  8
//data:  3  index:  2  behind num:  7
//data:  4  index:  3  behind num:  6
//data:  5  index:  4  behind num:  5
//data:  6  index:  5  behind num:  4
//data:  7  index:  6  behind num:  3
//data:  8  index:  7  behind num:  2
//data:  9  index:  8  behind num:  1
//data:  10  index:  9  behind num:  0
