package main

import (
    "errors"
    "fmt"
)
//自定义错误类型
type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

//errors.As 是 errors.Is 的更高级版本。它检查给定的错误（或其链中的任何错误）是否匹配特定的错误类型，并将其转换为该类型的值，返回 true。如果没有匹配，则返回 false。
func main() {

    _, err := f(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}
/*42
can't work with it
*/