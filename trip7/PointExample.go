package main

import "fmt"

// go中的指针不能+1/-1跳到下一个地址
func main() {
	var a *int
	fmt.Printf("%v\n", a)
	fmt.Println(a == nil)
	b := new(int)
	fmt.Printf("%v\n", b)
	fmt.Printf("b:%d\n", *b)
	fmt.Println(b == nil)
	*b = 10
	fmt.Printf("b1:%d\n", *b)
	var c int = 10
	a = &c
	fmt.Println("a的地址是:", a)
	b = &c
	fmt.Println("b的地址是:", b)
	fmt.Println(*a)
	fmt.Println(*b)
}
