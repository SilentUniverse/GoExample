package main

import (
    "fmt"
    "time"
)

//定时器

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped") 
    }

    time.Sleep(2 * time.Second)
}

//timer1.C 中的 C 是 time.Timer 结构体中的一个字段。time.Timer 结构体有一个名为 C 的字段，它是一个通道（channel），用于接收定时器到期时的通知。

//具体来说，当你创建一个新的定时器时，例如 timer1 := time.NewTimer(2 * time.Second)，Go 会在内部创建一个定时器，并在 C 通道中发送一个值，表示定时器已经到期。通过 <-timer1.C 语句，你可以阻塞当前的 goroutine，直到定时器到期并向通道发送一个值。

//因此，C 是 time.Timer 结构体的一个字段，表示定时器的通道。