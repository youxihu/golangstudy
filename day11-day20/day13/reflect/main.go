package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Student{
		Name: "喜羊羊",
		Age:  18,
	}
	studyReflect(stu)
}

func studyReflect(i interface{}) {
	reVal := reflect.ValueOf(i)
	fmt.Println(reVal)

	reTyp := reflect.TypeOf(i)
	fmt.Println(reTyp)

	// 使用 Interface() 方法获取实际值，并进行类型断言
	i2 := reVal.Interface()
	n, flag := i2.(Student)
	if flag {
		fmt.Println("Student:", n.Name, n.Age)
	} else {
		fmt.Println("Failed to assert as Student type")
	}
}

type Student struct {
	Name string
	Age  int
}
