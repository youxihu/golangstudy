package studyio

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读1
func TestIo() {
	file, err := ioutil.ReadFile("/workserver/work-local/work-study/gogo/day1-day10/day7/examplefile/test")

	if err != nil {
		fmt.Printf("出现了问题:%v", err)
	} else {
		fmt.Printf("文件内容是:\n%v", string(file))
	}
}

// 读2
func TestIo2() {
	file2, err := os.Open("/workserver/work-local/work-study/gogo/day1-day10/day7/examplefile/test")
	if err != nil {
		fmt.Printf("打开失败:%v", err)
	}
	//当函数退出时,让file关闭,防止内存泄露
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {

		}
	}(file2)
	//创建一个流
	reader := bufio.NewReader(file2)
	//读取操作
	for {
		file3, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF代表读完了
			break
		}
		fmt.Println(file3)
	}
}

// 写1
func TestIo3() {
	file4, err := os.OpenFile("/workserver/work-local/work-study/gogo/day1-day10/day7/examplefile/test", os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开失败:%v", err)
		return
	}
	defer func(file4 *os.File) {
		err := file4.Close()
		if err != nil {
			return
		}
	}(file4)
	write := bufio.NewWriter(file4)
	_, err = write.WriteString("我是你爹")
	if err != nil {
		return
	}
	err = write.Flush()
	if err != nil {
		return
	}
}

func TestCopy() {
	path1 := "/workserver/work-local/work-study/gogo/day1-day10/day7/examplefile/test1"
	path2 := "/workserver/work-local/work-study/gogo/day1-day10/day7/examplefile/test2"
	if _, err := os.Stat(path2); os.IsNotExist(err) {
		_, _ = os.Create(path2)
	} else {
		fmt.Println("path2早有了")
	}
	file6, err := ioutil.ReadFile(path1)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path2, file6, os.ModeAppend)
}
