package main

import (
    "fmt"
    "math"
)

//Interfaces are named collections of method signatures
type geometry interface {
    area() float64
    perim() float64 //周长
}

type compute interface {
	abc() float64
	def() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (r rect) abc() float64 {
    return r.width * r.height
}
func (r rect) def() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func (c circle) abc() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) def() float64 {
    return 2 * math.Pi * c.radius
}
//可以将接口类型用作函数参数，接受任何实现了该接口的类型
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func measure2(c compute) {
    fmt.Println("measure2", c)
    fmt.Println(c.abc())
    fmt.Println(c.def())
}
//if ... ok { 检查转换是否成功，只有在成功时才执行代码块
/*  c, ok := g.(circle) 这是类型断言的安全形式：
	c 将会是转换后的 circle 类型的值
	ok 是一个布尔值，表示转换是否成功
	如果 g 确实是一个 circle，则 ok 为 true，c 包含实际的 circle 值
	如果 g 不是 circle（比如是 rect），则 ok 为 false
*/
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
	/*    
	switch v := g.(type) {
    case circle:
        fmt.Printf("这是一个圆，半径为 %.2f\n", v.radius)
    case rect:
        fmt.Printf("这是一个矩形，宽度为 %.2f，高度为 %.2f\n", v.width, v.height)
    default:
        fmt.Println("未知的几何形状")
    }
	*/
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

	measure2(r)
    measure2(c)

    detectCircle(r)
    detectCircle(c)
}