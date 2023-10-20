package brimotaskwithgolang

import "sort"

func Pendulum(values []int) []int {
	sort.Ints(values)

	var num []int
	for i, v := range values {
		if i%2 == 0 {
			num = append([]int{v}, num...)
		} else {
			num = append(num, v)
		}
	}
	return num
}
