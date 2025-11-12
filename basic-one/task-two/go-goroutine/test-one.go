package main

import (
	"fmt"
	"time"
)

/**
1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
   1. 考察点 ： go 关键字的使用、协程的并发执行。
*/

func printOddNumbers() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
func printEvenNumbers() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	go printOddNumbers()
	go printEvenNumbers()

	time.Sleep(1 * time.Second)

}
