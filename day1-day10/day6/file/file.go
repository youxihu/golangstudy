package file

import (
	"fmt"
	"os"
)

func OpenFile() {
	//打开文件
	file, err := os.Open("/workserver/work-local/work-study/gogo/day1-day10/day6/file/a.txt")
	if err != nil {
		fmt.Println("打开文件出错，问题是:", err)
	} else {
		fmt.Println("文件对象是:", &file)
	}
	//关闭文件
	err2 := file.Close()
	if err2 != nil {
		fmt.Println("关闭的错误原因是:", err)
	}
}
