package main

import "fmt"

// https://leetcode.com/problems/split-a-string-in-balanced-strings

func main() {

	fmt.Print(balancedStringSplit("RLRRLLRLRL"))
}

//opt
func solution1(s string) int {
	count, ans := 0, 0

	for _, r := range s {
		if r == 'R' {
			count++
		} else {
			count--
		}

		if count == 0 {
			ans++
		}
	}

	return ans
}

func balancedStringSplit(s string) int {
	table := make(map[string]int)
	result := 0
	for x := 0; x < len(s); x++ {
		table[string(s[x])] += 1
		if table["L"] == table["R"] {
			result += 1
		}
	}

	return result
}
