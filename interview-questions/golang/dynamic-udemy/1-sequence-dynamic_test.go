package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//IS SEQUENCE? ex.: 1,2,3,4,5 yes; 1,2,4,5 no
func TestIsSequenceRecursive(t *testing.T) {

	elems := []int{1, 2, 3, 4, 5}
	r := isSequenceRecursive(elems, len(elems)-1)
	assert.True(t, r)

	elems = []int{1, 2, 4, 5}
	r = isSequenceRecursive(elems, len(elems)-1)
	assert.False(t, r)

	elems = []int{5, 4, 3, 2}
	r = isSequenceRecursive(elems, len(elems)-1)
	assert.False(t, r)

	elems = []int{1}
	r = isSequenceRecursive(elems, len(elems)-1)
	assert.True(t, r)

}

func TestIsSequenceBottomUp(t *testing.T) {

	elems := []int{1, 2, 3, 4, 5}
	r := isSequenceBottomUp(elems)
	assert.True(t, r)

	elems = []int{1, 2, 4, 5}
	r = isSequenceBottomUp(elems)
	assert.False(t, r)

	elems = []int{5, 4, 3, 2}
	r = isSequenceBottomUp(elems)
	assert.False(t, r)

	elems = []int{1}
	r = isSequenceBottomUp(elems)
	assert.True(t, r)

}
