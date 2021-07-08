package grammar

import (
	"testing"
	"time"
)

// 协程饥饿
// 启动 10 个不会结束执行，且不调用其他函数的 go 任务，用于耗尽 MPG 中的 M
// 接着启动一个简单 go 任务，发现在 go 1.15 时，简单任务不会饥饿，go 1.12 和 go 1.14 会饥饿
// 因为 go 语言的协程策略依赖协程自觉，但为了避免某些协程不自觉，造成别的协程拿不到 cpu，因此引入抢占机制。
// 12 版本中，协程判断自身是否被抢占，通过函数调用时由编译器插入的抢占检查，本例协程不调用函数。
// 14 版本在 STW 和 栈扫描时附着抢占判断，本例程也未生效。
// 15 版本为何不会饥饿，目前还不清楚
func TestHungry(t *testing.T) {
	for i := 0; i < 6; i++ {
		go func() {
			for {
				j := 1
				j++
			}
		}()
	}

	time.Sleep(time.Second * 1)
	go println(111)
	time.Sleep(time.Hour * 1)
}
