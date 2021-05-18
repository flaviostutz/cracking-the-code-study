package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestHousePaintCost1(t *testing.T) {
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 9},
	}
	minCost := houseMinCostRecur(costs)
	assert.Equal(t, 10, minCost)

	minCost = houseTab1(costs)
	assert.Equal(t, 10, minCost)

	minCost = houseTab2(costs)
	assert.Equal(t, 10, minCost)

	costs = [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 9},
		{1, 3, 2},
		{1, 99, 2},
	}
	minCost = houseMinCostRecur(costs)
	assert.Equal(t, 13, minCost)

	minCost = houseTab1(costs)
	assert.Equal(t, 13, minCost)

	minCost = houseTab2(costs)
	assert.Equal(t, 13, minCost)

}
