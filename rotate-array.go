// 将问题拆解为几个子问题，分别解决
// 之前做过这道题，但是面试的时候没做出来
// 整个问题可以分为输出方向和下标移动两个问题
package main

import (
	"fmt"
)

func direct(arr [][]int, i, j *int, count int, direct int) {
	if direct == 0 {
		for k := 0; k < count; k++ {
			*j++
			fmt.Printf("%d ", arr[*i][*j])
		}
	} else if direct == 1 {
		for k := 0; k < count; k++ {
			*i++
			fmt.Printf("%d ", arr[*i][*j])
		}
	} else if direct == 2 {
		for k := 0; k < count; k++ {
			*j--
			fmt.Printf("%d ", arr[*i][*j])
		}
	} else if direct == 3 {
		for k := 0; k < count; k++ {
			*i--
			fmt.Printf("%d ", arr[*i][*j])
		}
	}
}

func output(arr [][]int) {
	row := len(arr)
	column := len(arr[0])
	sum := row * column
	row--
	dir := 0
	i := 0
	j := -1
	count := 0
	for count < sum {
		var ln int
		if dir == 0 || dir == 2 {
			ln = column
		} else {
			ln = row
		}
		direct(arr, &i, &j, ln, dir)
		count += ln
		if dir == 0 || dir == 2 {
			column--
		} else {
			row--
		}
		dir = (dir + 1) % 4
	}
	fmt.Println()
}

func output02(arr [][]int) {
	m := len(arr)
	n := len(arr[0])
	i := 0
	j := 0
	flag := 0
	step := 0
	for m != 0 && n != 0 {
		fmt.Printf("%d ", arr[i][j])
		step++
		if flag%2 == 0 && step == n {
			flag = (flag + 1) % 4
			m--
			step = 0
		} else if flag%2 == 1 && step == m {
			flag = (flag + 1) % 4
			n--
			step = 0
		}
		if flag == 0 {
			j++
		} else if flag == 1 {
			i++
		} else if flag == 2 {
			j--
		} else if flag == 3 {
			i--
		}
	}
	fmt.Println()
}

func main() {
	arr := [][]int{}
	arr = append(arr, []int{1, 2, 3})
	arr = append(arr, []int{8, 9, 4})
	arr = append(arr, []int{7, 6, 5})
	output(arr)
	output02(arr)

	arr2 := [][]int{}
	arr2 = append(arr2, []int{1, 2, 3, 4, 5})
	arr2 = append(arr2, []int{16, 17, 18, 19, 6})
	arr2 = append(arr2, []int{15, 24, 25, 20, 7})
	arr2 = append(arr2, []int{14, 23, 22, 21, 8})
	arr2 = append(arr2, []int{13, 12, 11, 10, 9})
	output(arr2)
	output02(arr2)

	arr3 := [][]int{}
	arr3 = append(arr3, []int{1})
	output(arr3)
	output02(arr3)
}
