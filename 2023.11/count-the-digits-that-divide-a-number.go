/*
Date: 2023-10-26
ProblemID: 2520
ProblemName: 统计能整除数字的位数
*/
package leetcode

import "strconv"

func countDigits(num int) int {
	var res = 0
	str := strconv.Itoa(num)
	for _, val := range str {
		curr, _ := strconv.Atoi(string(val))
		if num%curr == 0 {
			res++
		}
	}
	return res
}
