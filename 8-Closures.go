package main

import "fmt"


//Go supports anonymous functions, which can form closures.
//This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func stringSeq() func() string {
	i := 0
	return func() string {
		i++
		return fmt.Sprintf("Hello %d\n", i)
	}
}

func main() {

    nextInt := intSeq()

	nextIntString := stringSeq()
	fmt.Println(nextIntString())

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}