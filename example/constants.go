//定义常量
package main

import "fmt"

const (
	a = 1
	b
	c
	d
)
const (
	i = iota
	j = iota
	x = iota
)
const xx = iota
const yy = iota

const (
	ii = 1 << iota
	jj = 3 << iota
	kk
	ll
)

// iota 表示从 0 开始自动加 1，所以 ii=1<<0, jj=3<<1（<< 表示左移的意思），即：ii=1, jj=6，这没问题，关键在 kk 和 l，从输出结果看 kk=3<<2，ll=3<<3。

// 简单表述:

// ii=1：左移 0 位,不变仍为 1;
// jj=3：左移 1 位,变为二进制 110, 即 6;
// kk=3：左移 2 位,变为二进制 1100, 即 12;
// ll=3：左移 3 位,变为二进制 11000,即 24。
// 注：<<n==*(2^n)。

func main() {
	fmt.Println(a)
	// b、c、d没有初始化，使用上一行(即a)的值
	fmt.Println(b) // 输出1
	fmt.Println(c) // 输出1
	fmt.Println(d) // 输出1

	//iota
	// 	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	// iota 可以被用作枚举值：
	// const (
	// 	a = iota
	// 	b = iota
	// 	c = iota
	// )
	//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	// const (
	// 	a = iota
	// 	b
	// 	c
	// )
	fmt.Println(i, j, x, xx, yy)
	// 输出是 0 1 2 0 0

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)

	fmt.Println("ii=", ii)
	fmt.Println("jj=", jj)
	fmt.Println("kk=", kk)
	fmt.Println("ll=", ll)
}
