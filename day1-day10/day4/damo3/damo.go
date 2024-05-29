package damo

// 封装

type exampleFZ struct {
	name string
	age  int
}

func FenPackage(name string) *exampleFZ {
	return &exampleFZ{
		name: name,
	}
}
