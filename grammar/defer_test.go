package grammar

import (
	"testing"
)

// defer 修饰的函数，会在执行 defer 的函数返回之前执行
func TestDefer(t *testing.T) {
	println("test defer begin")

	for i := 0; i < 5; i++ {
		defer fun()
	}

	println("test defer end")

}

func fun()  {
	println("defer")
}
