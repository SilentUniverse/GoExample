package main

import (
    "fmt"
    "net/http"
    "time"
)
//每当服务器接收到一个HTTP请求时，Go的net/http包会自动为该请求生成一个上下文对象。这个上下文对象用于在处理请求时传递信息，比如取消信号和截止日期。
func hello(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}