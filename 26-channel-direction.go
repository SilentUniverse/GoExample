package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
/*
发送操作总是将值推向 channel（ch <- value）
接收操作总是从 channel 拉取值（value := <-ch）
箭头 <- 的方向指示了数据的流动方向
*/