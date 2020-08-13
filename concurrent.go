/**
 * 编写一个函数，入参为 arr []int, num int, fn func(int) int，需求时使用不超过 num 个 goroutine 对 arr 中每个元素执行 fn，并返回执行后的 arr
 * 这是今天面试的一道题目，面试的时候卡在执行计算的几个 goroutine 了
 **/
package main

import (
	"fmt"
)

func concurrent(arr []int, num int, fn func(int) int) []int {
	ch := make(chan int)
	done := make(chan bool)
	ok := make(chan bool)
	go func() {
		for i, _ := range arr {
			ch <- i
		}
		done <- true
	}()
	for i := 0; i < num; i++ {
		go func(ch chan int) {
			for {
				select {
				case <-ok:
					return
				case x := <-ch:
					arr[x] = fn(arr[x])
					break
				}
			}
		}(ch)
	}
	<-done
	close(ok)
	return arr
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	num := 3
	a := func(x int) int {
		return x * 2
	}
	fmt.Println(concurrent(nums, num, a))
}
