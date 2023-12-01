/*
Date: 2023-10-12
ProblemID: 2562
ProblemName: 找出数组的串联值
*/

package leetcode

import "strconv"

func FindTheArrayConcVal(nums []int) int64 {
	var n = len(nums)
	var res int64 = 0
	if n%2 == 1 {
		res += int64(nums[n/2])
	}
	for i := 0; i < n/2; i++ {
		str := strconv.Itoa(nums[i]) + strconv.Itoa(nums[n-i-1])
		num, _ := strconv.Atoi(str)
		res += int64(num)
	}
	return res
}
