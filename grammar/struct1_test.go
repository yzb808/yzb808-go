package grammar

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type MyStruct1 struct {
	field1 int64
}

func (ms *MyStruct1) name() string {
	return "MyStruct1"
}

func (ms *MyStruct1) getName() {
	fmt.Println(unsafe.Pointer(ms))
	ms.field1 = 123456
	fmt.Println(reflect.TypeOf(ms))
	println(ms.name())
}

type MyStruct2 struct {
	field2 int32
	MyStruct1
}

func (ms MyStruct2) name() string {
	return "MyStruct2"
}

//func (ms *MyStruct2) getName() {
//	//fmt.Println(reflect.TypeOf(ms))
//	//var int1 interface1 = ms
//	//println(int1.name())
//	print(ms.name())
//}

func TestStruct1(t *testing.T) {
	var s MyStruct2 = MyStruct2{field2:1234567}
	s.field1 = 7654321
	fmt.Println(unsafe.Pointer(&s.field1), unsafe.Pointer(&s.field2))
	fmt.Println(s.field2)
	s.name()
}

