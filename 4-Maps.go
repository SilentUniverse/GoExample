package main

import (
    "fmt"
    "maps"
)

func main() {

    //make(map[key-type]val-type)
    m := make(map[string]int)

    m["k1"] = 4
    m["k2"] = 14

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    clear(m)
    fmt.Println("map:", m)

    //_ 是一个空白标识符，用于忽略返回的第一个值（通常是一个布尔值，表示键是否存在）;prs 是一个变量，用于存储键 "k2" 对应的值。
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    //中括号不能写到下一行
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}

/*
map: map[k1:4 k2:14]
v1: 4
v3: 0
len: 2
map: map[k1:4]
map: map[]
prs: false
map: map[bar:2 foo:1]
n == n2
*/