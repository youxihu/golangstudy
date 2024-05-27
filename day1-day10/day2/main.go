package main

import (
	"fmt"
	"time"
)

func main() {
	//processIfControl()
	//processForControl()
	//outPutOneHundred()
	//goTo()
	//cal()
	//calculate()
	//while()
	//jiaohuan()
	testTest()

}

func processIfControl() {
	a := 60
	var b float32
	fmt.Println("请输入你的成绩")
	_, err := fmt.Scanf("%f", &b)
	if err != nil {
		fmt.Println("输入无效,重新输入")
		return
	}

	if b < 0 || b > 100 {
		fmt.Println("输入无效,成绩必须在0到100之间")
		return
	}

	if a > int(b) {
		fmt.Printf("你的成绩是%.2f,你没有及格\n", b)
	} else if a == int(b) {
		fmt.Println("不错，及格了，继续努力")
	} else {
		fmt.Printf("你的成绩是%.2f,你很优秀\n", b)
	}

	studentScore := int(b) // 将用户输入的成绩转换为整数并赋给 studentScore

	switch {
	case studentScore >= 90:
		fmt.Printf("你的成绩是%d,成绩为A\n", studentScore)
	case studentScore >= 80:
		fmt.Printf("你的成绩是%d,成绩为A\n", studentScore)
	case studentScore >= 70:
		fmt.Printf("你的成绩是%d,成绩为B\n", studentScore)
	case studentScore >= 60:
		fmt.Printf("你的成绩是%d,成绩为B\n", studentScore)
	case studentScore >= 50:
		fmt.Printf("你的成绩是%d,成绩为C\n", studentScore)
	case studentScore >= 40:
		fmt.Printf("你的成绩是%d,成绩为C\n", studentScore)
	case studentScore >= 30:
		fmt.Printf("你的成绩是%d,成绩为D\n", studentScore)
	case studentScore >= 20:
		fmt.Printf("你的成绩是%d,成绩为D\n", studentScore)
	case studentScore >= 10:
		fmt.Printf("你的成绩是%d,成绩为E\n", studentScore)
	case studentScore >= 0:
		fmt.Printf("你的成绩是%d,成绩为E\n", studentScore)
	default:
		fmt.Printf("你的成绩是%d,输入无效\n", studentScore)
	}
}

func processForControl() {
	//var arr [15]int
	//var sum int
	for i := 0; i < 12; i++ {
		fmt.Println("我是吴仔妞的大爹")
		//fmt.Println(sum)
	}
	//for {
	//	fmt.Println("我是你爹")
	//}
}

func outPutOneHundred() {
	for i := 1; i <= 100; i++ {
		//if i%6 == 0 {
		//	fmt.Println(i)
		//}
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func goTo() {
	for i := 1; i <= 3; i++ {
		for j := 2; j <= 3; j++ {
			fmt.Printf("i:%v, j:%v\n", i, j)
			if i == 2 || j == 2 {
				//continue
				break
				//return
			}
		}
	}
}

func cal(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func calculate() {
	var num1, num2 int
	fmt.Println("请输入语文成绩：")
	fmt.Scanf("%d", &num1)
	fmt.Println("请输入数学成绩：")
	fmt.Scanf("%d", &num2)

	result := cal(num1, num2)
	fmt.Printf("你的分数总和是：%d\n", result)
}

func while() {
	for {
		fmt.Println("当前时间戳:", time.Now().Unix())
		time.Sleep(time.Second) // 每秒钟打印一次时间戳
		return
	}
}

func exchangeNum(num1 *int, num2 *int) {
	// 使用指针来交换变量的值
	var t int
	t = *num1
	*num1 = *num2
	*num2 = t
}

func jiaohuan() {
	var num1 int = 100
	var num2 int = 111
	fmt.Printf("交换前 两数分别是num1:%v,num2:%v\n", num1, num2)
	exchangeNum(&num1, &num2)
	fmt.Printf("交换后 两数分别是num1:%v,num2:%v\n", num1, num2)
}

func test(num int) {
	num = 30
	fmt.Println("test--------", num)
}

func testTest() {
	var num int = 10
	test(num)
	fmt.Println("testTest--------", num)
}
