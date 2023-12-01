/*
Date: 2023-10-19
ProblemID: 1726
ProblemName: 同积元组
*/
package leetcode

func tupleSameProduct(nums []int) int {
	var res, n = 0, len(nums)
	var cnt = make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt[nums[i]*nums[j]]++
		}
	}
	for _, val := range cnt {
		// val = 2，3，4，......
		res += val * (val - 1) * 4
	}

	return res * 8
}
