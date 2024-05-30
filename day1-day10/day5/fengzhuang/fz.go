package fengzhuang

import "fmt"

func DiaoyonFz() {
	d := NewScore("youxihu")
	d.SetScore(20)
	fmt.Println(d.Name)
	fmt.Println(d.GetScore())
	fmt.Println(*d)
}
