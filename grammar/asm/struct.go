package asm

type myStruct struct {
	field1 int32
	field2 int32
}

func (*myStruct) fun(in int) {
}
