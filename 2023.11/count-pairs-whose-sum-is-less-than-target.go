/*
Date: 2023-10-26
ProblemID: 2824
ProblemName: 统计和小于目标的下标对数目
*/
package leetcode

func countPairs(nums []int, target int) int {
	var res = 0
	var n = len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}
	return res
}
