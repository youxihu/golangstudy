package studylock

import (
	"fmt"
	"sync"
)

var rwlock sync.RWMutex

func RLock() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("开始读取数据")
	rwlock.Unlock()
	fmt.Println("读取数据成功")
}

func WLock() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("开始写入数据")
	rwlock.Unlock()
	fmt.Println("写入数据成功")
}

func RUnlock() {
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go RLock()
	}
	go WLock()
	wg.Wait()
	fmt.Println("运行结束")
}
