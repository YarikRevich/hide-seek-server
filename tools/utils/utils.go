package utils

import "math/rand"

func IsIntInArray(v int, s []int) bool {
	for _, q := range s {
		if q == v {
			return true
		}
	}
	return false
}

func GetSequenceOfRandomInt(num int, r0, r1 int) []int {
	r := make([]int, num)
	for len(r) != num {
		q := rand.Intn(r1-r0+1) + r0
		if !IsIntInArray(q, r) {
			r = append(r, q)
		}
	}
	return r
}
