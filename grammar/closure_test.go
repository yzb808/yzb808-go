package grammar

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// 闭包，实际返回的是结构体，封装了匿名函数和使用到的成员变量。
// 闭包中的匿名函数，使用闭包结构体寻址依赖到的局部变量
func adder() func(int) string {
	sum := 0
	return func(x int) string {
		sum += x
		return strconv.Itoa(sum)
	}
}

func adder1() func(int) string {
	y := 99
	y1 := 100
	var func1 = func(x int) string {
		y += x
		y += y1
		return strconv.Itoa(y)
	}
	return func1
}

// 匿名函数引用父函数局部变量，是通过在 DX 寄存器中记录父函数栈帧里局部变量地址，
// 利用 DX 将地址传递给匿名函数，从而被访问
func adder2() int {
	sum := 123
	sum1 := 124
	sum2 := 125
	sum2++
	sum3 := 126
	sum4 := 127
	func1 := func(x int) int {
		sum = sum1 + sum3 + sum4
		sum1 = 11
		sum += x
		return sum
	}
	func1(11)
	fmt.Println(sum1)
	return 1
}

func TestClosure(t *testing.T) {
	adder2()
	myAdder := adder()
	println(reflect.TypeOf(myAdder).String())

	// 从1加到10
	for i := 1; i <= 10; i++ {
		myAdder(i)
	}

	fmt.Println(myAdder(0))

	myAdder1 := adder()
	// 再加上45
	fmt.Println(myAdder1(45))

}