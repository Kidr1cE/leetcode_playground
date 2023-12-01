/*
Date: 2023-10-29
ProblemID: 274
ProblemName: H 指数
*/
package leetcode

func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, n+1)
	for _, c := range citations {
		cnt[min(c, n)]++
	}
	s := 0
	for i := n; ; i-- {
		s += cnt[i]
		if s >= i {
			return i
		}
	}
}
