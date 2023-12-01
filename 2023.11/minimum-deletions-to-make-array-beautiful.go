/*
Date: 2023-11-6
ProblemID: 2216
ProblemName: 美化数组的最少删除数
*/
package leetcode

func minDeletion(nums []int) int {
	var res = 0
	var l = len(nums)
	for i := 0; i < l-1; i++ {
		if (i+res)%2 == 0 && nums[i] == nums[i+1] {
			res++
		}
	}
	if (l-res)%2 == 1 {
		res++
	}
	return res
}
