/**
 * 编写一个函数，入参为 arr []int, num int, fn func(int) int，需求时使用不超过 num 个 goroutine 对 arr 中每个元素执行 fn，并返回执行后的 arr
 * 这是今天面试的一道题目，面试的时候卡在执行计算的几个 goroutine  怎么退出了
 **/
package main

import (
	"fmt"
	"sync"
)

func concurrent(arr []int, num int, fn func(int) int) []int {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	done := make(chan bool)
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case k := <-ch:
					arr[k] = fn(arr[k])
				case <-done:
					return
				}
			}
		}()
	}
	for i := 0; i < len(arr); i++ {
		ch <- i
	}
	close(done)
	wg.Wait()
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
