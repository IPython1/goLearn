package main

import "fmt"

func main() {
	var any interface{}
	var b int = 10
	var a *int
	any = a
	value, ok := any.(*int)
	fmt.Println(value) //nil空指针 没有指向任何内存地址
	fmt.Println(ok)
	fmt.Println(b)

	value1, ok1 := any.(int)
	fmt.Println(value1) //断言成功是正常值 断言失败是0
	fmt.Println(ok1)

	any = b
	value2, ok2 := any.(int)
	fmt.Println(value2)
	fmt.Println(ok2)

	//接口作函数参数

	//接口嵌套

}
