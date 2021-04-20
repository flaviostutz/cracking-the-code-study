package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermStr(t *testing.T) {

	result := make([]string, 0)
	permStr("", "ABC", &result)
	assert.Equal(t, 6, len(result))

	result = make([]string, 0)
	permStr("", "ABCD", &result)
	assert.Equal(t, 24, len(result))

	result = make([]string, 0)
	permStr("", "ABCDE", &result)
	assert.Equal(t, 120, len(result))
}

func TestPermArray(t *testing.T) {

	result := make([][]int, 0)
	permArray([]int{}, []int{1, 2, 3}, &result)
	assert.Equal(t, 6, len(result))

	result = make([][]int, 0)
	permArray([]int{}, []int{1, 2, 3}, &result)
	assert.Equal(t, 6, len(result))
	// fmt.Printf("%v\n", result)

}
