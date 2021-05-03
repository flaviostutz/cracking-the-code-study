package main

func isSequenceRecursive(elems []int, i int) bool {
	if i == 0 {
		return true
	}

	if elems[i]-elems[i-1] != 1 {
		return false
	}

	return isSequenceRecursive(elems, i-1)
}

func isSequenceBottomUp(elems []int) bool {
	// fmt.Printf("elems=%v\n", elems)
	tb := make([]bool, len(elems))
	tb[0] = true
	for i := 1; i < len(elems); i++ {
		tb[i] = tb[i-1] && (elems[i]-elems[i-1] == 1)
	}
	// fmt.Printf("%v\n", tb)
	return tb[len(elems)-1]
}
