package main

import (
	"strconv"
)

func digitSumRecursive(digits string, i int) int {
	if i == len(digits) {
		return 0
	}
	v, _ := strconv.Atoi(string(digits[i]))
	return v + digitSumRecursive(digits, i+1)
}

func digitSumBottomUp(digits string) int {
	tb := make([]int, len(digits))

	l := len(digits)

	dig := make([]int, 0)
	for _, s := range digits {
		v, _ := strconv.Atoi(string(s))
		dig = append(dig, v)
	}
	// fmt.Printf("%v\n", dig)

	tb[l-1] = dig[l-1]

	for i := l - 2; i >= 0; i-- {
		tb[i] = tb[i+1] + dig[i]
	}

	return tb[0]
}
