package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestHouseRob1(t *testing.T) {
	money := []int{1, 1, 1}
	r := houseRobTab1(money)
	assert.Equal(t, 2, r)

	money = []int{1, 10, 1}
	r = houseRobTab1(money)
	assert.Equal(t, 10, r)

	money = []int{1, 11, 56, 56, 47, 33}
	r = houseRobTab1(money)
	assert.Equal(t, 104, r)

}

func TestHouseRob2(t *testing.T) {
	money := []int{1, 1, 1}
	r := houseRobTab2(money)
	assert.Equal(t, 2, r)

	money = []int{1, 10, 1}
	r = houseRobTab2(money)
	assert.Equal(t, 10, r)

	money = []int{1, 11, 56, 56, 47, 33}
	r = houseRobTab2(money)
	assert.Equal(t, 104, r)

}
