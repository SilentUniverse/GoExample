package main

import "fmt"

//  the type of nums is equivalent to []int
//  the ... is the variadic parameter
// 在定义了一个可变参数的函数时，传入的参数会被处理为一个切片（slice），具体来说是一个整数切片（[]int）
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)  // 这将把切片中的每个元素作为单独的参数传递给 sum 函数
}