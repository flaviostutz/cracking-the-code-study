package main

import (
	"math"
)

// var c int

func longSubseqRecur(seq []int) []int {
	// c = 0
	ls, valid := longSubseqRecur1(seq, -1, []int{})
	if valid {
		return ls
	}
	return []int{}
}

//O(S^S)
//using backtracking pattern
func longSubseqRecur1(seq []int, i int, cand []int) ([]int, bool) {
	// fmt.Printf("%d\n", c)
	// c++
	// fmt.Printf("longSubseqRecur1 %v %d %v\n", seq, i, cand)

	c := make([]int, len(cand))
	copy(c, cand)

	//is valid candidate
	if len(c) >= 2 && c[len(c)-1] <= c[len(c)-2] {
		// fmt.Printf("invalid %v\n", c)
		return c, false
	}

	//base case
	if i == len(seq)-1 {
		// fmt.Printf("end\n")
		return c, true
	}

	//explore candidates
	longest := c
	for j := i + 1; j < len(seq); j++ {
		// fmt.Printf("explore cand %d\n", j)

		//add candidate
		c = append(c, seq[j])

		//evolve solution
		// fmt.Printf("evolve %v %d %d\n", c, j, seq[j])
		ls, valid := longSubseqRecur1(seq, i+1, c)
		// fmt.Printf("evolve result %v %v\n", ls, valid)
		if valid && len(ls) > len(longest) {
			// fmt.Printf("longest %d\n", ls)
			longest = ls
		}

		//remove candidate (backtrack)
		c = c[:len(c)-1]
	}

	// fmt.Printf("valid %v\n", longest)

	return longest, true
}

func longSubseqTab1(seq []int) []int {
	// state
	// parameters
	//// array position i
	// cost function - what are we trying to find out?
	//// "find the longest increasing subsequence until pos i in array"
	lis := make([]int, len(seq))

	//transitions
	////base case - if there is only one element, it is the longest increasing subsequence itself
	lis[0] = 1
	////recurrence relation
	//lis[i] = max(lis[j]+1 if seq[i]>seq[j])
	//         max(lis[j]   if seq[i]<seq[j])

	for i := 1; i < len(seq); i++ {
		ma := math.MinInt32
		for j := 0; j < i; j++ {
			l := lis[j]
			if seq[i] > seq[j] {
				l++
			}
			if l > ma {
				ma = l
			}
		}
		lis[i] = ma
	}

	s := make([]int, 0)
	pl := lis[0]
	first := true
	for i := 1; i < len(lis); i++ {
		if lis[i] > pl {
			if first {
				s = append(s, seq[i-1])
				first = false
			}
			s = append(s, seq[i])
		}
		pl = lis[i]
	}

	//WRONG ANSWER BUT CORRECT LENGTH
	// fmt.Printf("%v %v %v\n", seq, lis, s)

	return s
}
