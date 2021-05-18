package main

import (
	"fmt"
	"math"
)

func houseRobTab1(money []int) int {
	//state
	tab := make([]int, len(money))

	//base cases
	for i := 0; i < len(money); i++ {
		tab[i] = money[i]
	}

	//transition
	// tab[a] = tab[a] + max(tab[bs]...)

	// fill table with optimization
	for a := len(money) - 1; a >= 0; a-- {
		maxa := math.MinInt32
		for b := len(money) - 1; b >= a+2; b-- {
			if tab[b] > maxa {
				maxa = tab[b]
			}
		}
		tab[a] = max(tab[a], tab[a]+maxa)
	}

	//select best
	maxa := math.MinInt32
	for b := 0; b < len(money); b++ {
		if tab[b] > maxa {
			maxa = tab[b]
		}
	}
	fmt.Printf("%v\n", tab)
	return maxa
}

func houseRobTab2(money []int) int {
	//state
	//best gain for robbing houses from [0-i]
	gain := make([]int, len(money)+1)

	//base cases
	gain[0] = 0
	gain[1] = money[0]

	//transition
	//gain[h] = max(gain[h-1], money[h]+gain[h-2])

	//fill table with optimization
	for i := 2; i <= len(money); i++ {
		gain[i] = max(gain[i-1], money[i-1]+gain[i-2])
	}

	fmt.Printf("%v\n", gain)

	//select best
	return gain[len(gain)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
