package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// go语言函数参数传递只有值传递
func update(arr [2]int) {
	arr[0] = 0
}
func update2(arr1 []int) {
	arr1[0] = 0
}
func print() {
	fmt.Println("hi")
}
func main() {
	//a, b := swap("hello", "world")
	arr := [...]int{1, 2}
	fmt.Println(arr)
	update(arr)
	fmt.Println(arr)

	fmt.Println()

	arr1 := []int{1, 2} //切片的底层是一个结构体 结构体中有指针 在函数传递的过程中指针也被复制过去了 所以值被修改了
	fmt.Println(arr1)
	update2(arr1)
	fmt.Println(arr1)
	//fmt.Println(a, b)
	fmt.Println()
	fmt.Println("hello")
	print()

	num := 100
	word := "google"
	//匿名函数
	func(i int) {
		tmp := 5
		fmt.Println(i)
		fmt.Println(word)
		fmt.Println(tmp)
	}(num)
}
