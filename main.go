package main

import "fmt"

// Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
// Output: [0,4,1,3,2]
// Explanation:
// nums       index     target
// 0            0        [0]
// 1            1        [0,1]
// 2            2        [0,1,2]
// 3            2        [0,1,3,2]
// 4            1        [0,4,1,3,2]

func main() {
	fmt.Print(createTargetArray([]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1}))
}

func createTargetArray(nums []int, index []int) []int {
	m := make(map[int][]int)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[index[i]]; ok {
			v = append([]int{nums[i]}, v...)
			m[index[i]] = v
		} else {
			m[index[i]] = []int{nums[i]}
		}
	}

	for _, v := range m {
		if len(v) > 1 {
			fmt.Println(v)
			result = append(result, v...)

		}
		result = append(result, v...)
	}

	return result
}
