package main

import (
	"fmt"
)

func search(nums []int, val int) int {
	l := 0
	r := len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 7}
	fmt.Println(search(arr, 1))
	fmt.Println(search(arr, 4))
	fmt.Println(search(arr, 7))
	arr = []int{1, 2, 3, 3, 4, 5}
	fmt.Println(search(arr, 1))
	fmt.Println(search(arr, 3))
	fmt.Println(search(arr, 5))
}
