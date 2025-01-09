package main

import (
    "errors" //errros包
    "fmt"
)

//  Go 语言中，函数通常会返回两个值：一个是正常的结果，另一个是错误值（如果有的话）
//errors.New constructs a basic error value with the given error message.
func f(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

//哨兵错误是一个预先声明的变量，用于表示特定的错误条件。
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}

func main() {
    for _, i := range []int{7, 42} {

        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil {

            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }

        fmt.Println("Tea is ready!")
    }
}