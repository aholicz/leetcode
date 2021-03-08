package main

import "fmt"

// https://leetcode.com/problems/richest-customer-wealth/

func main() {
	fmt.Print(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	sum := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}

		if sum > max {
			max = sum
		}

		sum = 0
	}

	return max
}
