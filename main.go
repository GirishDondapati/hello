package main

import (
	"fmt"
	hel "hello/examples"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println("argsWithProg ", argsWithProg)
	fmt.Println("argsWithoutProg ", argsWithoutProg)
	fmt.Printf("hello, world\n")
	var1 := "girish"
	ownFuncaton(var1)

	hel.Variables2()
	fmt.Printf("Constants example started")
	hel.ConstantsEx()

	fmt.Printf("For loop example started")
	hel.ForLoopEx()

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
