package main

import "fmt"

const SetName01 = "test-set-01"
const SetName02 = "test-set-02"
const SetName03 = "test-set-03"
const SetName04 = "test-set-04"
const SetName05 = "test-set-05"

func setInit() {
	res := rdb.SAdd(ctx, SetName01, "set-01", "set-02")
	fmt.Println(res.String())

	res = rdb.SAdd(ctx, SetName02, "set-01", "set-03")
	fmt.Println(res.String())
}

func scard() {
	res := rdb.SCard(ctx, SetName01)
	fmt.Println(res.String())

	res = rdb.SCard(ctx, SetName02)
	fmt.Println(res.String())
}

func smembers() {
	res := rdb.SMembers(ctx, SetName01)
	strs, _ := res.Result()
	fmt.Println(strs)

	res = rdb.SMembers(ctx, SetName02)
	strs, _ = res.Result()
	fmt.Println(strs)
}

func sdiff() {
	res := rdb.SDiff(ctx, SetName01, SetName02)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func sdiffstore() {
	res := rdb.SDiffStore(ctx, SetName03, SetName01, SetName02)
	fmt.Println(res.String())

	sres := rdb.SMembers(ctx, SetName03)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func sinter() {
	res := rdb.SInter(ctx, SetName01)
	strs, _ := res.Result()
	fmt.Println(strs)

	res = rdb.SInter(ctx, SetName02, SetName02)
	strs, _ = res.Result()
	fmt.Println(strs)
}

func sinterstore() {
	res := rdb.SInterStore(ctx, SetName04, SetName01, SetName02)
	fmt.Println(res.String())

	sres := rdb.SMembers(ctx, SetName04)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func sismember() {
	res := rdb.SIsMember(ctx, SetName01, "set-01")
	fmt.Println(res.Val())

	res = rdb.SIsMember(ctx, SetName01, "set-05")
	fmt.Println(res.Val())
}

// smove 原子操作
func smove() {
	res := rdb.SMove(ctx, SetName01, SetName02, "set-02")
	fmt.Println(res.Val())

	sres := rdb.SMembers(ctx, SetName01)
	strs, _ := sres.Result()
	fmt.Println(strs)

	sres = rdb.SMembers(ctx, SetName02)
	strs, _ = sres.Result()
	fmt.Println(strs)
}

func spop() {
	res := rdb.SPop(ctx, SetName02)
	str, _ := res.Result()
	fmt.Println(str)

	sres := rdb.SMembers(ctx, SetName02)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

// srandmember 随机返回但并不移除
func srandmember() {
	res := rdb.SRandMember(ctx, SetName01)
	str, _ := res.Result()
	fmt.Println(str)
}

func srem() {
	res := rdb.SRem(ctx, SetName01, "set-01")
	fmt.Println(res.String())

	sres := rdb.SMembers(ctx, SetName01)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func sunion() {
	res := rdb.SUnion(ctx, SetName01, SetName02, SetName03, SetName04)
	strs, _ := res.Result()
	fmt.Println(strs)
}

func sunionstore() {
	res := rdb.SUnionStore(ctx, SetName05, SetName01, SetName02, SetName03, SetName04)
	fmt.Println(res.String())

	sres := rdb.SMembers(ctx, SetName05)
	strs, _ := sres.Result()
	fmt.Println(strs)
}

func sscan() {
	var strs []string
	var cursor uint64
	for {
		res := rdb.SScan(ctx, SetName05, cursor, "*", 1)
		strs, cursor, _ = res.Result()
		fmt.Println(strs, cursor)
		if cursor == 0 {
			break
		}
	}
}

func Set() {
	setInit()
	scard()
	smembers()
	sdiff()
	sdiffstore()
	sinter()
	sinterstore()
	sismember()
	smove()
	spop()
	srandmember()
	srem()
	sunion()
	sunionstore()
	sscan()
}
