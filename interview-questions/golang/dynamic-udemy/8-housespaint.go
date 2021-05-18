package main

import "math"

func houseMinCostRecur(costs [][]int) int {
	return houseRecur(costs, 0, 0, -1)
}

//O((c-1)^h)
func houseRecur(costs [][]int, houseN int, currentCost int, prevColor int) int {
	if houseN == len(costs) {
		return currentCost
	}

	min := math.MaxInt32
	for i := 0; i < len(costs[houseN]); i++ {
		if i != prevColor {
			cc := houseRecur(costs, houseN+1, currentCost+costs[houseN][i], i)
			if cc < min {
				min = cc
			}
		}
	}

	return min
}

//O(hc)
func houseTab1(costs [][]int) int {
	mc := make([]int, len(costs)+1)
	mc[0] = 0

	prevI := -1
	for h := 1; h <= len(costs); h++ {
		min := math.MaxInt32
		mi := -1
		for i := 0; i < len(costs[h-1]); i++ {
			if i != prevI {
				if costs[h-1][i] < min {
					min = costs[h-1][i]
					mi = i
				}
			}
		}
		prevI = mi
		mc[h] = min + mc[h-1]
	}

	return mc[len(costs)]
}

//O(h*c*(c-1))
func houseTab2(costs [][]int) int {
	//edge cases
	if len(costs) == 0 {
		return 0
	}

	//1-define state: house, color
	tab := make([][]int, len(costs))
	for i := 0; i < len(tab); i++ {
		tab[i] = make([]int, len(costs[0]))
	}

	//2-fill base case: aggregated last house cost is self cost
	nc := len(costs[0])
	for i := 0; i < nc; i++ {
		tab[len(costs)-1][i] = costs[len(costs)-1][i]
	}

	//3-transition: min_cost(h,c) = cost(h,c) + min(min_cost(h+1, not c), min_cost(h+1, not c)...)
	//4-fill table with transition
	for i := len(costs) - 2; i >= 0; i-- {
		for c := 0; c < len(costs[i]); c++ {
			min := math.MaxInt32
			for cn := 0; cn < len(costs[i]); cn++ {
				if cn == c {
					continue
				}
				v := tab[i+1][cn]
				if v < min {
					min = v
				}
			}
			tab[i][c] = costs[i][c] + min
		}
	}

	//5-select desired case
	min := math.MaxInt32
	for c := 0; c < len(costs[0]); c++ {
		v := tab[0][c]
		if v < min {
			min = v
		}
	}

	return min
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
