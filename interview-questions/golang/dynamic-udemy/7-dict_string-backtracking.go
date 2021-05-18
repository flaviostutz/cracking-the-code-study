package main

import "fmt"

func dictString(input string, dict []string) [][]string {
	result := make([][]string, 0)
	dic := make(map[string]bool)
	for _, v := range dict {
		dic[v] = true
	}
	dictStr(input, dic, 0, make([]string, 0), &result)
	return result
}

func dictStr(input string, dict map[string]bool, i int, partial []string, result *[][]string) {
	// fmt.Printf("input=%s i=%d\n", input, i)

	part := make([]string, len(partial))
	copy(part, partial)

	//is valid (goal)
	if i == len(input) {
		*result = append(*result, part)
		return
	}

	//explore candidates
	for b := i + 1; b <= len(input); b++ {
		s := input[i:b]
		fmt.Printf("s=%s i=%d b=%d\n", s, i, b)
		_, ok := dict[s]
		if ok {
			//add candidate
			part = append(part, s)

			//evolve solution
			dictStr(input, dict, b, part, result)

			//remove candidate
			part = part[:len(part)-1]
		}
	}
}
