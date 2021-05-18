package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//PRINT ALL COMBINATIONS OF K ELEMENTS FROM INPUT

func TestCombinationSum1(t *testing.T) {

	r := combinationSum([]int{1}, 1)
	fmt.Printf("%v\n", r)
	assert.Equal(t, 1, len(r))

	r = combinationSum([]int{1, 2, 3, 5, 7, 9, 16}, 6)
	fmt.Printf("%v\n", r)
	assert.Equal(t, 2, len(r))

	r = combinationSum([]int{1, 2, 3, 4}, 4)
	fmt.Printf("%v\n", r)
	assert.Equal(t, 2, len(r))

}
