package main

import "fmt"

// https://leetcode.com/problems/shuffle-the-array/

// odd = 2 *n - 1
// event = 2 * n - 2

func main() {
	fmt.Print(shuffle([]int{1, 3, 5, 2, 4, 7}, 3))
}

//otp
func Solution2(nums []int, n int) []int {
	res := make([]int, len(nums))
	j := 0
	for i := 0; i < n; i++ {
		res[j], res[j+1] = nums[i], nums[i+n]
		j += 2
	}
	return res
}

// me
func Solution1(nums []int, n int) []int {
	result := make([]int, len(nums))
	for i := 0; i < n; i++ {
		result[2*(i+1)-1] = nums[i+n]
		result[2*(i+1)-2] = nums[i]
	}

	return result
}

func shuffle(nums []int, n int) []int {
	head, tail := nums[:n], nums[n:]
	result := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		result = append(result, head[i], tail[i])
	}

	return result
}
