package asm

import "testing"

// Golang 有自己的汇编语法，与 Intel 汇编类似
// 本例编写了简单的汇编代码，并在 Go 语言中调用
// Go SDK 中也存在汇编逻辑，以 .s 结尾
// 为了解程序执行细节，在 go build 或 go test 等命令执行时，添加 "-gcflags -S" 获得 Go 语言的汇编结果

// 使用 Go 汇编求负数
func neg(x uint64) int8

func TestASM(t *testing.T) {
	println(neg(42))
}