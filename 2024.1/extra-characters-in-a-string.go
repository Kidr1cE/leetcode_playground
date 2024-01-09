/*
Date: 2023-12-9
ProblemID: 2707
ProblemName: 字符串中的额外字符
*/
package leetcode

func minExtraChar(s string, dictionary []string) int {
	// make dictionary a set
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}

	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	// dfs(i) returns the minimum number of extra characters in s[:i+1]
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		res := dfs(i-1) + 1

		for j := 0; j <= i; j++ {
			if has[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}

		*p = res
		return res
	}
	return dfs(n - 1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
