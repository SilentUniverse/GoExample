package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)
	//non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}