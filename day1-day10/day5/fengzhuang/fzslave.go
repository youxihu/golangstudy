package fengzhuang

import "fmt"

type score struct {
	Name  string
	score int
}

func NewScore(name string) *score {
	return &score{
		Name: name,
	}
}

func (s *score) SetScore(score int) {
	if score > 0 && score < 150 {
		s.score = score
	} else {
		fmt.Println("你输入的成绩范围不正确")
	}
}

func (s *score) GetScore() int {
	return s.score
}
