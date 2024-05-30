package inter

import "fmt"

type integer int

func (i integer) test() {
	fmt.Println("say hi+ ", i)
}
func InterR() {
	var a integer = 10
	a.test()
}
