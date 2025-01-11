package main

import (
    "fmt"
    "time"
)

func main() {
    //带bffer
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()
    //选择结构会优先处理第一个准备好的接收操作”，这意味着如果 c1 在 1 秒内有值可接收，程序将执行 case res := <-c1 的代码块。如果 c1 在 1 秒内没有值，程序将执行超时的代码块 case <-time.After(1 * time.Second)。
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}