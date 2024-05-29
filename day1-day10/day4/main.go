package main

import (
	"fmt"
)

func main() {
	//makeFunc()
	//makeFuncA()
	//cutFuncA()
	//cutCopy()
	//mapStudy()
	//mapRange()
	//map1Range()
	//objectStudy()
	//people1()
	//people2()
	//people3()
	//people4()
	student1()
	p.test()
}

func makeFunc() {
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Printf("切片的长度: %d, 切片的容量: %d\n", len(slice), cap(slice))
	slice[0] = 66
	slice[1] = 99
	slice[2] = 00
	slice[3] = 17
	slice[4] = 91
	fmt.Println(slice)
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	fmt.Printf("切片的长度: %d, 切片的容量: %d\n", len(slice2), cap(slice2))
}

func makeFuncA() {
	slice := make([]int, 5, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 10 // 为每个索引分配不同的值
		fmt.Printf("slice[%v] = %v \n", i, slice[i])
	}
	fmt.Println("---------------")
	for key, value := range slice {
		value = 10
		slice[key] = value
		fmt.Printf("切片的索引:%v,切片的值:%v\n", key, value)
	}
}

func cutFuncA() {
	var slice1 [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice1)
	var slice3 []int = slice1[0:4]
	fmt.Println(slice3)
	var slice4 = append(slice3, 88, 22)
	fmt.Println(slice4)
}

func cutCopy() {
	var attr1 = []int{1, 2, 3}
	fmt.Println(attr1)
	var attr2 = make([]int, 3)
	fmt.Println(attr2)
	copy(attr2, attr1)
	attr2[0] = 100
	fmt.Println(attr2)
}
func mapStudy() {
	//var map1 map[string]string
	var map1 = make(map[string]interface{}, 10)
	var student = []string{"yrzs", "youxihu"}
	map1["六班"] = student
	map1["四班"] = "小王"
	map1["三班"] = "小李"
	fmt.Println(map1)

	class := make(map[int]string)
	class[1] = "一班"
	class[2] = "二班"
	//delete(class, 1)
	//class = make(map[int]string)
	fmt.Println(class)
	value, flag := class[1]
	fmt.Println(value)
	fmt.Println(flag)

	age := map[int]string{
		18: "youxihu",
		12: "zhangzhic",
	}
	fmt.Println(age)
}
func mapRange() {
	map17 := make(map[string]int)
	map17["三班"] = 99
	map17["四班"] = 88
	map17["五班"] = 77
	for k, v := range map17 {
		fmt.Printf("map的key:%v,map的value:%v\n", k, v)
	}
}

func map1Range() {
	example1 := make(map[string]interface{})
	example1["yxh"] = 1
	example1["xh"] = 2
	fmt.Println(example1)
	classroom := make(map[int]map[int]string)
	classroom[1] = make(map[int]string, 3)
	classroom[1][1] = "yxh"
	classroom[1][2] = "zy"
	classroom[1][3] = "zzc"
	fmt.Println(classroom)
	for k1, v1 := range classroom {
		for k2, v2 := range v1 {
			fmt.Printf("k1是:%v,k2是:%v,v2是:%v\n", k1, k2, v2)
		}
	}
}

func objectStudy() {
	var name = []string{"yxh", "zzc", "zy"}
	handsomeName := name[1:3]
	fmt.Println(name)
	fmt.Println(handsomeName)
}

// 自定义结构体
type People struct {
	Name   string
	Age    int
	School string
	Sexy   string
}

func people1() {
	var p1 People
	fmt.Println(p1)
	p1.Name = "youxihu"
	p1.Age = 20
	p1.School = "清华大学"
	fmt.Println(p1)
}

func people2() {
	var p2 = People{"youxihu", 22, "清华大学", "男"}
	fmt.Println(p2)
}
func people3() {
	var p3 *People = new(People)
	p3.Name = "youxihu"
	p3.Age = 20
	p3.School = "清华大学"
	p3.Sexy = "男"
	fmt.Println(*p3)
}

func people4() {
	var p4 *People = &People{}
	p4.Name = "youxihu"
	p4.Age = 20
	fmt.Println(*p4)
}

// 结构体转换
type student struct {
	age int
}
type Stu student

func student1() {
	var s1 = student{19}
	var s2 = Stu{20}
	fmt.Println(s1, s2)
	s1 = student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

// 方法
type person struct {
	Name string
}

func (p person) test() {
	var p person
	p.Name = "youxihu"
	fmt.Println(p.Name)
}
