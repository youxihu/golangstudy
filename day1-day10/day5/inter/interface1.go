package inter

import "fmt"

type Alive interface {
	LongAlive()
	die()
}

type bbx struct {
}

type xyn struct {
}

func (b bbx) die() {}
func (y xyn) die() {}

func (b bbx) LongAlive() {
	fmt.Println("bbx卡住了")
}
func (y xyn) LongAlive() {
	fmt.Println("xyn好起来了")
}

func call(a Alive) {
	a.LongAlive()
}

func Inter() {
	bbx := bbx{}
	xyn := xyn{}
	//bbx.LongAlive()
	//xyn.LongAlive()
	call(bbx)
	call(xyn)
}
