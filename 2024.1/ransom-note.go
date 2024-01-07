/*
Date: 2023-12-7
ProblemID: 383
ProblemName: 赎金信
*/

package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{} // Counter()
	for _, c := range magazine {
		cnt[c-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
