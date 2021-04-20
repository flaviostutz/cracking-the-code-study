package main

// func main() {
// 	//backtracking
// 	arr := []int{4, 1, 3, 2, 1, 5}
// 	selected := make([]int, 0)
// 	lis(arr, selected)
// 	fmt.Printf("%v length: %d", selected, len(selected))
// }

func lisdiscover(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		// fmt.Printf(">>>\n")
		r, sel := lis(arr[i:], make([]int, 0))
		if r {
			if len(sel) > max {
				// fmt.Printf("%v\n", sel)
				max = len(sel)
			}
		}
	}
	return max
}

func lis(arr []int, selected []int) (bool, []int) {
	// fmt.Printf("-- %v %v\n", arr, selected)
	//goal
	if len(arr) == 0 {
		return true, selected
	}

	if len(selected) == 0 {
		selected = append(selected, arr[0])
	}

	//choice - exploration
	for i := 0; i < len(arr); i++ {
		// fmt.Printf("arr=%v i=%d\n", arr, i)
		//select choice
		selected = append(selected, arr[i])

		//constraints
		if arr[i] > arr[0] {
			// fmt.Printf("%d %v %v\n", i, arr, selected)
			var r bool
			r, selected = lis(arr[i:], selected)
			if r {
				//confirm choice
				return true, selected
			}
		}

		//unselect choice
		selected = selected[:len(selected)-1]
	}

	// fmt.Printf("arr=%v FALSE\n", arr)
	//normally this is "false", but in this case it is true
	return true, selected
}
