package main

import (
	"fmt"
	"sync"
	"time"
)

const ListName01 = "list-name-01"
const ListName02 = "list-name-02"
const ListName03 = "list-name-03"
const ListName04 = "list-name-04"
const ListName05 = "list-name-05"
const ListName06 = "list-name-06"

func ListInit() {
	_ = rdb.Del(ctx, ListName01)
	_ = rdb.Del(ctx, ListName02)
	_ = rdb.Del(ctx, ListName03)
	_ = rdb.Del(ctx, ListName04)

	res := rdb.LPush(ctx, ListName01, "list-01", "list-02")
	fmt.Println(res.String())

	res = rdb.LPush(ctx, ListName02, "list-01", "list-03", "list-05")
	fmt.Println(res.String())
}

func BLpop() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			res := rdb.BLPop(ctx, 3*time.Second, ListName03, ListName04)
			strs, _ := res.Result()
			fmt.Println(strs)
		}
	}()
	for i := 0; i < 3; i++ {
		go func() {
			res := rdb.LPush(ctx, ListName03, "list-01")
			fmt.Println(res.String())
			res = rdb.LPush(ctx, ListName04, "list-01")
			fmt.Println(res.String())
		}()
	}
	wg.Wait()
}

func BRpop() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			res := rdb.BRPop(ctx, 3*time.Second, ListName03, ListName04)
			strs, _ := res.Result()
			fmt.Println(strs)
		}
	}()
	for i := 0; i < 3; i++ {
		go func() {
			res := rdb.LPush(ctx, ListName03, "list-01")
			fmt.Println(res.String())
			res = rdb.LPush(ctx, ListName04, "list-01")
			fmt.Println(res.String())
		}()
	}
	wg.Wait()
}

func LIndex() {
	res := rdb.LIndex(ctx, ListName01, 0)
	str, _ := res.Result()
	fmt.Println(str)

	res = rdb.LIndex(ctx, ListName02, -1)
	str, _ = res.Result()
	fmt.Println(str)
}

func LInsert() {
	res := rdb.LInsert(ctx, ListName02, "BEFORE", "list-01", "list-02")
	fmt.Println(res.String())

	res = rdb.LInsert(ctx, ListName02, "AFTER", "list-05", "list-04")
	fmt.Println(res.String())

	sres := rdb.LRange(ctx, ListName02, 0, -1)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func LLen() {
	res := rdb.LLen(ctx, ListName01)
	fmt.Println(res.String())
}

func LPop() {
	res := rdb.LPop(ctx, ListName01)
	str, _ := res.Result()
	fmt.Println(str)
}

func LPush() {
	res := rdb.LPush(ctx, ListName01, "list-02")
	fmt.Println(res.String())
}

func LPushX() {
	res := rdb.LPushX(ctx, ListName01, "list-03")
	fmt.Println(res.String())

	sres := rdb.LRange(ctx, ListName01, 0, -1)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func LRange() {
	sres := rdb.LRange(ctx, ListName01, 0, -1)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func LRem() {
	res := rdb.LRem(ctx, ListName01, 0, "list-01")
	fmt.Println(res.String())
}

func LSet() {
	res := rdb.LSet(ctx, ListName01, 0, "list-set-01")
	fmt.Println(res.String())

	sres := rdb.LRange(ctx, ListName01, 0, -1)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func LTrim() {
	res := rdb.LTrim(ctx, ListName01, 0, 0)
	fmt.Println(res.String())

	LRange()
}

func RPop() {
	res := rdb.RPop(ctx, ListName02)
	str, _ := res.Result()
	fmt.Println(str)
}

func RPopLPush() {
	res := rdb.RPopLPush(ctx, ListName01, ListName02)
	str, _ := res.Result()
	fmt.Println(str)

	LRange()
}

func RPush() {
	_ = rdb.RPush(ctx, ListName01, "list-rpush-01")
	_ = rdb.RPush(ctx, ListName01, "list-rpush-02")
	LRange()
}

func RPushX() {
	_ = rdb.RPushX(ctx, ListName01, "list-rpush-03")
	_ = rdb.RPushX(ctx, ListName01, "list-rpush-04")
	LRange()
}

func BRPopLPush() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		res := rdb.BRPopLPush(ctx, ListName05, ListName06, 3*time.Second)
		fmt.Println(res.String())
	}()
	go func() {
		defer wg.Done()
		res := rdb.LPush(ctx, ListName05, "list-brpoplpush-01")
		fmt.Println(res.String())
	}()
	wg.Wait()

	sres := rdb.LRange(ctx, ListName06, 0, -1)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func List() {
	ListInit()
	BLpop()
	BRpop()
	LIndex()
	LInsert()
	LLen()
	LPop()
	LPush()
	LPushX()
	LRange()
	LRem()
	LSet()
	LTrim()
	RPop()
	RPopLPush()
	RPush()
	RPushX()
	BRPopLPush()
}
