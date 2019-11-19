package strategy

import "fmt"

type NavigationStrategy interface {
	Route(*Context)
}

// 用来存储上下文状态的结构体
type Context struct {
	From     string
	To       string
	Strategy NavigationStrategy
}

// 创建上下文环境
func NewContext(from, to string, payment NavigationStrategy) *Context {
	return &Context{
		From:     from,
		To:       to,
		Strategy: payment,
	}
}

func (p *Context) Route() {
	p.Strategy.Route(p)
}

// Bus 策略，实现了 NavigationStrategy 接口
type Bus struct{}

func (*Bus) Route(ctx *Context) {
	fmt.Printf("Bus (%s => %s) Spending Time: 10h, Spending Money: 200￥\n", ctx.From, ctx.To)
}

// Ship 策略，实现了 NavigationStrategy 接口
type Ship struct{}

func (*Ship) Route(ctx *Context) {
	fmt.Printf("Ship (%s => %s) Spending Time: 4h, Spending Money: 500￥\n", ctx.From, ctx.To)

}

// 用户使用逻辑
func main() {
	// 采用 BUS 进行路线花费判断
	ctx := NewContext("Beijing", "Shanghai", &Bus{})
	ctx.Route()
	// 采用 SHIP 进行路线花费的判断
	ctx = NewContext("Beijing", "Shanghai", &Ship{})
	ctx.Route()
}

//output:
//Bus (Beijing => Shanghai) Spending Time: 10h, Spending Money: 200￥
//Ship (Beijing => Shanghai) Spending Time: 4h, Spending Money: 500￥
