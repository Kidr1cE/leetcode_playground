/*
Date: 2023-10-27
ProblemID: 1465
ProblemName: 切割后面积最大的蛋糕
*/
package leetcode

import "sort"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	var max_h, max_v int
	horizontalCuts = append(append([]int{0}, horizontalCuts...), h)
	verticalCuts = append(append([]int{0}, verticalCuts...), w)
	for i := 0; i < len(horizontalCuts)-1; i++ {
		max_h = max(max_h, horizontalCuts[i+1]-horizontalCuts[i])
	}
	for i := 0; i < len(verticalCuts)-1; i++ {
		max_v = max(max_v, verticalCuts[i+1]-verticalCuts[i])
	}
	return max_v * max_h % 1000000007
}
