package grammar

import (
	"fmt"
	"testing"
	"unsafe"
)

// golang 赋值操作
// 基础数据类型是值拷贝，java 语言基于不变设计模式，也是值拷贝
// 非基础数据类型的赋值操作，go 语言依然采用值拷贝，而 java 使用引用拷贝
// 值拷贝使得两个变量能解耦，而代价是要完成数据复制，在复杂结构体或频繁赋值的场景下，会有较大性能开销
// go 语言中要显式使用指针来实现引用拷贝
func TestAssign(t *testing.T) {
	// 基础数据类型，值拷贝
	x := 1
	y := x
	fmt.Println(unsafe.Pointer(&x), unsafe.Pointer(&y)) // 0xc000110278 0xc000110280
	y = y + 1
	fmt.Println(x, y)

	// 结构体，值拷贝
	var s1 = structX{m: 1}
	var s2 = s1
	fmt.Println(unsafe.Pointer(&s1), unsafe.Pointer(&s2))
	s2.m = 2
	fmt.Println(s1.m, s2.m)
	// 结构体，使用指针，引用拷贝
	// 在 golang 中，'->' 和 '.' 都用 '.' 表示
	var s3 = &s1
	s3.m = 3
	fmt.Println(s1.m, s3.m)
}

// 测试函数入参和返回值的传递特性
// 默认是值传递，引用传递要显示使用地址
func TestReturn(t *testing.T) {
	var s structX
	var sx = structX{m: 1} // 变量 sx 作用域仅限当前方法，存储在栈
	s.m = sx.m
	fmt.Println("入参", unsafe.Pointer(&s)) // 变量 s 逃逸出当前方法，存储在 heap
	s1, s11 := fun1(s, &s);
	fmt.Println("返回值out", unsafe.Pointer(&s1))
	fmt.Println("返回值copy", unsafe.Pointer(s11))
}

func fun1(s structX, ss *structX) (structX, *structX) {
	fmt.Println("入参copy", unsafe.Pointer(&s))
	fmt.Println("入参in", unsafe.Pointer(ss))
	var s1 structX;
	fmt.Println("返回值", unsafe.Pointer(&s1))
	return s1, &s1;
}

type structX struct {
	m int32
	n int64
}
