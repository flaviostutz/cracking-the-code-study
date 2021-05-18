package main

func allAnagrams(str string) []string {
	result := make([]string, 0)
	nextAnagram("", str, &result)
	return result
}

func nextAnagram(prefix string, suffix string, result *[]string) {
	//is valid (goal)
	if suffix == "" {
		*result = append(*result, prefix)
		return
	}
	//explore candidates
	for i := 0; i < len(suffix); i++ {
		s := suffix[:i] + suffix[i+1:]
		p := prefix + string(suffix[i])
		nextAnagram(p, s, result)
	}
}
