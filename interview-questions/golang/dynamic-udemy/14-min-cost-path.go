package main

import (
	"math"
)

// func minCostPathRecur(input string, words []string) int {
// }

func minCostPathTab(costs [][]int) int {
	//state
	dp := make([][]int, len(costs))
	for i := 0; i < len(costs); i++ {
		dp[i] = make([]int, len(costs[i]))
	}

	//base base
	//dp[last][last] = cost[last][last]
	dp[len(costs)-1][len(costs)-1] = costs[len(costs)-1][len(costs)-1]

	//transition
	for i := 1; i <= len(costs)*2; i++ {
		// fmt.Printf("%d\n", i)

		for j := 0; j <= i; j++ {
			a := len(costs) - i + j - 1
			b := len(costs) - j - 1
			if a < 0 || b < 0 {
				continue
			}

			// fmt.Printf("%d, %d\n", a, b)

			//base case skip
			if a == len(costs)-1 && b == len(costs[a]) {
				continue
			}

			//select min path
			v1 := math.MaxInt32
			// vv1 := fmt.Sprintf("%d,%d", a+1, b)
			if a+1 < len(costs) {
				v1 = costs[a][b] + dp[a+1][b]
			}

			v2 := math.MaxInt32
			// vv2 := fmt.Sprintf("%d,%d", a, b+1)
			if b+1 < len(costs[a]) {
				v2 = costs[a][b] + dp[a][b+1]
			}

			dp[a][b] = v1
			if v2 < v1 {
				dp[a][b] = v2
				// fmt.Printf("elem=%s v2=%d\n", vv2, v2)
				// } else {
				// fmt.Printf("elem=%s v1=%d\n", vv1, v1)
			}

			// fmt.Printf("dp=%v\n", dp)
		}
	}

	// fmt.Printf("%v\n", dp)

	return dp[0][0]
}
