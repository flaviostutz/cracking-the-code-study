package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestLongestSubseq1(t *testing.T) {
	seq := []int{1, 5, 4}
	r := longSubseqRecur(seq)
	// fmt.Printf("%v\n", r)
	assert.Equal(t, 2, len(r))

	r = longSubseqTab1(seq)
	// fmt.Printf("%v\n", r)
	assert.Equal(t, 2, len(r))

	seq = []int{56, 4, 5, 6, 7}
	r = longSubseqRecur(seq)
	// fmt.Printf("%v\n", r)
	assert.Equal(t, 4, len(r))

	r = longSubseqTab1(seq)
	// fmt.Printf("%v\n", r)
	assert.Equal(t, 4, len(r))

	seq = []int{56, 5, 4, 78, 5, 9, 3, 2, 99, 8, 101, 7, 102}
	r = longSubseqRecur(seq)
	// fmt.Printf("%v\n", r)
	// assert.Equal(t, 6, len(r))

	r = longSubseqTab1(seq)
	fmt.Printf("%v\n", r)
	assert.Equal(t, 6, len(r))

}
