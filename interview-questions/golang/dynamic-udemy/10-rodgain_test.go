package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestRodGain1(t *testing.T) {
	p := []int{1, 1, 1}
	r := rodGainRecur1(p, 10)
	assert.Equal(t, 10, r)
	r = rodGainTab(p, 10)
	assert.Equal(t, 10, r)

	p = []int{1, 5, 8}
	r = rodGainRecur1(p, 40)
	assert.Equal(t, 106, r)
	r = rodGainTab(p, 40)
	assert.Equal(t, 106, r)

	p = []int{1, 5, 8, 43, 21, 65}
	r = rodGainRecur1(p, 1140)
	assert.Equal(t, 12350, r)
	r = rodGainTab(p, 1140)
	assert.Equal(t, 12350, r)

}
