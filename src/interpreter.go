package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

// 节点通用接口
type Node interface {
	Interpret() int
}

// 值节点
type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

// 加号节点
type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

// 减号节点
type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

// AST
type Parser struct {
	exp   []string
	index int
	prev  Node // AST root 节点
}

// 解析函数
func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

// 创建新的加号节点
func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

// 创建新的减号节点
func (p *Parser) newMinNode() Node {
	p.index++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

// 创建值节点
func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: v,
	}
}

func (p *Parser) Result() Node {
	return p.prev
}

// 用户逻辑
func main() {
	// 三个待解析的语句
	test1 := "1 + 2 + 3 - 4 + 5 - 6"
	test2 := "6 - 9 + 7 + 1 + 3 + 8"
	test3 := "0 + 9 + 8 - 7 + 6 - 5"
	p1 := &Parser{}
	// 进行解析
	p1.Parse(test1)
	// 计算结果
	res1 := p1.Result().Interpret()
	fmt.Println(test1, " = ", res1)
	p2 := &Parser{}
	p2.Parse(test2)
	res2 := p2.Result().Interpret()
	fmt.Println(test2, " = ", res2)
	p3 := &Parser{}
	p3.Parse(test3)
	res3 := p3.Result().Interpret()
	fmt.Println(test3, " = ", res3)
}

//output:
//1 + 2 + 3 - 4 + 5 - 6  =  1
//6 - 9 + 7 + 1 + 3 + 8  =  16
//0 + 9 + 8 - 7 + 6 - 5  =  11
