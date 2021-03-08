package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/shuffle-string

func main() {
	fmt.Print(solution2("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

func solution2(s string, indices []int) string {
	r := make([]rune, len(s))
	for i, c := range s {
		r[indices[i]] = c
	}

	return string(r)
}

// otp
func solution1(s string, indices []int) string {
	m := make(map[int]string)
	for i := 0; i < len(indices); i++ {
		m[indices[i]] = string(s[i])
	}

	var g string
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
		g = g + m[i]
	}
	return g
}

func restoreString(s string, indices []int) string {
	appendex := make([]string, len(s))
	for i, c := range s {
		appendex[i] = string(c)
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			temp := 0
			var stemp string
			if indices[i] < indices[j] {
				temp = indices[i]
				stemp = appendex[i]
				indices[i] = indices[j]
				appendex[i] = appendex[j]

				indices[j] = temp
				appendex[j] = stemp
			}
		}
	}

	return strings.Join(appendex, "")
}
