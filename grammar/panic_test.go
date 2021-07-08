package grammar

import (
	"fmt"
	"testing"
)

// 吃掉本函数的异常
func TestPanic(t *testing.T) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("panic once")
}

// 吃掉下游可能产生的异常
func TestEatPanic(t *testing.T) {
	funPanic1()
}

func funPanic1()  {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	funPanic2()
	fmt.Println("fun1")
}

func funPanic2()  {
	funPanic3()
	fmt.Println("fun2")
}

func funPanic3()  {
	panic("panic from fun3")
	fmt.Println("fun3")
}

// 吃掉异常后，返回指定数值
func TestPanicReturn(t *testing.T) {
	fmt.Println(funPanicReturn())
}

// 通过显示声明返回值名称的方式赋值
func funPanicReturn() (ret string) {
	defer fmt.Println("defer 1")
	defer func() {
		if x := recover(); x != nil {
			ret = "panic recovered"
		}
	}()
	defer fmt.Println("defer 2")
	return "pass"
}

func TestPanicEmbedded(t *testing.T) {
	defer func() {
		if x := recover(); x != nil {
		}
		panic("throw again")
	}()
	panic("throw")
}