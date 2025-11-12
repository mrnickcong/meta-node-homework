package main

import (
	"fmt"
	"math"
)

/*
*
1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
 1. 考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rectangle := Rectangle{width: 5, height: 3}
	fmt.Println("矩形的面积和周长：", rectangle.Area(), rectangle.Perimeter())

	circle := Circle{radius: 4}
	fmt.Println("圆形的面积和周长：", circle.Area(), circle.Perimeter())

}
