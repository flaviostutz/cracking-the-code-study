package main

import (
	"math"
)

func minJumpCostTab(cost []int, skip int) int {
	cost = append([]int{0}, cost...)

	//state
	//cost for jumping from pos i until end
	dp := make([]int, len(cost))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	//cost function
	//dp[i] = min((dp[i+1]+cost[i+1]), (dp[i+2],cost[i+2]) +... until skip count)

	//base case
	l := len(dp) - 1
	dp[l] = 0

	//describe solution
	sol := make([]int, len(dp))

	//transition/fill table
	for i := l - 1; i >= 0; i-- {
		for j := 0; j <= skip; j++ {
			if i+j+1 < len(dp) {
				// fmt.Printf("i=%d j=%d dp=%v cost=%v\n", i, j, dp, cost)
				if dp[i+j+1]+cost[i+j+1] < dp[i] {
					sol[i] = i + j + 1
				}
				dp[i] = min3(dp[i], dp[i+j+1]+cost[i+j+1])
			}
		}
	}
	// fmt.Printf("%v\n", sol)

	elems := make([]int, 0)
	i := 0
	for i < len(cost)-1 {
		i = sol[i]
		elems = append(elems, i)
	}

	// fmt.Printf("elems=%v\n", elems)
	// fmt.Printf("dp=%v\n", dp)

	return dp[0]
}

func min3(i, j int) int {
	if i < j {
		return i
	}
	return j
}
