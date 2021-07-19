package main

import "fmt"

const HashName01 = "hash-name-01"
const HashName02 = "hash-name-02"
const HashName03 = "hash-name-03"

func HashInit() {
	_ = rdb.Del(ctx, HashName01)
	_ = rdb.Del(ctx, HashName02)

	res := rdb.HSet(ctx, HashName01, "hash-key-1", "hash-value-1", "hash-key-2", "hash-value-2")
	fmt.Println(res.String())
	Scan(HashName01)

	res = rdb.HSet(ctx, HashName02, map[string]interface{}{"hash-key-1": "hash-value-1", "hash-key-2": "hash-value-2"})
	fmt.Println(res.String())
	Scan(HashName02)
}

func Scan(name string) {
	var strs []string
	var cursor uint64
	for {
		res := rdb.HScan(ctx, name, cursor, "*", 1)
		strs, cursor, _ = res.Result()
		fmt.Println(strs)
		if cursor == 0 {
			break
		}
	}
}

func HDel() {
	res := rdb.HDel(ctx, HashName01, "hash-key-1")
	fmt.Println(res.String())
	Scan(HashName01)
}

func HExists() {
	res := rdb.HExists(ctx, HashName01, "hash-key-2")
	b, _ := res.Result()
	fmt.Println(b)

	res = rdb.HExists(ctx, HashName01, "hash-key-3")
	b, _ = res.Result()
	fmt.Println(b)
}

func HGet() {
	res := rdb.HGet(ctx, HashName01, "hash-key-2")
	str, _ := res.Result()
	fmt.Println(str)
}

func HGetAll() {
	res := rdb.HGetAll(ctx, HashName01)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func HIncrBy() {
	res := rdb.HIncrBy(ctx, HashName03, "count", 1)
	fmt.Println(res.String())

	sre := rdb.HGet(ctx, HashName03, "count")
	str, _ := sre.Result()
	fmt.Println(str)

	for i := 0; i < 10; i++ {
		_ = rdb.HIncrBy(ctx, HashName03, "count", 3)
	}

	sre = rdb.HGet(ctx, HashName03, "count")
	str, _ = sre.Result()
	fmt.Println(str)
}

func HIncrByFloat() {
	res := rdb.HIncrByFloat(ctx, HashName03, "count-float", 1)
	fmt.Println(res.String())

	sre := rdb.HGet(ctx, HashName03, "count-float")
	str, _ := sre.Result()
	fmt.Println(str)

	for i := 0; i < 10; i++ {
		_ = rdb.HIncrByFloat(ctx, HashName03, "count-float", 0.3)
	}

	sre = rdb.HGet(ctx, HashName03, "count-float")
	str, _ = sre.Result()
	fmt.Println(str)
}

func HKeys() {
	res := rdb.HKeys(ctx, HashName01)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func HLen() {
	res := rdb.HLen(ctx, HashName01)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func HMGet() {
	res := rdb.HMGet(ctx, HashName02, "hash-key-1", "hash-key-2")
	strs, _ := res.Result()
	fmt.Println(strs)
}

func HMSet() {
	res := rdb.HSet(ctx, HashName01, "hash-key-1", "hash-new-value-1", "hash-key-2", "hash-new-value-2")
	fmt.Println(res.String())

	sres := rdb.HGetAll(ctx, HashName01)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func HSet() {
	res := rdb.HSet(ctx, HashName01, "hash-key-1", "hash-hset-value-1")
	fmt.Println(res.String())

	re := rdb.HGet(ctx, HashName01, "hash-key-1")
	str, _ := re.Result()
	fmt.Println(str)
}

func HSetNX() {
	res := rdb.HSetNX(ctx, HashName01, "hash-key-1", "hash-setnax-value-1")
	b, _ := res.Result()
	fmt.Println(b)

	sre := rdb.HSetNX(ctx, HashName01, "hash-key-3", "hash-setnax-value-3")
	str, _ := sre.Result()
	fmt.Println(str)
}

func HVals() {
	res := rdb.HVals(ctx, HashName01)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func Hash() {
	HashInit()
	HDel()
	HExists()
	HGet()
	HGetAll()
	HIncrBy()
	HIncrByFloat()
	HKeys()
	HLen()
	HMGet()
	HMSet()
	HSet()
	HSetNX()
	HVals()
}
