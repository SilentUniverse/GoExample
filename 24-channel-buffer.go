package main

import "fmt"

func main() {
//默认情况下，通道是非缓冲的，这意味着它们仅在存在相应的接收（<-chan）准备好接收发送的值时才接受发送（chan<-）。缓冲通道接受有限数量的值，而无需相应的接收器来接收这些值。
    messages := make(chan string, 4)

    messages <- "buffered" // 向 channel 发送数据
    messages <- "channel"
	messages <- "is"
	messages <- "cool"

    fmt.Println(<-messages) // 向 channel 接受数据
    fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
buffered
channel
is
cool
*/