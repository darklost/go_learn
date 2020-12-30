// goroutine 资源访问 互斥锁
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

//incCount 不对goroutine 做访问控制
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

//incCountAtomic atomic 原子操作 只支持有数的基础数据类型
func incCountAtomic() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		// value := atomic.LoadInt32(&count)
		runtime.Gosched()
		// value++
		atomic.AddInt32(&count, 1)
		// atomic.StoreInt32(&count, value)
	}
}

// incCountSync Mutex 锁
func incCountSync() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}
func main() {

	count = 0
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println("不加锁", count)

	count = 0
	wg.Add(2)
	go incCountAtomic()
	go incCountAtomic()
	wg.Wait()
	fmt.Println("atomic 加锁", count)

	count = 0
	wg.Add(2)
	go incCountSync()
	go incCountSync()
	wg.Wait()
	fmt.Println("sync 加锁", count)
}
