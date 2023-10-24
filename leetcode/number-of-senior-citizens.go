/*
Date: 2023-10-23
ProblemID: 2678
ProblemName: 老人的数目
*/

package leetcode

// charInt & 15 = int
func countSeniors(details []string) int {
	res := 0
	for _, s := range details {
		if s[11:13] > "60" {
			res++
		}
	}
	return res
}
