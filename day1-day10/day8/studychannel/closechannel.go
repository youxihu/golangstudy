package studychannel

import "fmt"

func CloseChannel() {
	chan1 := make(chan int, 3)
	chan1 <- 1
	chan1 <- 2
	close(chan1) //关闭管道可读不可写
	num1 := <-chan1
	num2 := <-chan1
	//chan1 <- 1
	fmt.Println(num1, num2)
}
