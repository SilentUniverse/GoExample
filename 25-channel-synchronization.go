package main

import (
    "fmt"
    "time"
)


//done是一个变量名

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
	//go 关键字启动新的 goroutine,worker 函数在独立的 goroutine 中异步执行
    go worker(done)

    <-done
}