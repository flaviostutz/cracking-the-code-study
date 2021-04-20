package main

import "fmt"

var c int
var memo map[int]int
var memo2 map[string]int

// var memo3 map[string]int

func init() {
	memo = make(map[int]int)
	memo2 = make(map[string]int)
	// memo3 = make(map[string]int)
}

func knapsackRecursiveInfItems(items [][]int, availWeight int) int {
	//memoization
	if true {
		v, ok := memo[availWeight]
		if ok {
			return v
		}
	}

	// fmt.Printf("availWeight=%d\n", availWeight)
	//base case
	if availWeight <= 0 {
		return 0
	}

	//subproblems
	// var mv []int
	maxValue := 0
	for i := 0; i < len(items); i++ {
		m := knapsackRecursiveInfItems(items, availWeight-items[i][0]) + items[i][1]
		// fmt.Printf("m=%d\n", m)
		if m > maxValue {
			// fmt.Printf("items[i]=%d\n", items[i])
			maxValue = m
			// mv = items[i]
		}
	}
	c++
	// fmt.Printf("mv=%d\n", mv)

	memo[availWeight] = maxValue

	return maxValue
}

func knapsackRecursiveItems(W int, wt []int, val []int, n int) int {
	// Base Case
	if n == 0 || W == 0 {
		return 0
	}

	// If weight of the nth item is more than Knapsack capacity W, then
	// this item cannot be included in the optimal solution
	if wt[n-1] > W {
		return knapsackRecursiveItems(W, wt, val, n-1)
	}

	// Return the maximum of two cases: (1) nth item included (2) not included
	vv := max2(
		val[n-1]+knapsackRecursiveItems(W-wt[n-1], wt, val, n-1),
		knapsackRecursiveItems(W, wt, val, n-1),
	)
	return vv
}

func knapsackBottomUp(W int, wt []int, val []int) int {

	wwt := make([][]int, W+1)
	n := len(wt)

	for ww := 0; ww <= W; ww++ {
		wwt[ww] = make([]int, n+1)
		for nn := 0; nn <= n; nn++ {
			if nn == 0 || ww == 0 {
				wwt[ww][nn] = 0
				continue
			}
			v2 := wwt[ww][nn-1]
			if wt[nn-1] <= ww {
				v1 := val[nn-1] + wwt[ww-wt[nn-1]][nn-1]
				if v1 > v2 {
					wwt[ww][nn] = v1
				} else {
					wwt[ww][nn] = v2
				}
			} else {
				wwt[ww][nn] = v2
			}
		}
	}

	fmt.Printf(">>> %v\n", wwt)

	return wwt[W][n]
}

func max2(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
