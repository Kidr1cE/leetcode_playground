/*
Date: 2023-10-09
ProblemID: 2578
ProblemName: 最小和分割
*/
package leetcode

import "sort"

func SplitNum(num int) int {
	var nums = make([]int, 10)
	for i := 0; num > 0; i++ {
		nums[i] = num % 10
		num /= 10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	var res = 0
	var carry, pos = 0, 1
	for i := 0; i < 5; i++ {
		curr := nums[i*2] + nums[i*2+1] + carry

		if curr >= 10 {
			curr %= 10
			carry = 1
		} else {
			carry = 0
		}

		res += curr * pos
		println(curr, carry)

		pos *= 10
	}
	return res + carry*pos
}
