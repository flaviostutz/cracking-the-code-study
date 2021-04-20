package main

import "fmt"

var (
	memo1 = make(map[string]int)
)

// func main0() {
// 	//recursive
// 	// chocolates := []int{4, 3, 2, 1, 5}
// 	// r := eat(chocolates, 1)
// 	// fmt.Printf("Pleasure: %d", r)

// 	choco := []int{4, 3, 2, 1, 5}
// 	memo := JoyTab(choco)
// 	fmt.Println(memo[0][len(choco)].joy)
// 	for i, j := 0, len(choco); i < j; {
// 		if memo[i][j].pickLeft {
// 			fmt.Print("left ")
// 			i++
// 		} else {
// 			fmt.Print("right ")
// 			j--
// 		}
// 	}
// 	fmt.Println()
// }

func eat(chocolates []int, day int) int {
	//memoization
	k := fmt.Sprintf("%v-%d", chocolates, day)
	v, ok := memo1[k]
	if ok {
		return v
	}
	// fmt.Printf("%v %d %d\n", chocolates, day, plr)

	//goal
	n := len(chocolates)
	if n == 1 {
		return chocolates[0] * day
	}

	//constraints
	//none

	//choice
	left := chocolates[0]*day + eat(chocolates[1:], day+1)
	right := chocolates[n-1]*day + eat(chocolates[:n-1], day+1)
	r := right
	if left > right {
		r = left
	}
	memo1[k] = r
	return r
}

// JoyTab returns a table containing the joys memo[i,j] that
// you get if you eat the chocolates choco[i:j] in an optimal
// order starting at day 1 + len(choco) - (j - i).
func JoyTab(choco []int) (memo [][]int) {
	n := len(choco)
	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	for i := range memo { // base cases
		memo[i][i+1] = n * choco[i]
	}
	// fmt.Printf("memo %v choco %v\n", memo, choco)
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n; j++ {
			day := 1 + n - (j - i)
			left := day*choco[i] + memo[i+1][j]
			right := day*choco[j-1] + memo[i][j-1]
			memo[i][j] = max(left, right)
			// memo[i][j].pickLeft = left > right
			// fmt.Printf("i=%d j=%d day=%d left=%d right=%d %v\n", i, j, day, left, right, memo[i][j])
		}
	}
	// fmt.Printf("memo %v\n", memo)
	return memo
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
