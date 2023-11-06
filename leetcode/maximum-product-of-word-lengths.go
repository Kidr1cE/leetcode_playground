/*
Date: 2023-11-6
ProblemID: 318
ProblemName: 最大单词长度乘积
*/
package leetcode

func maxProduct(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
