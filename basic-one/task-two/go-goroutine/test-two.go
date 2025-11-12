package main

import (
	"fmt"
	"time"
)

/**
2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
   1. 考察点 ：协程原理、并发任务调度。
*/

func taskScheduler(tasks []func()) {
	for _, task := range tasks {
		go func() {
			start := time.Now()
			task()
			end := time.Now()
			fmt.Println("任务执行完成，耗时：", end.Sub(start))
		}()
	}
}

func main() {
	tasks := []func(){}

	func1 := func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务1完成")
	}

	tasks = append(tasks, func1)

	taskScheduler(tasks)

	time.Sleep(2 * time.Second)

}
