package grammar

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type MyStruct struct {
	field1 int32
	field2 int64
	field3 int64
	field4 int64
	//field5 int64
	//field6 int64
	//field7 int64
	//field8 int64
	//field9 int64
}

func (myStruct *MyStruct) structFunc() {
	myStruct.field4 = myStruct.field2 + myStruct.field3
}

func normalFunc(myStruct MyStruct)  {
	myStruct.field4 = myStruct.field2 + myStruct.field3
}

func TestStruct(t *testing.T) {
	s := MyStruct{
		field1: 99,
		//field2: 108,
	}
	fmt.Println(*(*[100]byte)(unsafe.Pointer(&s)))
}

func TestStructFunc(t *testing.T) {
	sss := new(MyStruct)
	normalFunc(*sss)
	sss.structFunc()
	fmt.Println(sss)
}

func TestStructOffset(t *testing.T) {
	typ := reflect.TypeOf(MyStruct{})
	fmt.Println(typ.Size())
	num := typ.NumField()
	for i := 0; i < num; i++  {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n", field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}
}