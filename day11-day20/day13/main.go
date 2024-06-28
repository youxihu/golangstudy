package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b int = 200
	testReflect(b)
}

func testReflect(i interface{}) {
	// reflect反射练习
	v := reflect.ValueOf(i)
	fmt.Println("Value:", v)
	t := reflect.TypeOf(i)
	fmt.Println("Type:", t)

	// 将 reflect.Value 转换为实际值并打印
	fmt.Println("Interface value:", v.Interface())

	// 如果要直接将 reflect.Value 转换为空接口，可以这样做：
	var emptyInterface interface{} = v.Interface()
	fmt.Println("Empty Interface value:", emptyInterface)

}
