//定义变量
package main

import "fmt"

var xx, yy int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	aa int
	bb bool
)

var cc, dd int = 1, 2
var ee, ff = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//gg, hh := 123, "hello"

func main() {
	//变量

	//1 指定变量类型，如果没有初始化，则变量默认为零值。
	//var identifier type
	var a string = "test"
	fmt.Println(a)
	//定义多个变量
	// var identifier1, identifier2 type
	var b, c int = 1, 2
	fmt.Println(b, c)
	//未初始化默认值
	// 数值类型（包括complex64/128）为 0
	// 布尔类型为 false
	// 字符串为 ""（空字符串）
	// 以下几种类型为 nil：
	// var a *int
	// var a []int
	// var a map[string]int
	// var a chan int
	// var a func(string) int
	// var a error // error 是接口
	var iVal int
	var fVal float64
	var bVal bool
	var sVal string
	fmt.Printf("%v %v %v %q\n", iVal, fVal, bVal, sVal)

	//2 定义变量自动推导
	//var v_name v_type
	// v_name = value
	var d = true
	fmt.Println(d)

	//3 省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	//v_name := value
	var intVal int

	// intVal := 1 // 这时候会产生编译错误

	intVal, intVal1 := 1, 2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Println(intVal, intVal1)

	f := "test2" // var f string = "test2"

	fmt.Println(f)

	//类型相同多个变量, 非全局变量
	// var vname1, vname2, vname3 type
	// vname1, vname2, vname3 = v1, v2, v3

	// var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

	// vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

	// // 这种因式分解关键字的写法一般用于声明全局变量
	// var (
	// 	vname1 v_type1
	// 	vname2 v_type2
	// )

	g, h := 123, "hello"
	fmt.Println("test", xx, yy, aa, bb, cc, dd, ee, ff, g, h)

}
