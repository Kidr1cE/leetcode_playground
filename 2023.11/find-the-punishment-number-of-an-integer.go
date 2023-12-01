/*
Date: 2023-10-25
ProblemID: 2698
ProblemName: 求一个整数的惩罚数
*/
package leetcode

import "strconv"

var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 符合要求
			}
			x := 0
			for j := p; j < n; j++ { // 枚举分割出从 s[p] 到 s[j] 的子串
				x = x*10 + int(s[j]-'0') // 子串对应的整数值
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
