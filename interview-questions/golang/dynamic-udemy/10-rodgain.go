package main

import "math"

func rodGainRecur1(prices []int, length int) int {
	memo := make(map[int]int, length)
	return rodGainRecur(prices, length, memo)
}

func rodGainRecur(prices []int, length int, memo map[int]int) int {
	v, ok := memo[length]
	if ok {
		return v
	}

	//base case
	if length == 0 {
		return 0
	}

	if length < 0 {
		return math.MinInt32
	}

	ma := math.MinInt32
	for i := 0; i < len(prices); i++ {
		pl := i + 1
		v := prices[i] + rodGainRecur(prices, length-pl, memo)
		if v > ma {
			ma = v
		}
	}
	memo[length] = ma
	return ma
}

func rodGainTab(prices []int, length int) int {
	//state
	dp := make([]int, length+1)
	dp[0] = 0

	//transition
	//gain[i] = max(p[a] + gain[i-a]) - a => all prices

	//fill table
	for l := 1; l <= length; l++ {
		dp[l] = math.MinInt32
		for a := 0; a < len(prices); a++ {
			pl := a + 1
			if l-pl < 0 {
				continue
			}
			v := prices[a] + dp[l-pl]
			if v > dp[l] {
				dp[l] = v
			}
		}
	}
	return dp[length]
}
