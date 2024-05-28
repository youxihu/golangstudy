package main

import (
	"errors"
	"fmt"
)

func main() {
	//study()
	//errDiscover()
	//err := diyErr()
	//if err != nil {
	//	fmt.Println("错误是:", err)
	//}
	//ScoreAvg()
	//avgScore()
	//scoreAvg()
	//collRange()
	//anonymous()
	//getArr()
	//rangeArrArr()
	cutStudy()
}
func diyErr() (err error) {
	num3 := 10
	num4 := 0
	if num4 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num3 / num4
		fmt.Println(result)
		return nil
	}
}

func study() {
	str := "golang"
	fmt.Println(len(str))
	fmt.Println(str)
	num := new(int)
	fmt.Printf("num的类型是是森么:%T，num的值:%v,num的地址是:%v,num指针对应的是:%v\n", num, num, &num, *num)
}

func errDiscover() {
	//defer+recover
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经处理")
			fmt.Println("err discover:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func ScoreAvg() {
	var score1, score2, score3, score4, score5 int64
	fmt.Println("请输入语文，数学，英语，地理，物理的成绩")
	fmt.Print("语文成绩：")
	fmt.Scan(&score1)
	fmt.Print("数学成绩：")
	fmt.Scan(&score2)
	fmt.Print("英语成绩：")
	fmt.Scan(&score3)
	fmt.Print("地理成绩：")
	fmt.Scan(&score4)
	fmt.Print("物理成绩：")
	fmt.Scan(&score5)

	sum := score1 + score2 + score3 + score4 + score5
	avg := sum / 5
	fmt.Println("平均分：", avg)
}
func avgScore() {
	scores := [20]int{100, 77, 32, 57, 74, 64, 100, 99, 79, 69, 66, 71, 70, 56, 98, 64, 88, 85, 78, 97}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Println(len(scores))
	fmt.Printf("scores0地址是:%p\n", &scores[0])
	fmt.Printf("scores1地址是:%p\n", &scores[1])
	fmt.Printf("scores2地址是:%p\n", &scores[2])
	fmt.Printf("scores3地址是:%p\n", &scores[3])
	fmt.Printf("scores4地址是:%p\n", &scores[4])
	avg := (sum) / (len(scores))
	fmt.Printf("班级学生成绩总和是:%v,平均分是:%v\n", sum, avg)
}
func scoreAvg() {
	var scoresS [7]int
	sum := 0
	for i := 0; i < len(scoresS); i++ {
		fmt.Printf("请输入第%d门科目的成绩: ", i+1)
		fmt.Scanf("%d", &scoresS[i])
		sum += scoresS[i]
		fmt.Printf("第%d门成绩是:%d\n", i+1, scoresS[i])
	}
	avg := sum / len(scoresS)
	fmt.Printf("成绩总和是:%v,平均分是:%v\n", sum, avg)
}
func collRange() {
	var coll [6]int
	for index, _ := range coll {
		fmt.Printf("请输入第%d个同学的成绩:", index+1)
		fmt.Scanf("%d", &coll[index])

	}
	fmt.Println("所有同学的成绩:", coll)
}

func anonymous() {
	//定义匿名函数,定义的同时调用
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(result)

	//将匿名函数赋给变量
	sub := func(n3 int, n4 int) int {
		return n3 - n4
	}
	result01 := sub(30, 70)
	fmt.Println(result01)
	fmt.Println(func01(30, 20))
}

// 全局变量
var func01 = func(c1, c2 int) int {
	return c1 + c2
}

func getArr() {
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
	var arr2 = [...]int{3, 4, 5, 4}
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)
	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3)
	var arr4 = [...]int{7, 8, 9, 6}
	fmt.Println(arr4)
}
func rangeArrArr() {
	var arr5 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr5)
	fmt.Println("----------------")
	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Println(arr5[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("------------------")
	for key, value := range arr5 {
		for k, v := range value {
			fmt.Printf("key:[%v]  k:[%v]  v:%v\n", key, k, v)
			fmt.Println("key就代表[第一个下标],k就代表[第二个下标],v就代表[value代表的数组里的值]\n")
		}
	}
}

func cutStudy() {
	var intArr = [...]int{1, 2, 3}
	var sliceArr []int = intArr[:3]
	fmt.Println(sliceArr)
	fmt.Printf("sliceArr元素的个数:%v\n", len(sliceArr))
	fmt.Printf("sliceArr元素的容量:%v\n", cap(sliceArr))
	fmt.Printf("sliceArr下表为1的地址是:%v\n", &sliceArr[1])
	fmt.Printf("sliceArr下表为2的地址是:%v\n", &sliceArr[2])
	fmt.Printf("sliceArr下表为0的地址是:%v\n", &sliceArr[0])
	sliceArr[1] = 16
	fmt.Println(sliceArr)
}
