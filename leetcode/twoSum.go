package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	temps := [][]int{}
	/*타켓보다 큰 숫자를 뺀 숫자열을 인덱스화 */
	for index, value := range nums {

		// if value <= target {

		// 	row := []int{index, value}

		// 	temps = append(temps, row)

		// }

		row := []int{index, value}

		temps = append(temps, row)
	}

	//var sum int
	for i, v := range temps {

		//fmt.Printf("%d : %d : %d\n", i, v[0], v[1])

		/*0, 1, 2*/
		/*2, 7, 5*/
		/*최초 루프 1에 맨처음 배열의 값이 셋팅*/
		var _index1 int = v[0] // 0
		var _value1 int = v[1] // 2

		//fmt.Printf("_index1:%d _value1:%d\n", _index1, _value1)

		/*최초 배열 이후의 인덱스와 값을 저장하는 변수*/
		var _index2 int = 0
		var _value2 int = 0

		// /*두번째 배열*/
		for _i, _v := range temps {
			/*첫번째 배열과 인덱스값이 같지 않을경우 변수에 값을 셋팅 */
			if i != _i {
				_index2 = _v[0]
				_value2 = _v[1]
				//fmt.Printf("_index2:%d _value2:%d\n", _index2, _value2)

				if (_value1 + _value2) == target {
					// result = append(result, _index1)
					// result = append(result, _index2)

					result = []int{_index1, _index2}

					//fmt.Printf("result %d\n", result)
				}

			} else {
				_index2 = 0
				_value2 = 0
			}
		}

	}
	//fmt.Printf("sum : %d\n", sum)

	fmt.Printf("%d", result)

	return result
}
