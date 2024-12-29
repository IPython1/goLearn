package trip1

import "fmt"

// 常量一般用于对枚举值进行定义 提高程序的可读性
// 0000 0011 ->0000 1100 4+8=12
func const_example() {
	const (
		x = 1 << iota //iota=0 i=1<<0=3*2^0,iota++
		y = 3 << iota //iota=1 i=3<<1=3*2^1,iota++
		z             //iota=2 i=3<<2=3*2^2,iota++
		q
	)
	fmt.Println(x, y, z, q)
	const (
		a = iota //默认初始值为0
		b
		c
		d = "ha" //独立值 iota中断 iota+=1
		e        //iota+=1
		f = 100  //
		g
		h = iota //恢复计数
		i
	)
	const j = iota
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(j)
}
