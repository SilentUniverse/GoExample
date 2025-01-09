package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base
    str string
}


//在 Go 语言中，嵌入（embedding）允许一个结构体（如容器）包含另一个结构体（如基类）。这意味着容器不仅可以拥有自己的字段和方法，还可以直接使用基类的字段和方法。
type base2 struct {
    num string
}

func (b base2) describe2() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container2 struct {
    base2
    num int
}

func main() {

    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)

    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }

    var d describer = co
    fmt.Println("describer:", d.describe())

	co2 := container2{
        base2: base2{
            num: "1",
        },
        num: 1,
    }

	fmt.Printf("co2={num: %v, str: %v}\n", co2.num, co2.num)
    fmt.Println("also num:", co2.base2.num)
}