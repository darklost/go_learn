//函数 方法 和 接口
package main

import "fmt"

// 函数是指不属于任何结构体、类型的方法，也就是说，函数是没有所有者的
//模板
// func 函数名(参数1，……)(返回值1，……){
// 	doSomething
// }

//PrintInt 例子
func PrintInt(a int) int {
	fmt.Println(a)
	return 0
}

//方法是有所有者的，所有者 可以是 struct，interface，甚至我们可以重定义基本数据类型
//模板
// func (所有者名 类型) 方法名(参数列表) 返回值列表 {
// 	doSomething
// }

//Animal 构造Animal结构体，即所有者类型
type Animal struct {
	Name string
}

//Run 这里是值接收，
func (an Animal) Run() {
	fmt.Println("Run ", an.Name)
}

//setName 使用指针接收
func (an *Animal) setName(name string) {
	an.Name = name
}

// 接口(interface)
// 和int，string等一样，接口也是一种类型，和struct类似；区别是struct中存放各种属性，而接口中存放各种方法的声明。另外有几点需要注意：
// 1、接口内的方法数可以为0，即空接口；默认所有对象都实现了空接口
// 2、同一个接口内所有的方法都被实现后，该接口才能被正常使用；例如下方的run和eat方法都需要被实现
// 3、建议将相同的行为放在同一个接口内，例如下方的Eat接口，通过绑定到不同的对象上实现多态
/*
type Eat interface{
    eat() int
}
*/

/*Action 接口是一种类型，和struct类似，
包含n个方法,其中n>=0*/
type Action interface {
	//模板 方法名() 返回类型
	run() int
	eat() int
}

//NullInterface 空接口
type NullInterface interface{}

//Human 接口
type Human struct {
	weight float32 //属性
	name   string
}

func (h Human) run() { //实现方法
	fmt.Println("run run run")
}

func (h Human) eat() { //实现方法
	fmt.Println("eat eat eat")
}

//print 可变参数
func print(a ...interface{}) {
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()
}
func main() {

	fmt.Println("函数和方法")

	//函数调用
	PrintInt(2)

	//方法调用
	var an Animal //声明所有者
	an.Name = "Pig"
	an.Run() //所有者调用方法
	an.setName("pig 1")
	an.Run()

	//接口
	//使用方法
	/*空接口可以接受任意类型，故可以赋值
	  string，struct等其他类型*/
	var n NullInterface = 1
	fmt.Println("null interface", n)

	/*理解接口：接口类似工厂，我们希望生成一个
	  工具箱（Human），它的相关参数是120.0和
	  Bob；我们将生产需求下单给工厂Action，工
	  厂Action生产了一个名字为act的工具箱，里
	  面包含run和eat两种工具*/

	act1 := Human{120.0, "Bob"}
	fmt.Println("hello", act1)
	act1.run()
	act1.eat()

	print(1, 2, 3, n, act1)
}
