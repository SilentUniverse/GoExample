package main

import "fmt"

//通过对解引用的指针赋值，可以改变指针所引用地址中的值。这与通过值传递的方式不同，后者只会在函数内部创建一个参数的副本，无法影响原始变量的值。和C是一样的。
func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}