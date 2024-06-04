package main

import (
	"fmt"
	"time"
)

//	func TestGoRoutine() {
//		for i := 0; i < 15; i++ {
//			fmt.Println("hello golang +", i)
//			time.Sleep(1 * time.Second)
//		}
//	}
func main() {
	//go TestGoRoutine() //开启协程
	TestGoRoutine2()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("hello world +", i)
	//	time.Sleep(1 * time.Second)
	//}
	//for i := 1; i <= 5; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//time.Sleep(2 * time.Second)
}

func TestGoRoutine2() {
	for i := 1; i <= 5; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
