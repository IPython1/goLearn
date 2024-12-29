package main

import (
	"fmt"
	"math"
)

func getNumber() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	/*nextNumber为一个函数，函数中i为日*/
	nextNumber := getNumber()

	/*调用 nextNumber 函数，i变量自增1并返回*/
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/*创建新的函数 nextNumber1，并查看结果*/
	nextNumber1 := getNumber()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	/**
	声明函数变量
	*/
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x) //求一个数的平方根
	}
	fmt.Println(getSquareRoot(9))
}
