package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

const ZSetName01 = "zset-name-01"
const ZSetName02 = "zset-name-02"
const ZSetName03 = "zset-name-03"
const ZSetName04 = "zset-name-04"
const ZSetName05 = "zset-name-05"
const ZSetName06 = "zset-name-06"

func ZSetInit() {
	_ = rdb.Del(ctx, ZSetName01)
	_ = rdb.Del(ctx, ZSetName02)
	_ = rdb.Del(ctx, ZSetName03)
	_ = rdb.Del(ctx, ZSetName04)
	_ = rdb.Del(ctx, ZSetName05)
	_ = rdb.Del(ctx, ZSetName06)

	res := rdb.ZAdd(ctx, ZSetName01, &redis.Z{Score: 1.0, Member: "zset-key-1"}, &redis.Z{Score: 2.0, Member: "zset-key-2"}, &redis.Z{Score: 3.0, Member: "zset-key-3"})
	fmt.Println(res.String())
	ZScan(ZSetName01)

	res = rdb.ZAdd(ctx, ZSetName02, &redis.Z{Score: 1.0, Member: "zset-key-1"}, &redis.Z{Score: 2.0, Member: "zset-key-2"}, &redis.Z{Score: 3.0, Member: "zset-key-3"})
	fmt.Println(res.String())
	ZScan(ZSetName02)

	_ = rdb.ZAdd(ctx, ZSetName03, &redis.Z{Score: 1.0, Member: "zset-key-1"}, &redis.Z{Score: 2.0, Member: "zset-key-2"}, &redis.Z{Score: 3.0, Member: "zset-key-3"})
	_ = rdb.ZAdd(ctx, ZSetName04, &redis.Z{Score: 0.5, Member: "zset-key-1"}, &redis.Z{Score: 1.5, Member: "zset-key-5"}, &redis.Z{Score: 2.5, Member: "zset-key-6"})
}

func ZScan(name string) {
	var cursor uint64
	var strs []string
	for {
		res := rdb.ZScan(ctx, name, cursor, "*", 1)
		strs, cursor, _ = res.Result()
		fmt.Println(strs)
		if cursor == 0 {
			break
		}
	}
}

func ZCard() {
	res := rdb.ZCard(ctx, ZSetName01)
	r, _ := res.Result()
	fmt.Println(r)

	res = rdb.ZCard(ctx, ZSetName02)
	r, _ = res.Result()
	fmt.Println(r)
}

func ZCount() {
	res := rdb.ZCount(ctx, ZSetName01, "1", "2")
	r, _ := res.Result()
	fmt.Println(r)
}

func ZIncrBy() {
	res := rdb.ZIncrBy(ctx, ZSetName01, 5, "zset-key-1")
	r, _ := res.Result()
	fmt.Println(r)

	res = rdb.ZScore(ctx, ZSetName01, "zset-key-1")
	r, _ = res.Result()
	fmt.Println(r)
}

func ZRange() {
	res := rdb.ZRange(ctx, ZSetName01, 0, -1)
	strs, _ := res.Result()
	fmt.Println(strs)

	res = rdb.ZRange(ctx, ZSetName01, 0, 1)
	strs, _ = res.Result()
	fmt.Println(strs)
}

func ZRangeByScore() {
	res := rdb.ZRangeByScore(ctx, ZSetName01, &redis.ZRangeBy{Min: "0", Max: "2"})
	strs, _ := res.Result()
	fmt.Println(strs)
}

func ZRank() {
	res := rdb.ZRank(ctx, ZSetName01, "zset-key-1")
	r, _ := res.Result()
	fmt.Println(r)
}

func ZRem() {
	res := rdb.ZRem(ctx, ZSetName01, "zset-key-1")
	r, _ := res.Result()
	fmt.Println(r)

	ZScan(ZSetName01)
}

func ZRemRangeByRank() {
	res := rdb.ZRemRangeByRank(ctx, ZSetName01, 0, -1)
	r, _ := res.Result()
	fmt.Println(r)

	ZScan(ZSetName01)
}

func ZRemRangeByScore() {
	res := rdb.ZRemRangeByScore(ctx, ZSetName02, "0", "2")
	r, _ := res.Result()
	fmt.Println(r)

	ZScan(ZSetName02)
}

func ZRevRange() {
	res := rdb.ZRevRange(ctx, ZSetName02, 0, 2)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func ZRevRangeByScore() {
	res := rdb.ZRevRangeByScore(ctx, ZSetName02, &redis.ZRangeBy{Min: "0", Max: "10"})
	strs, _ := res.Result()
	fmt.Println(strs)
}

func ZRevRank() {
	res := rdb.ZRevRank(ctx, ZSetName02, "zset-key-3")
	fmt.Println(res.String())
}

func ZScore() {
	res := rdb.ZScore(ctx, ZSetName02, "zset-key-3")
	val, _ := res.Result()
	fmt.Println(val)
}

func ZUnionStore() {
	res := rdb.ZUnionStore(ctx, ZSetName05, &redis.ZStore{Keys: []string{ZSetName03, ZSetName04}})
	fmt.Println(res)

	ZScan(ZSetName05)
}

func ZInterStore() {
	res := rdb.ZInterStore(ctx, ZSetName06, &redis.ZStore{Keys: []string{ZSetName03, ZSetName04}})
	fmt.Println(res)

	ZScan(ZSetName06)
}

func ZSet() {
	ZSetInit()
	ZCard()
	ZCount()
	ZIncrBy()
	ZRange()
	ZRangeByScore()
	ZRank()
	ZRem()
	ZRemRangeByRank()
	ZRemRangeByScore()
	ZRevRange()
	ZRevRangeByScore()
	ZRevRank()
	ZScore()
	ZUnionStore()
	ZInterStore()
}
