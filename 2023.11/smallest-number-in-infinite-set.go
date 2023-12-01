/*
Date: 2023-11-29
ProblemID: 2336
ProblemName: 无限集中的最小数字
*/
package leetcode

type SmallestInfiniteSet struct {
	Arr [1001]int
}

func Constructor() SmallestInfiniteSet {
	arr := new(SmallestInfiniteSet)
	for i := range arr.Arr {
		arr.Arr[i] = 1
	}
	return *arr
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; i < len(this.Arr); i++ {
		if this.Arr[i] == 1 {
			this.Arr[i] = 0
			return i
		}
	}
	return -1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.Arr[num] = 1
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
