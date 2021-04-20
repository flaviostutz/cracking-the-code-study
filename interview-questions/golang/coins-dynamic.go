package main

import "fmt"

func coinsMinRecursive(memo map[string]int, coins []int, amount int, cn int) int {
	k := fmt.Sprintf("%d-%d", amount, cn)
	v, ok := memo[k]
	if ok {
		return v
	}

	s := ""
	for i := 0; i < (10 - cn); i++ {
		s += "  "
	}
	// fmt.Printf("%s select %d %d\n", s, amount, cn)
	if amount == 0 {
		// fmt.Printf("%s ZERO %d %d\n", s, amount, cn)
		return 0
	}

	if amount < 0 || cn == 0 {
		// fmt.Printf("%s IMPOSSIBLE %d %d\n", s, amount, cn)
		return -1
	}

	cur := 999999999

	//branch: select current and try same coin again
	// fmt.Printf(">>select coin %d and try same\n", coins[cn-1])
	r1 := coinsMinRecursive(memo, coins, amount-coins[cn-1], cn)
	if r1 != -1 {
		cur = r1 + 1
	}
	// r1 := 0
	// r3 := 0
	//branch: select current and try next coin
	// fmt.Printf(">>select coin %d and try next\n", coins[cn-1])
	r2 := coinsMinRecursive(memo, coins, amount-coins[cn-1], cn-1)
	if r2 != -1 && (r2+1) < cur {
		cur = r2 + 1
	}

	//branch: skip current and try next coin
	// fmt.Printf(">>skip current %d and try next\n", coins[cn-1])
	r3 := coinsMinRecursive(memo, coins, amount, cn-1)
	if r3 != -1 && r3 < cur {
		cur = r3
	}

	// fmt.Printf("%s %d %d %d\n", s, r1, r2, r3)
	if cur == 999999999 {
		cur = -1
	}

	memo[k] = cur

	return cur
}

func coinsMinBottomUp(coins []int, amount int, cn int) int {
	t := make([][]int, amount+1)
	for am := 0; am <= amount; am++ {
		t[am] = make([]int, cn+1)
		for c := 0; c <= cn; c++ {
			if am == 0 {
				t[am][c] = 0
				continue
			}
			//am < 0 ||
			if c == 0 {
				t[am][c] = -1
				continue
			}

			cur := 999999999

			if am-coins[c-1] >= 0 {
				r1 := t[am-coins[c-1]][c]
				if r1 != -1 {
					cur = r1 + 1
				}
			}

			if am-coins[c-1] >= 0 {
				r2 := t[am-coins[c-1]][c-1]
				if r2 != -1 && (r2+1) < cur {
					cur = r2 + 1
				}
			}

			// c < len(coins) &&
			r3 := t[am][c-1]
			if r3 != -1 && r3 < cur {
				cur = r3
			}

			if cur == 999999999 {
				cur = -1
			}

			t[am][c] = cur

			// fmt.Printf("111 %v\n", t)
		}
	}
	// fmt.Printf("222 %v\n", t)
	return t[amount][cn]
}
