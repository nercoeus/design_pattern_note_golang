package singleton

import (
	"fmt"
	"sync"
)

// 直接使用了 sync.Once 来进行实现，也可以采取加锁的方式，均可以
type singleton struct {
	name string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			name: "Singleton",
		}
	})
	return instance
}

func main() {
	fmt.Println(GetInstance().name)
	fmt.Println(GetInstance().name)
	fmt.Println(GetInstance().name)
}

//output:
//Singleton
//Singleton
//Singleton
