package main

// func main() {
// 	// result := make([]string, 0)
// 	// permStr("", "ABC", &result)
// 	// fmt.Printf("Permutations: %v", result)

// 	result := make([][]int, 0)
// 	permArray([]int{}, []int{1, 2, 3}, &result)
// 	fmt.Printf("Permutations: %v", result)
// }

func permStr(prefix string, suffix string, result *[]string) {
	if suffix == "" {
		*result = append(*result, prefix)
	}
	for i := 0; i < len(suffix); i++ {
		permStr(prefix+string(suffix[i]), suffix[:i]+suffix[i+1:], result)
	}
}

func permArray(prefix []int, suffix []int, result *[][]int) {
	// fmt.Printf("p=%v s=%v\n", prefix, suffix)
	if len(suffix) == 0 {
		*result = append(*result, prefix)
	}
	for i := 0; i < len(suffix); i++ {
		// fmt.Printf(">>SS %d %v %v %d %v %v\n", prefix, i, suffix, suffix[i], suffix[:i], suffix[i+1:])
		permArray(copyAndAppend(prefix, suffix[i]), copyAndAppend(suffix[:i], suffix[i+1:]...), result)
	}
}

func copyAndAppend(i []int, vals ...int) []int {
	j := make([]int, len(i), len(i)+len(vals))
	copy(j, i)
	return append(j, vals...)
}
