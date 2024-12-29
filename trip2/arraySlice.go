package main

import "fmt"

func main() {
	//数组初始化

	//var array1 = [2]int{}
	//array1[0] = 1
	//array1[1] = 2
	//
	//array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(array2)
	//arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(arr2)
	//var strAry = [10]string{"aa", "bb", "cc", "dd", "ee"}

	//切片初始化
	//var slice1 = []int{}
	//slice1 = append(slice1, 1)
	//slice1 = append(slice1, 2)
	//fmt.Println(slice1)
	//slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(slice2)
	//[]string代表切片元素的类型 0 表示这个切片被初始化的长度 第三个参数控制容量（不写的话默认跟第二个参数大小一致）
	var sliceAry = make([]int, 2)
	sliceAry = append(sliceAry, 1)
	sliceAry = append(sliceAry, 2)
	sliceAry = append(sliceAry, 3) //go会自动处理容量不足的情况 不会导致错误
	fmt.Println(sliceAry)
	//字典(map)初始化 方式一
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	fmt.Println(m)

	var dic = map[string]int{ //方式二
		"apple":      1,
		"watermelon": 2,
		"banana":     3,
	}
	//fmt.Printf("strAry %+v\n", strAry)
	fmt.Printf("sliceAry %+v\n", sliceAry)
	fmt.Printf("dic %+v\n", dic)
}
