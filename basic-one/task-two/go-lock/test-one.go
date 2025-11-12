package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
    1. 考察点 ： sync.Mutex 的使用、并发数据安全。
*/

func addTenThousand() {
	var count int
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}

func main() {
	addTenThousand()
}
