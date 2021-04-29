// 面试的时候跟面试官聊到一个数据同步的问题，当时提到可以通过监听 mysql 的 binlog 实现，但是具体该怎么做不了解；
// 查了下有现成的工具可以做，写个 demo
package main

import (
	"github.com/go-mysql-org/go-mysql/canal"

	"github.com/zizaimengzhongyue/go-demo/go-mysql/config"
)

func main() {
	root := config.RootDir()
	cfg := canal.NewDefaultConfig()
	if err := config.Load(&cfg, root+"/cfg/db.toml"); err != nil {
		panic(err)
	}

	can, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	can.SetEventHandler(&Bootstrap{})

	can.Run()
}
