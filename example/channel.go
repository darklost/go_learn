// channel 通道
package main

import "fmt"

func main() {

	//无缓冲通道 （同步通道）
	ch := make(chan int)

	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()

	fmt.Println(<-ch)

	//管道

	one := make(chan int)
	two := make(chan int)

	go func() {
		fmt.Println("写入管道 one")
		one <- 100
	}()

	go func() {
		v := <-one
		fmt.Println("读取管道 one", v)
		fmt.Println("写入管道 two", v)
		two <- v
	}()

	fmt.Println("读取管道 two", <-two)

	//有缓冲的通道
	ch1 := make(chan int, 3)
	fmt.Println(cap(ch1), len(ch1))

	//单向通道
	//有时候，我们有一些特殊场景，比如限制一个通道只可以接收，但是不能发送；有时候限制一个通道只能发送，但是不能接收，这种通道我们称为单向通道。
	// 定义单向通道也很简单，只需要在定义的时候，带上<-即可。
	// var send chan<- int    //只能发送
	// var receive <-chan int //只能接收
	// 注意<-操作符的为止，在后面是只能发送，对应发送操作；在前面是只能接收，对应接收操作。

	// 单向通道应用于函数或者方法的参数比较多，比如
	// func counter(out chan<- int) {
	// }
}
