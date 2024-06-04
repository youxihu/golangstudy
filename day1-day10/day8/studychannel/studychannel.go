package studychannel

import "fmt"

func Chan() {
	var ch chan int = make(chan int, 3)
	ch <- 1
	num := 2
	ch <- num
	fmt.Printf("ch的长度:%v,ch的容量:%v\n", len(ch), cap(ch))
	fmt.Println(ch)
	num1 := <-ch
	num2 := <-ch
	fmt.Println(num1, num2)
}
