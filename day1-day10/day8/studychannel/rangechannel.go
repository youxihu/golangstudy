package studychannel

import (
	"fmt"
)

func RangeChannel() {
	chan2 := make(chan int, 101)
	for i := 0; i <= 100; i++ {
		chan2 <- i
	}
	close(chan2)
	//管道遍历只有 for range，没有for i ..索引
	for v := range chan2 {
		fmt.Printf("value= %d\n", v)
	}
}
