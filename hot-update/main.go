/**
 * 不重启服务更新服务配置
 * 使用 kill -10 给进程发送信号在 mac 环境似乎有 bug，linux 运行正常
 */
package main

import (
	"fmt"
	"time"
)

func Run() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(GetConfig())
	}
}

func main() {
	Run()
}
