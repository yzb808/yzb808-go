package grammar

import (
	"testing"
)

func TestClosure1(t *testing.T) {

}

type ClosureStruct struct {
	func1 func(int) int
	value int
}
