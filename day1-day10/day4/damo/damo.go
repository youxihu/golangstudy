package damo

import (
	"fmt"
	damo "gogo/day1-day10/day4/damo2"
)

type Student struct {
	Name string
}

// 定义方法
func (s Student) sy() {
	fmt.Println(s.Name)
}

// 定义函数
func StudentFunc(s Student) {
	fmt.Println(s.Name)
}

func ShiYon() {
	var s Student = Student{"youxihu"}
	StudentFunc(s)
	s.sy()
}

// 创建结构体指定字段值
type account struct {
	name string
	//age  int
}

func (a *account) test() {
	a.name = "youxihu"
	fmt.Println(a.name)
}
func AccountFuncFunc() {
	var a account
	a.name = "youxihu"
	a.test()
}

//	func AccountFunc() {
//		//按照顺序传值
//		a1 := account{"youxihu", 19}
//		fmt.Println(a1)
//		//指定类型传值
//		a2 := account{
//			name: "youxihu",
//			age:  19,
//		}
//		fmt.Println(a2)
//		//返回结构体对应指针
//		a3 := &account{"youxihu", 19}
//		fmt.Println(*a3)
//	}
func Admin() {
	x := damo.Admin{"youxihu", 20}
	fmt.Println(x)
	x1 := damo.Admin{
		Name: "youxihu",
		Age:  20,
	}
	fmt.Println(x1)
	x2 := &damo.Admin{"youxihu", 20}
	fmt.Println(*x2)
}
