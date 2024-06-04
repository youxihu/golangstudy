package studygoroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func FiveGoroutines() {
	for i := 1; i <= 7; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//time.Sleep(time.Second * 2)
}
