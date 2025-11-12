package main

import "fmt"

func plusOne(digits []int) []int {

	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		digits[i]++

		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	//运行到这里说明是 9999999，结果就是 10000000
	digits = make([]int, n+1)
	digits[0] = 1

	return digits

}

func main() {
	digits := []int{9, 9, 9, 9}
	result := plusOne(digits)
	fmt.Println(result)

}
