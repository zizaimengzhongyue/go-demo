package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Config struct {
	Response string `json:"response"`
}

var cfgs = []Config{Config{}, Config{}}
var index = 0
var dir string
var mu = &sync.Mutex{}

func GetConfig() Config {
	return cfgs[index]
}

func Update() error {

	mu.Lock()
	defer mu.Unlock()

	var path string
	if index == 0 {
		path = dir + "/conf01.json"
	} else {
		path = dir + "/conf02.json"
	}
	target := (index + 1) % 2
	if err := load(path, &cfgs[target]); err != nil {
		return err
	}
	index = target

	return nil
}

func load(path string, data interface{}) error {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, data); err != nil {
		return err
	}
	return nil
}

func registerSignal() error {
	ch := make(chan os.Signal, 5)
	signal.Notify(ch, syscall.SIGUSR1)
	go func() {
		for {
			_ = <-ch
			fmt.Println("捕获更新配置信号")
			Update()
		}
	}()
	return nil
}

func initDir() error {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	if err := initDir(); err != nil {
		panic(err)
	}
	if err := registerSignal(); err != nil {
		panic(err)
	}
	if err := Update(); err != nil {
		panic(err)
	}
}
