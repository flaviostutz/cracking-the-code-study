package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestMinJumpCost1(t *testing.T) {
	cost := []int{2, 1, 3, 1}
	r := minJumpCostTab(cost, 1)
	assert.Equal(t, 2, r)

	cost = []int{2, 1, 3, 1}
	r = minJumpCostTab(cost, 2)
	assert.Equal(t, 2, r)

	cost = []int{2, 1, 3, 1}
	r = minJumpCostTab(cost, 3)
	assert.Equal(t, 1, r)

	cost = []int{1, 1, 1, 1, 1, 1, 1}
	r = minJumpCostTab(cost, 1)
	assert.Equal(t, 4, r)

	cost = []int{2, 5, 6, 3, 1}
	r = minJumpCostTab(cost, 2)
	assert.Equal(t, 6, r)

}

func TestMinJumpCost2(t *testing.T) {
	cost := []int{0, 20, 30, 40, 25, 15, 20, 28}
	r := minJumpCostTab(cost, 0)
	assert.Equal(t, 178, r)

	r = minJumpCostTab(cost, 1)
	assert.Equal(t, 98, r)

	r = minJumpCostTab(cost, 2)
	assert.Equal(t, 73, r)

	r = minJumpCostTab(cost, 3)
	assert.Equal(t, 53, r)

	r = minJumpCostTab(cost, 4)
	assert.Equal(t, 43, r)

	r = minJumpCostTab(cost, 5)
	assert.Equal(t, 43, r)

	r = minJumpCostTab(cost, 6)
	assert.Equal(t, 28, r)

}
