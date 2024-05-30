package jicheng

import "fmt"

type Animal struct {
	age    int
	weight float32
}

func (a *Animal) shut() {
	fmt.Println("I am a big fat cat, barking every day")
}

func (a *Animal) show() {
	fmt.Printf("hat is the age of an animal:%v,What is the weight of an animal:%.1f\n", a.age, a.weight)
}

type cat struct {
	Animal
	age int
}

func (c *cat) show() {
	fmt.Printf("～～～～～～～～What is the age of an animal:%v,What is the weight of an animal:%.1f\n", c.age, c.weight)
}

func (c *cat) eat() {
	fmt.Println("I am a big fat cat, and I love to eat canned cats")
}

func DiaoyonJc() {
	cat := &cat{}
	cat.weight = 10.5
	cat.age = 3
	cat.Animal.age = 14
	cat.shut() //就近原则
	cat.Animal.show()
	cat.show()
	cat.eat()
}

// 多继承
type A struct {
	a int
	b string
}
type B struct {
	a int
	b string
}
type C struct {
	A
	B
}

func D() {
	d1 := C{A{12, "aaa"}, B{12, "vvv"}}
	fmt.Println(d1)
}

type occupation struct {
	Player string
	age    int
	a      int
}

type school struct {
	player string
	age    int
}

type littileBoy struct {
	occupation
	school
	a int
}

func Jack() {
	jack := littileBoy{
		occupation{
			"footballPlayer", 17,
			12,
		},
		school{"EnglishMaster",
			17,
		},
		11,
	}
	fmt.Println(jack)
	fmt.Println(jack.occupation)
	fmt.Println(jack.school)
	fmt.Println(jack.a)
}
