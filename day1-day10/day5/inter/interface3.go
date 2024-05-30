package inter

import (
	"fmt"
)

// 一个自定义类型可以实现多个接口
type Ainterface interface {
	a()
}

type Binterface interface {
	b()
}

type youxihu struct {
}

func (y youxihu) a() {
	fmt.Println("aaaaaa")
}
func (y youxihu) b() {
	fmt.Println("bbbbbb")
}

func Stu() {
	var s youxihu
	var a Ainterface = s
	var b Binterface = s
	a.a()
	b.b()
}

// 一个接口可以继承多个别的接口
type Cinterface interface {
	Ainterface
	Binterface
	C()
}

type classroom struct {
}

func (c classroom) a() {
	fmt.Println("aaaaaa")
}
func (c classroom) b() {
	fmt.Println("bbbbbb")
}
func (c classroom) c() {
	fmt.Println("cccccc")
}

func Usejicheng() {
	var d classroom
	d.a()
	d.b()
	d.c()
}
