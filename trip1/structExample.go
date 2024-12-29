package trip1

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func main() {
	//var st = Student{
	//	ID:    1,
	//	Name:  "Tom",
	//	Age:   18,
	//	Score: 100,
	//}
	st := Student{
		ID:    1,
		Name:  "Tom",
		Age:   18,
		Score: 100,
	}
	fmt.Println(st)
	//值列表初始化需要1-1对应且个数相同
	st1 := Student{1, "Tom", 18, 100}
	fmt.Println(st1)
	//结构体成员的访问
	fmt.Println(st.ID, st.Name, st.Age, st.Score)
	//struct嵌套
	type Teacher struct {
		Name string
		Age  int
	}
	type Student struct {
		ID    int
		Name  string
		Age   int
		Score int
		Teacher
	}
	st2 := Student{
		ID:    1,
		Name:  "Tom",
		Age:   18,
		Score: 100,
		Teacher: Teacher{
			Name: "Jerry",
			Age:  20,
		},
	}
	fmt.Println(st2.Teacher.Name)
}
