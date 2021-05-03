package main

import "fmt"

func numberCombinations(input []int, k int) [][]int {
	result := make([][]int, 0)
	numbersComb(input, k, 0, []int{}, &result)
	return result
}

//USING BACKTRACKING PATTERN
func numbersComb(input []int, k int, start int, set []int, result *[][]int) {
	//is valid (goal)
	if len(set) == k {
		fmt.Printf(">> %v\n", set)
		*result = append(*result, set)
	}

	if start == len(input) {
		return
	}

	//explore candidates
	for i := start; i < len(input); i++ {
		//add candidate
		set = append(set, input[i])

		//evolve solution
		numbersComb(input, k, i+1, set, result)

		//remove candidate
		set = set[:len(set)-1]
	}

}
