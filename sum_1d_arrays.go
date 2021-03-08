package main

import "fmt"

// https://leetcode.com/problems/running-sum-of-1d-array/

func main() {
	fmt.Print(runningSum([]int{1, 2, 3, 4}))
}

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]

	for index := 1; index < len(nums); index++ {
		result[index] = result[index-1] + nums[index]
	}

	return result
}
