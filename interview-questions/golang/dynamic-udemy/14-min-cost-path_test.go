package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestMinCostPath1(t *testing.T) {
	m := [][]int{
		{0, 47, 8, 18, 1},
		{43, 25, 39, 36, 13},
		{22, 8, 13, 38, 46},
		{41, 41, 40, 25, 44},
		{29, 43, 22, 50, 10},
	}
	r := minCostPathTab(m)
	assert.Equal(t, 187, r)

	m = [][]int{
		{0, 1},
		{29, 43},
	}
	r = minCostPathTab(m)
	assert.Equal(t, 44, r)

	m = [][]int{
		{0, 1, 2},
		{29, 43, 3},
		{19, 53, 1},
	}
	r = minCostPathTab(m)
	assert.Equal(t, 7, r)

	m = [][]int{
		{100, 1, 200},
		{1, 43, 3},
		{2, 3, 112},
	}
	r = minCostPathTab(m)
	assert.Equal(t, 218, r)

}
