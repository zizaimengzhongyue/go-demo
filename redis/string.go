package main

import (
	"fmt"
	"time"
)

func String() {
	err := rdb.Set(ctx, "key", "value", S).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	res := rdb.SetNX(ctx, "key", "value", S).Val()
	if !res {
		fmt.Println("setnx failed")
	}

	_ = rdb.SetEX(ctx, "key", "value1", S)
	val, err = rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	for i := 0; ; i++ {
		res := rdb.SetNX(ctx, "key", "value2", S).Val()
		if !res {
			fmt.Printf("第 %d 次写入失败\n", i)
		} else {
			fmt.Printf("第 %d 次写入成功\n", i)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
