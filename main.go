package main

func main() {
	var funX func(int) int = func(i int) int {
		return i + 1
	}
	var funY = func(fun func(int) int, i int) {
		println(fun(i + 1))
	}

	funY(funX, 2)
}