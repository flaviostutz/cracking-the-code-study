package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//SUM DIGITS ex.: 1234 => 10; 685 => 19

func TestDigitSumRecursive(t *testing.T) {

	r := digitSumRecursive("12345", 0)
	assert.Equal(t, 15, r)

	r = digitSumRecursive("123", 0)
	assert.Equal(t, 6, r)

	r = digitSumRecursive("1", 0)
	assert.Equal(t, 1, r)

	r = digitSumRecursive("4321", 0)
	assert.Equal(t, 10, r)

}

func TestDigitSumBottomUp(t *testing.T) {

	r := digitSumBottomUp("1234")
	assert.Equal(t, 10, r)

	r = digitSumBottomUp("123")
	assert.Equal(t, 6, r)

	r = digitSumBottomUp("1")
	assert.Equal(t, 1, r)

	r = digitSumBottomUp("4321")
	assert.Equal(t, 10, r)

}
