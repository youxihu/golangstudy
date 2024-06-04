package studylock

import (
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup
var Count int

func AddAdd() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		lock.Lock()
		Count += 1
		lock.Unlock()
	}
}

func SubSub() {
	defer wg.Done()
	for u := 0; u < 10; u++ {
		lock.Lock()
		Count -= 1
		lock.Unlock()
	}
}

func Cal() {
	wg.Add(2)
	go AddAdd()
	go SubSub()
	wg.Wait() // 等待所有子例程完成
}
