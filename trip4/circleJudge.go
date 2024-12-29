package main

import (
	"fmt"
	"sort"
)

func main() {
	//for range 中的坑 go1.22之后就没有这个问题了
	arr := [2]int{1, 2}
	res := []*int{}
	for _, v := range arr {
		fmt.Println(&v)
		res = append(res, &v)
	}
	fmt.Println(*res[0], *res[1])
	//迭代变量时的闭包问题
	var funcs []func()

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f()
	}
	//修改切片中的元素
	/*修改迭代变量不会影响原始切片
	for range 会创建每个元素的副本
	而不是直接操作原始切片中的元素。*/
	arr2 := []int{1, 2, 3, 4, 5}
	for _, v := range arr2 {
		v *= 10
	}
	fmt.Println(arr2)
	//遍历字典时的顺序
	/*
			在 Go 中，使用 for range 遍历字典时，遍历顺序是随机的
			每次运行程序时，顺序可能不同
		解决方法：可以先对键进行排序，然后再遍历。

	*/
	dic := map[string]int{
		"a": 3,
		"b": 2,
		"c": 1,
	}
	keys := []string{}
	for k := range dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, dic[k])
	}
	//字符串遍历问题
	/*
			for range 遍历字符串时，每次迭代会返回 Unicode 代码点(rune)，
			而不是字节。如果字符串包含多字节字符，这一点尤其重要。
		解决方法：用常规的for循环
	*/

}
