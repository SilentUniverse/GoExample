package main

import "fmt"

type rect struct {
    width, height int
}

//Methods can be defined for either pointer or value receiver types.
//Here’s an example of a value receiver.
//使用指针接收器（如*rect）的主要原因有两个：
/*避免复制：当结构体较大时，值接收器会在方法调用时都复制整个结构体，会影响性能。指针接收器只传递指针，避免这种开销。
允许修改：如果方法需要修改接收的结构体的内容（例如，改变其属性），则必须使用指针接收器。
值接收器会创建接收结构体的副本，因此对副本的修改不会影响原始结构体。
*/
func (r *rect) area() int {
	r.height = 100
    fmt.Println(r.height)
	r.width = 100
    fmt.Println(r.width)
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}