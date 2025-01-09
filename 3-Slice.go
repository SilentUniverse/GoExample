package main

import (
    "fmt"
    "slices"
)

func main() {

    // An uninitialized slice equals to nil and has length 0.声明一个未初始化的切片，等于nil，长度为0
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

	a := [...]string{"a", "b", "c"}
	fmt.Println("a:", a)

    // 创建一个长度为3的切片，初始值为空字符串
    // 使用make创建切片，传入一个capacity参数，表示切片的最大容量
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    // 设置切片元素，len拿到长度
    // 使用append向切片添加元素
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // 创建一个长度为len(s)的切片，并复制s的元素
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    //切片
    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    // 创建一个包含"g", "h", "i"的切片
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // 比较两个切片是否相等Equal
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    // 创建一个包含3个空切片的切片
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}