/*
Date: 2023-10-20
ProblemID: 2525
ProblemName: 根据规则将箱子分类
*/
package leetcode

func categorizeBox(length int, width int, height int, mass int) string {
	x := length >= 1e4 || width >= 1e4 || height >= 1e4 || length*width*height >= 1e9
	y := mass >= 100
	switch {
	case x && y:
		return "Both"
	case x:
		return "Bulky"
	case y:
		return "Heavy"
	default:
		return "Neither"
	}
}
