package fun1

import "fmt"

func init() {
	fmt.Println("Fun1 init")
}

func Fun1() {
	fmt.Println("print fun1")
}

type Phone interface {
	Call()
	//play()
}