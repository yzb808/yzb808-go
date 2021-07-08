package grammar

import (
	"fmt"
	"reflect"
)

func Call(interface1 Interface1) {
	fmt.Println(reflect.TypeOf(interface1))
	println("line")
	interface1.getName()
}

type Interface1 interface {
	name() string
	getName()
}