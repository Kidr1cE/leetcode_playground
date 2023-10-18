/*
Date: 2023-10-18
ProblemID: 2530
ProblemName: 执行 K 次操作后的最大分数
*/
package leetcode

import "sort"

func MaxKelements(nums []int, k int) int64 {
	var n = len(nums)
	var res int64 = 0
	var queue = make(chan int, n)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var ni, a = 1, 0
	var qmax int
	var qk = false

	// nums[0]
	res += int64(nums[0])
	if nums[0]%3 != 0 {
		a = 1
	}
	qmax = nums[0]/3 + a

	// curr := queue[qi] < nums[i] ? nums[i] : queue[qi]
	for i := 1; i < k; i++ {
		// get current max
		var curr int
		if ni < n && qmax < nums[ni] {
			curr = nums[ni]
			ni++
		} else {
			curr = qmax
			qk = true
		}
		// add to res
		res += int64(curr)
		// /3 & get qmax
		if curr%3 != 0 {
			a = 1
		} else {
			a = 0
		}
		queue <- curr/3 + a
		if qk {
			qmax = <-queue
			qk = false
		}
	}
	return res
}
