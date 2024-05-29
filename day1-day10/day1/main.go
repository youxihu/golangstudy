package main

import (
	"fmt"
	"reflect"
)

// 全局变量1
var author = "youxihu"
var year = "2024.05.24"

// 全局变量2
var (
	author2 = "youxihu"
	year2   = "2024.05.24"
)

func main() {
	//VarStudy()
	//TypeStudy()
	//calculate()
	ForScanLn()
}

func VarStudy() {
	//第一种类型 (指定类型并赋值)
	var num int = 100
	fmt.Println(num)
	//第二种类型 (不写类型 自动判断)
	num2 := 10
	fmt.Println(num2)
	//第三种类型 （不赋值使用默认值）
	var num3 int
	fmt.Println(num3)
	//第四种类型
	sex := "男"
	fmt.Println(sex)
	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	//声明多个变量 定义值
	var n4, name, age = 10, "jack", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(age)
	//声明多个变量 定义值
	n5, n6, n7 := 10, 1000, 100
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	//全局变量调用
	fmt.Println(author)
	fmt.Println(year)
	//全局变量调用2
	fmt.Println(author2)
	fmt.Println(year2)
}

func TypeStudy() {
	//基本数据类型
	//数值区分整数和浮点数
	var age = -23
	fmt.Println("age类型", reflect.TypeOf(age))
	var age2 = 127
	fmt.Printf("age类型 %T", age2)
	fmt.Println()
	var age3 = 4.1
	fmt.Println("age类型", reflect.TypeOf(age3))
	//字符型 使用byte来保存单个字符
	var c1 byte = 'a' //97
	fmt.Println(c1)
	var c2 byte = '6' //54
	fmt.Println(c2)
	var c3 byte = ')' //41
	fmt.Println(c3)
	//字符类型本质上就是一个整数 等同于uint8 地层是ascii码
	var c4 byte = '9' //57
	fmt.Println(c4)
	fmt.Println(c3 + c4) // 41+57=98
	var c5 int = '中'
	fmt.Printf("格式化输出c5的的具体字符为: %c", c5)
	fmt.Println()
	///转义字符
	fmt.Println("aaa\nbbb") //换行
	fmt.Println("aaa\bbbb") //退格
	fmt.Println("aaa\rbbb") //光标回到本行开头 后续输入替换原有字符
	fmt.Println("aaab\rbbaaab")
	fmt.Println("aaaaaaa\tbbbb") //每八个成一个制表符
	fmt.Println("aaaaaaaaa")
	fmt.Println(`\aaaa`)
	//布尔型 false true
	bool1 := 5 < 9
	fmt.Println(bool1)
	bool2 := 12 > 100
	fmt.Println(bool2)
	var bool3 bool = false
	fmt.Println(bool3)
	var bool4 bool = true
	fmt.Println(bool4)
	//字符串 string
	var s1 string = "你好,Golang"
	fmt.Println(s1)
	var s2 string = "ab c"
	fmt.Println(s2)
	var s3 string = `
	//全局变量调用
	fmt.Println(author)
	fmt.Println(year)
	//全局变量调用2
	fmt.Println(author2)
	fmt.Println(year2)
`
	fmt.Println(s3)
	var s4 string = "abc" + "eey"
	s4 += "hijk"
	fmt.Println(s4)
	var s5 string = "abc" + "abc" + "abc" +
		"abc" + "abc" + "abc" +
		"abc" + "abc" + "abc" +
		"abc" + "abc" + "abc" +
		"abc" + "abc" + "abc" +
		"abc" + "abc" + "abc"
	fmt.Println(s5)

	//类型转换
	var n1 int = 10
	//var n2 float32 = n1
	fmt.Println(n1)
	//fmt.Println(n2)
	var n2 float32 = float32(n1)
	fmt.Println(n2)
	fmt.Printf("%T\n", n1)
	//n1类型还是int类型 只是将值100 转换成了float32

	var q1 int = 10
	var q2 float32 = 4.01
	var q3 bool = true
	var q4 byte = 'a'

	var w1 string = fmt.Sprintf("%v", q1)
	fmt.Printf("w1对应的类型是:%T, 值是:%q\n", w1, w1)
	var w2 string = fmt.Sprintf("%v", q2)
	fmt.Printf("w2对应的类型是:%T, 值是:%q\n", w2, w2)
	var w3 string = fmt.Sprintf("%v", q3)
	fmt.Printf("w3对应的类型是:%T, 值是:%q\n", w3, w3)
	var w4 string = fmt.Sprintf("%v", q4)
	fmt.Printf("w4对应的类型是:%T, 值是:%q\n", w4, w4)

	//交换两个数并输出结果
	var data1 int = 9
	var data2 int = 3
	data1, data2 = data2, data1
	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)
	//复杂数据类型 派生数据类型
	//指针
	//数组
	//结构体
	//管道
	//函数
	//切片

	//接口
	//map
}

func calculate() {
	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(3 * 4)
	fmt.Println(4 / 2)
	fmt.Println(3 % 2)
	var num1 = (20+10)%3 + 7 - 6
	fmt.Println(num1)
	fmt.Println(5 == 9)
	fmt.Println(5 != 9)
	fmt.Println(5 > 9)
	fmt.Println(5 < 9)
	fmt.Println(5 >= 9)
	fmt.Println(5 <= 9)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(false || true)
	fmt.Println(!true)
	fmt.Println(!false)
	var un1 int = 18
	var ptr *int = &un1
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

func ForScanLn() {
	var age int
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	var name string
	//fmt.Println("请输入名字")
	//fmt.Scanln(&name)
	var score float32
	//fmt.Println("请输入成绩")
	//fmt.Scanln(&score)
	//fmt.Println("录入成功")
	//fmt.Printf("已录入成功，姓名:%v,年龄:%v,成绩:%v", name, age, score)
	fmt.Println("请录入学生的姓名、年龄、成绩，用空格隔开")
	fmt.Scanf("%s %d %f", &name, &age, &score)
	fmt.Printf("已录入成功，姓名:%v,年龄:%v,成绩:%v", name, age, score)
}
