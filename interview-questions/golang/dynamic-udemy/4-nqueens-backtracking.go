package main

type tup struct {
	x, y int
}

func nqueens(size int, nqueens int) []tup {
	candidates := make([]tup, 0)
	valid, cand := nqueensSolve(size, nqueens, 0, candidates)
	if valid {
		return cand
	}
	return make([]tup, 0)
}

//USING BACKTRACKING PATTERN
func nqueensSolve(size int, nqueens int, y int, cand []tup) (bool, []tup) {
	//is valid (goal)
	if !isValidQueens(cand, nqueens) {
		// fmt.Printf("INVALID %v\n", cand)
		//give up this solution path sooner as possible
		return false, cand
	}
	if len(cand) == nqueens {
		return true, cand
	}
	if y == size {
		return false, cand
	}

	// for i := 0; i < y; i++ {
	// 	fmt.Printf(" ")
	// }
	// fmt.Printf(".\n")

	candidates := make([]tup, len(cand))
	copy(candidates, cand)

	//explore candidates
	for x := 0; x < size; x++ {
		t := tup{x: x, y: y}

		//add candidate
		candidates = append(candidates, t)

		//evolve solution
		ret, cand := nqueensSolve(size, nqueens, y+1, candidates)
		if ret {
			return true, cand
		}

		//remove candidate
		candidates = candidates[:len(candidates)-1]
	}

	return false, candidates
}

func isValidQueens(queens []tup, nqueens int) bool {
	// fmt.Printf("isValidQueens %v %d\n", queens, nqueens)

	mx := make(map[int]bool)
	my := make(map[int]bool)
	for a := 0; a < len(queens); a++ {
		q := queens[a]

		//horizontal attack
		_, ok := mx[q.x]
		if ok {
			// fmt.Printf("HORIZ FAIL q.x=%d\n", q.x)
			return false
		}
		mx[q.x] = true

		//vertical attack
		_, ok = my[q.y]
		if ok {
			// fmt.Printf("VERT FAIL q.y=%d\n", q.y)
			return false
		}
		my[q.y] = true

		//diagonal attack
		for b := a + 1; b < len(queens); b++ {
			qb := queens[b]
			if abs(q.x-qb.x) == abs(q.y-qb.y) {
				// fmt.Printf("DIAG FAIL q.x=%d q.y=%d qb.x=%d qb.y=%d\n", q.x, q.y, qb.x, qb.y)
				return false
			}
		}
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}
