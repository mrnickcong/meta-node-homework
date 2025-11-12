package main

import "fmt"

/**

2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
    1. 考察点 ：通道的缓冲机制。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}
func main() {
	employee := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		EmployeeID: 12345,
	}

	employee.PrintInfo()

}
