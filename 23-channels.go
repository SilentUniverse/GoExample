package main

import "fmt"

func main() {

	//create new channel
    messages := make(chan string)

    go func() { messages <- "ping 4.4.4.4" }()

	//By default sends and receives block
    msg := <-messages
    fmt.Println(msg)
}