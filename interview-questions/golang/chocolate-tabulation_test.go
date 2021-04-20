package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChocolateRecursive1(t *testing.T) {
	chocolates := []int{1, 4, 5, 3, 2, 3}
	r := eat(chocolates, 1)
	assert.Equal(t, 75, r)
}

func TestChocolateRecursive2(t *testing.T) {
	chocolates := []int{1, 4, 5, 2, 3}
	r := eat(chocolates, 1)
	assert.Equal(t, 54, r)
}

func TestChocolateTabulation1(t *testing.T) {
	choco := []int{4, 5, 3}

	memo := JoyTab(choco)
	r := memo[0][len(choco)]
	fmt.Println(r)
	// for i, j := 0, len(choco); i < j; {
	// 	if memo[i][j].pickLeft {
	// 		fmt.Print("left ")
	// 		i++
	// 	} else {
	// 		fmt.Print("right ")
	// 		j--
	// 	}
	// }
	// fmt.Println()

	assert.Equal(t, 26, r)
}
