package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*
2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

 1. 考察点 ：原子操作、并发数据安全。
*/
func increment(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func main() {
	var counter int64
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Println(counter)
}

/**
关键修改点
添加 sync.WaitGroup：用于等待所有协程执行完成
修改 increment 函数签名：增加 wg *sync.WaitGroup 参数
在 increment 函数开始处添加 defer wg.Done()：确保协程结束时通知 WaitGroup
在 main 函数中调用 wg.Wait()：等待所有协程完成后再打印结果
这样修改后，程序会正确输出 10000（10个协程 × 1000次递增）
*/
