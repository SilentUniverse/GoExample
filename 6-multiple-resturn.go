package main

import "fmt"

func vals(a int, b int) (int, int, int) {
    return a, b, a+b
}

func main() {

    a, b, c := vals(1, 2)
    fmt.Println(a)
    fmt.Println(b)
	fmt.Println(c)

    _, c, d := vals(3, 4)
    fmt.Println(c)
	fmt.Println(d)
}