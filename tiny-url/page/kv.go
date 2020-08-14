package page

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
)

const base = 32
const step = 1

type KV struct {
	Key   string
	Value string
}

var k int64 = 0

var data sync.Map = sync.Map{}

func Load(key string) (KV, error) {
	val, ok := data.Load(key)
	if !ok {
		return KV{}, errors.New("unknown key")
	}
	value, ok := val.(string)
	if !ok {
		return KV{}, errors.New("unknown value")
	}
	return KV{
		Key:   key,
		Value: value,
	}, nil
}

func Store(value string) KV {
	kv := KV{
		Key:   convert(getKey()),
		Value: value,
	}
	data.Store(kv.Key, kv.Value)
	return kv
}

func getKey() int64 {
	return atomic.AddInt64(&k, step)
}

func convert(x int64) string {
	return strconv.FormatInt(x, base)
}
