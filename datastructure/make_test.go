package datastructure

import (
	"fmt"
	"testing"
)

// golang make 函数用于创建和初始化内置数据结构
// 具体是指切片、哈希表和 Channel

func TestMake(t *testing.T) {
	slice := make([]int, 0, 100)
	fmt.Println(len(slice))
	//
	//hash := make(map[int]bool, 10)
	//ch := make(chan int, 5)
}
