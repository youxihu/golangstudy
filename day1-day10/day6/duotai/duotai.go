package duotai

import "fmt"

type SayHello interface {
	Say()
}

type NiuYangGe interface {
	Niu()
}

type chinese struct{}

type american struct{}

func (c chinese) Say() {
	fmt.Println("你好")
}

func (c chinese) Niu() {
	fmt.Println("东北文化扭秧歌")
}

func (a american) Say() {
	fmt.Println("hello")
}
func (a american) Niu() {
	fmt.Println("美国人不会扭秧歌")
}

func use(u SayHello) {
	u.Say()
	//if d, ok := u.(NiuYangGe); ok {
	//	d.Niu()
	//} else {
	//	fmt.Println("美国人不会扭秧歌")
	//}
	switch u.(type) {
	case chinese:
		ch := u.(chinese)
		ch.Niu()
	case american:
		am := u.(american)
		am.Niu()
	}
}

func Using() {
	c := chinese{}
	//c := chinese{}
	use(c)
}
