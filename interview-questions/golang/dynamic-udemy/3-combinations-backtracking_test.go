package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//PRINT ALL COMBINATIONS OF K ELEMENTS FROM INPUT

func TestCombination1(t *testing.T) {

	r := numberCombinations([]int{1}, 1)
	assert.Equal(t, 1, len(r))

	r = numberCombinations([]int{1, 5}, 2)
	assert.Equal(t, 1, len(r))

	r = numberCombinations([]int{1, 5, 7}, 2)
	assert.Equal(t, 3, len(r))

	r = numberCombinations([]int{1, 5, 7, 8}, 3)
	assert.Equal(t, 4, len(r))

}
