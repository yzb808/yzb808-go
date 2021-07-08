package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestReflect(t *testing.T) {
	str := "hello"
	fmt.Println("TypeOf author:", reflect.TypeOf(str))
	fmt.Println("ValueOf author:", reflect.ValueOf(str))

	sx1 := structX{1, 2}
	sx2 := structX{1, 2}
	fmt.Println(reflect.DeepEqual(sx1, sx2))
}

func TestInterface(t *testing.T) {
	str := "hello"
	data := (*[32]byte) (unsafe.Pointer(&str))
	fmt.Println(data)
	sx := structX{
		f1: 111111,
		f2: 222,
	}
	data1 := (*[1280]byte) (unsafe.Pointer(&sx))
	fmt.Println(data1)
}

type structX struct {
	f1 int32
	f2 int64
}
