package state

import "fmt"

// 周一到周日所有天均实现的接口
type Day interface {
	Today()
	Next(*DayContext)
}

// 日历的 Context 类
type DayContext struct {
	today Day
}

// 创建 DayContext 的工厂函数
func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

// DayContext 打印今天的日期
func (d *DayContext) Today() {
	d.today.Today()
}

// 下一天，状态转换
func (d *DayContext) Next() {
	d.today.Next(d)
}

// 下面依次实现了不同天数的对应操作
// Sunday
type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Today is Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

// Monday
type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Today is Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

// Tuesday
type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Today is Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

// Wednesday
type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Today is Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

// Thursday
type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Today is Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

// Friday
type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Today is Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

// Saturday
type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Today is Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}

// 用户逻辑，随着天数的前进，日期进行转换
func main() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
}

//output：
//Today is Sunday
//Today is Monday
//Today is Tuesday
//Today is Wednesday
//Today is Thursday
//Today is Friday
//Today is Saturday
//Today is Sunday
