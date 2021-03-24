package main

import (
	"fmt"
	"runtime"
	"time"
)

// 在 timeout 中 select 退出之后，remoter 所在的 goroutine 因为无法向 ch 写入数据会导致 goroutine 泄漏
func remoter(ch chan bool) {
	defer func() {
		fmt.Println("remoter finished")
	}()
	time.Sleep(3 * time.Second)
	ch <- true
}

func timeout() {
	ch := make(chan bool)
	defer func() {
		fmt.Println("timeout finished")
	}()
	go remoter(ch)
	select {
	case <-ch:
		return
	case <-time.After(2 * time.Second):
		return
	}
}

func Run() {
	for i := 0; i < 5; i++ {
		go timeout()
	}
}

func main() {
	Run()
	// 随着 timeout 的退出，goroutine 的数量会从 11 个降到 6 个，然后不再下降
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println(runtime.NumGoroutine())
	}
	fmt.Println("finished")
}
