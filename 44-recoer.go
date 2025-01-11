package main

import "fmt"

func mayPanic() {
    panic("a problem")
}

func main() {

    defer func() {
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
}

//Go 通过使用内置函数 recover 来实现从panic中恢复,如果服务器出现严重错误，但是不希望崩溃