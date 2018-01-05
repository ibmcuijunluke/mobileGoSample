package main

import (
	"MergeSort"
	"container/list"
	"fmt"
)

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	fl := list.New()
	fl.PushBack(int(0))
	fl.PushBack(int(0))
	return func() int {
		a := fl.Front()
		fl.Remove(a)
		b := fl.Front()
		tmp := int(a.Value.(int)) + int(b.Value.(int)) //类型断言
		if tmp == 0 {
			tmp = 1
		}
		fl.PushBack(tmp)
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	var arr = []int{6, 5, 3, 1, 8, 7, 2, 4, 9}
	fmt.Println(arr)
	fmt.Println(MergeSort(arr))
}
