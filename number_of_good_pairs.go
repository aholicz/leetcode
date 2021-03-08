package main

import "fmt"

// https://leetcode.com/problems/number-of-good-pairs/

func main() {
	fmt.Print(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

//opt
func solution1(nums []int) int {
	c, m := 0, make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		c += m[nums[i]]
		m[nums[i]]++
	}
	return c
}

func numIdenticalPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				result++
			}
		}
	}
	return result
}
