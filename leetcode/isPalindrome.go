package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func isPalindrome(x int) bool {
	result := false
	fmt.Printf("x : %d\n", x)

	if x < 0 {
		result = false
	} else {

		var s string
		s = strconv.Itoa(x)
		length := len(s)

		reverseString := make([]string, length)
		for i := 0; i < length; i++ {
			reverseString[i] = strings.Split(s, "")[(length-1)-i]
		}

		tempLength := (length / 2) + (length % 2)

		// fmt.Printf("result : %v\n", reverseString)
		// fmt.Printf("len / 2 : %d\n", length/2)
		// fmt.Printf("len %% 2 : %d\n", (length % 2))

		fmt.Printf(">>>> : %v\n", strings.Split(s, "")[0:tempLength])
		fmt.Printf(">>> : %v\n", reverseString[0:tempLength])

		s1 := strings.Join(strings.Split(s, "")[0:tempLength], "")
		s2 := strings.Join(reverseString[0:tempLength], "")

		if s1 == s2 {
			result = true
		}

	}

	fmt.Printf("result : %v\n", result)
	return result
}
