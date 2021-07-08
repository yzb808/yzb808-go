package main

import "testing"

type Phone interface {
	call()
}
// 2. 声明结构体
type IPhone struct {
	f1 int32
}
// 3. 为结构体挂载方法
func (iPhone *IPhone) call() {
}
// 4. 接口指向 struct 实体，并调用挂在 struct 上的方法
func TestInterface(t *testing.T) {
	pi := new(IPhone)
	var p Phone = pi
	p.call()
}
