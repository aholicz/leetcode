package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/defanging-an-ip-address

func main() {
	fmt.Print(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
