package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {

	fmt.Printf("x : %d\n", x)

	var s string
	s = strconv.Itoa(x)

	s2 := strings.Split(s, "")

	temps := make([]string, len(s2)) //[]string{}
	//fmt.Printf("temps len : %d\n", (len(temps)))

	for i := 0; i < len(s2); i++ {

		//fmt.Printf(">>>>:%d\n", i)

		//fmt.Println(s2[(len(s2)-i)-1])

		temps[i] = s2[len(s2)-i-1]
	}

	fmt.Println(strings.Join(temps[:], "")) // world!

	str := strings.Join(temps[:], "")
	if x < 0 {
		str = "-" + str
	}

	result, err := strconv.Atoi(str)

	if err == nil {
		return 0

	}

	return result
}
