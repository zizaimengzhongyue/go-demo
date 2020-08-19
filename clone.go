// 通过 Clone 的方式创建一个新的对象
package main

import (
	"fmt"
)

var list []Incre = []Incre{}

type Item struct {
	Index int
}

type Incre interface {
	Init(*Item)
	Increase()
	GetVal() int
	Clone() Incre
}

type Node struct {
	Val *Item
}

func (this *Node) Init(item *Item) {
	this.Val = item
}

func (this *Node) Increase() {
	this.Val.Index++
}

func (this *Node) GetVal() int {
	return this.Val.Index
}

func (this *Node) Clone() Incre {
	return &Node{}
}

type Node2 struct {
	Val *Item
}

func (this *Node2) Init(item *Item) {
	this.Val = item
}

func (this *Node2) Increase() {
	this.Val.Index += 2
}

func (this *Node2) GetVal() int {
	return this.Val.Index
}

func (this *Node2) Clone() Incre {
	return &Node2{}
}

func Init() {
	list = append(list, &Node{})
	list = append(list, &Node2{})
}

func main() {
	Init()
	for _, v := range list {
		item1 := &Item{Index: 1}
		item2 := &Item{Index: 10}
		tmp1 := v.Clone()
		tmp1.Init(item1)
		tmp2 := v.Clone()
		tmp2.Init(item2)
		tmp1.Increase()
		tmp2.Increase()
		fmt.Println(tmp1.GetVal(), item1.Index, tmp2.GetVal(), item2.Index)
	}
}
