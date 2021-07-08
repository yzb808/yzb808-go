package fun2

import "fmt"

func Fun2() {
	fmt.Println("print fun2")
}

type Phone interface {
	Call()
	//play()
}