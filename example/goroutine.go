// goroutine
package main

// goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

// goroutine 语法格式：
// go 函数名( 参数列表 )
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
	// runtime.GOMAXPROCS(1)
	// fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	// sync.WaitGroup其实是一个计数的信号量，
	// 使用它的目的是要main函数等待两个goroutine执行完成后再结束，
	// 不然这两个goroutine还在运行的时候，程序就结束了，看不到想要的结果。
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}
