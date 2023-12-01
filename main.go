package main

import (
	"fmt"
	"playground/leetcode"
)

func main() {
	ints := []int{238838724, 196963851, 539418658, 120132686, 273213807, 57187185, 68854249, 619718192}
	res := leetcode.MaxKelements(ints, 7)
	fmt.Println(res)
}
