package main

import (
	"fmt"
	hel "hello/examples"
)

func main() {
	fmt.Printf("hello, world\n")
	var1 := "girish"
	ownFuncaton(var1)

	hel.Variables2()

}

func main2() {
	var a = 10
	var b = "girish"
	c := 10.5

	fmt.Println(a, b, c)

}
func ownFuncaton(str string) {
	hel.Variables()
	fmt.Printf(str)
}
