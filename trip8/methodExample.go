package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p *People) GetName() string {
	return p.Name
}

type Student struct {
	ID    int
	Score int
	People
}

func (st *Student) GetScore() int {
	return st.Score
}

func main() {
	st := &Student{ //直接定义成一个地址
		ID:    100,
		Score: 100,
		People: People{
			Name: "tom",
			Age:  18,
		},
	}
	fmt.Println(st.GetName())
	fmt.Println(st.GetScore())

	var a *int
	var i interface{}
	i = a
	value, ok := i.(int)
	fmt.Println(value)
	fmt.Println(ok)

}
