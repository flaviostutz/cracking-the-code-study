package main

func combinationSum(input []int, k int) [][]int {
	result := make([][]int, 0)
	combSum(input, k, 0, 0, []int{}, &result)
	return result
}

//USING BACKTRACKING PATTERN
func combSum(input []int, target int, start int, sum int, candidates []int, result *[][]int) {
	cand := make([]int, len(candidates))
	copy(cand, candidates)
	// fmt.Printf("input=%v target=%d start=%d candidates=%v\n", input, target, start, cand)
	//is valid (goal)
	// cand := make([]int, len(candidates))
	// copy(cand, candidates)
	if sum > target {
		return
	}
	if sum == target {
		// v := make([]int, len(cand))
		// copy(v, cand)
		*result = append(*result, cand)
		// fmt.Printf("FOUND %v\n", cand)
		return
	}

	//explore candidates
	for i := start; i < len(input); i++ {
		//add candidate
		cand = append(cand, input[i])

		//evolve solution
		combSum(input, target, i+1, sum+input[i], cand, result)

		//remove candidate
		cand = cand[:len(cand)-1]
	}

}
