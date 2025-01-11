package main

import "fmt"

type person struct {
    name string
    age  int
}

//newPerson constructs a new person with the given name.封装了一个构造函数
func newPerson(name string) *person {
	//age 0 可以不写
    //p := person{name: name, age: 0}
	p := person{name: name}
    p.age = 42

    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    //Omitted fields will be zero-valued.
    //省略的字段默认为0
    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    //匿名结构体
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
/*
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
*/