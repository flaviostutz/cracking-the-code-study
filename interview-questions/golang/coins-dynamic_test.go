package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinsMinRecursive1(t *testing.T) {
	memo := make(map[string]int)
	coins := []int{1, 2, 5}
	r := coinsMinRecursive(memo, coins, 11, len(coins))
	assert.Equal(t, 3, r)
}

func TestCoinsMinBottomUp1(t *testing.T) {
	coins := []int{1, 2}
	r := coinsMinBottomUp(coins, 4, len(coins))
	assert.Equal(t, 2, r)
}

func TestCoinsMinBottomUp2(t *testing.T) {
	coins := []int{1, 2, 5}
	r := coinsMinBottomUp(coins, 11, len(coins))
	assert.Equal(t, 3, r)
}
