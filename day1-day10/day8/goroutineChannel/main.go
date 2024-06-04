package main

import (
	"fmt"
	"sync"
	"time"
)

var Wg1 sync.WaitGroup

func WriteData(channel1 chan int) {
	defer Wg1.Done() //方法执行完后 才执行defer  done一次完成一个任务
	for i := 0; i <= 10; i++ {
		channel1 <- i
		fmt.Println("write data", i)
		time.Sleep(time.Second)
	}
	close(channel1)
}

func ReadData(channel1 chan int) {
	defer Wg1.Done()
	for v := range channel1 {
		fmt.Println(v)
		time.Sleep(time.Second)
	}
}

func main() {
	channel1 := make(chan int, 50)
	Wg1.Add(1)
	go WriteData(channel1)
	go ReadData(channel1)
	Wg1.Wait()
}
