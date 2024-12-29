package trip1

import "fmt"

func main() {
	//变量的定义 =必须与var同时出现 或者修改值的时候可以用=
	var a int //默认为0
	a = 4
	var b int = 2
	//自行推导类型
	var c = 3
	fmt.Println(a, b, c)
	d := 5
	fmt.Println(d)
	//可一次性定义多个变量
	num1, num2 := 10, 11
	fmt.Println(num1, num2)
}
