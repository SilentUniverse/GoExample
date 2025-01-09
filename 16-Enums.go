package main

import "fmt"

//定义的枚举类型 ServerState 实际上是基于整数类型（int）构建的。这意味着 ServerState 的每个可能值都可以用一个整数来表示。
type ServerState int

const (
    StateIdle ServerState = iota //iota 是一个特殊的预定义标识符，用于在常量声明中生成一系列连续的整数值。每当使用 iota 时，它的值会自动递增，初始值为 0。
    StateConnected
    StateError
    StateRetrying
)
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

//implementing the fmt.Stringer interface
func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}