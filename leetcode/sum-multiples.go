/*
Date: 2023-10-17
ProblemID: 2652
ProblemName: 倍数求和
*/

package leetcode

func sumOfMultiples(n int) int {
	var res = 0
	for i := 3; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}
