package grammar

import "testing"

func TestEmbedded(t *testing.T) {
	v1 := myStructChild{}
	v1.field1 = 1
	v2 := myStruct{2, 3}
	v1.myStruct = v2

	println(v1.field1) // v1.field1
	println(v1.field2) // v2.field2

	// myStructChild 的 fun 函数继承自 v2，实现是把 v2 对象伪装成 v1（拨动指针指向内嵌 struct），随后作为入参调用 myStruct 的 fun 函数
	println(v1.fun()) // 使用的 v2.field1 + v2.field2
}

type myStruct struct {
	field1 int32
	field2 int32
}

func (ms *myStruct) fun() int32 {
	return ms.field1 + ms.field2
}

type myStructChild struct {
	myStruct
	field1 int32
}