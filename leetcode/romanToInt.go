package leetcode

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	result := 0

	sArr := strings.Split(s, "")
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	nums := make([]int, len(s))
	for i, item := range sArr {

		nums[i] = romanMap[item]

		result += nums[i]
		if i > 0 {
			if nums[i-1] < nums[i] {
				fmt.Printf(">>>>>>>>>> %v - %v\n", nums[i], nums[i-1])

				result = result - nums[i-1]*2

			}
		}

	}
	fmt.Printf(">>>>>>>>>> %v\n", result)
	return result
}
