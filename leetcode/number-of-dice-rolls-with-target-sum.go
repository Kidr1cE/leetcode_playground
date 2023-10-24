/*
Date: 2023-10-24
ProblemID: 1155
ProblemName: 掷骰子等于目标和的方法数
*/

package leetcode

func numRollsToTarget(n, k, target int) int {
	if target < n || target > n*k {
		return 0 // 无法组成 target
	}
	const mod = 1_000_000_007
	f := make([]int, target-n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		ma := min(i*(k-1), target-n) // i 个骰子至多掷出 i*(k-1)
		for j := 1; j <= ma; j++ {
			f[j] += f[j-1] // 原地计算 f 的前缀和
		}
		for j := ma; j >= k; j-- {
			f[j] = (f[j] - f[j-k]) % mod // f[j] 是两个前缀和的差
		}
	}
	return f[target-n]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
