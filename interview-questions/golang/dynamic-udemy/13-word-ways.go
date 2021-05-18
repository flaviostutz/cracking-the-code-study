package main

func wordWaysRecur(input string, words []string) int {
	dict := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}
	memo := make(map[int]int)
	r := wordWaysRecur1(input, dict, 0, memo)
	return r
}

func wordWaysRecur1(input string, dict map[string]bool, i int, memo map[int]int) int {
	cv, ok := memo[i]
	if ok {
		return cv
	}

	//end of a full solution
	if i == len(input) {
		return 1
	}

	c := 0
	for b := i; b <= len(input); b++ {
		v := input[i:b]
		_, ok := dict[v]
		if ok {
			c += wordWaysRecur1(input, dict, i+len(v), memo)
		}
	}

	memo[i] = c

	return c
}

func wordWaysTab(input string, words []string) int {
	dict := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}

	//state
	dp := make([]int, len(input)+1)

	//base case
	dp[0] = 1

	//cost function
	// dp[i] = sum(dp[j-1]) for all input[j:i] in dict

	//transition
	for a := 0; a < len(input); a++ {
		for b := a; b <= len(input); b++ {
			v := input[a:b]
			_, ok := dict[v]
			if ok {
				dp[b] += dp[a]
			}
		}
	}

	return dp[len(input)]
}
