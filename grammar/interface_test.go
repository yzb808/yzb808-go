package grammar

import (
	"fmt"
	"testing"
	"unsafe"
)

// 接口简单用例
// 使用接口指向不同结构体，执行对应结构体的方法，多态
func TestInterface(t *testing.T) {
	var p Phone = new(IPhone)
	p.call()
	p = new(NokiaPhone)
	p.call()
}

// struct 的方法既可以使用值传递，也可以使用引用传递
func TestInterface1(t *testing.T) {
	var p1 IPhone = IPhone{callTimes:0}
	p1.call()
	var p2 *IPhone = &IPhone{callTimes:0}
	p2.call()
}

// 接口在 golang 中使用 iface 实现
// iface 由 itab 和 data 组成，itab 描述了接口和引用结构体类型，data 指向 struct 实体
// itab 实体在汇编层面声明，接口每指向一种具体类型，都会生成一个 itab 实体，因此一个接口会对应多个 itab 对象
// 结构体向接口赋值的语句，在汇编层面转化成对具体 itab 实体的操作
// 例如：
//     LEAQ	go.itab.*"".NokiaPhone,"".Phone(SB), CX
//     MOVQ	CX, "".p+96(SP)
func TestInterface2(t *testing.T) {
	pi := new(IPhone)
	var p Phone = pi
	p.call()
	fmt.Println(unsafe.Pointer(pi))
	fmt.Println(((*iface)(unsafe.Pointer(&p)).data))
}

// 两个接口绑定到同一个 struct 实体，实际调用的是一个 call 方法
// 例如：
//     LEAQ	go.itab.*"".IPhone,"".Phone(SB), AX
//     MOVQ	24(AX), AX
//     CALL	AX
//     LEAQ	go.itab.*"".IPhone,"".NoPhone(SB), AX
//     MOVQ	24(AX), AX
//     CALL	AX
//     go.itab.*"".IPhone,"".Phone 和 go.itab.*"".IPhone,"".NoPhone 的第 24 位都存放着 "".(*IPhone).call
// 其中 24(AX) 是 call 方法函数句柄
//     rel 24+8 t=1 "".(*NokiaPhone).call+0
func TestInterface4(t *testing.T) {
	p := IPhone{callTimes:1}
	var phone Phone = &p;
	phone.call()
	var noPhone NoPhone = &p
	noPhone.call()
}


type Phone interface {
	call()
}

type NoPhone interface {
	call()
}

type NokiaPhone struct {
	field1 int32
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("nokia call")
}

// IPhone 结构体包含 callTimes 成员变量，记录调用 call 函数次数。这种设计用于测试触发方法的结构体，值传递和引用传递的区别
// IPhone 使用值传递
type IPhone struct {
	callTimes int64
}

func (iPhone *IPhone) call() {
	iPhone.callTimes++
	fmt.Println("iphone call", iPhone.callTimes)
}

// IPhone1 使用引用传递
type IPhone1 struct {
	callTimes int64
}

func (iPhone *IPhone1) call() {
	iPhone.callTimes++
	fmt.Println("iphone call")
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter *uint64
	_type *uint64
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}