/*
Date: 2023-10-11
ProblemID: 2512
ProblemName: 奖励最顶尖的 K 名学生
*/
package leetcode

import (
	"sort"
	"strings"
)

func topStudents(positiveFeedback, negativeFeedback, report []string, studentId []int, k int) []int {
	words := map[string]int{}
	for _, w := range positiveFeedback {
		words[w] = 3
	}
	for _, w := range negativeFeedback {
		words[w] = -1
	}
	type pair struct{ score, id int }
	A := make([]pair, len(report))
	for i, r := range report {
		score := 0
		for _, w := range strings.Split(r, " ") {
			score += words[w]
		}
		A[i] = pair{score, studentId[i]}
	}
	sort.Slice(A, func(i, j int) bool {
		a, b := A[i], A[j]
		return a.score > b.score || a.score == b.score && a.id < b.id
	})
	res := make([]int, k)
	for i, p := range A[:k] {
		res[i] = p.id
	}
	return res
}
