package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLISBacktracking1(t *testing.T) {
	arr := []int{4, 1, 3, 2, 1, 5}
	len := lisdiscover(arr)
	fmt.Printf("length: %d\n", len)
	assert.Equal(t, 3, len)
}

func TestLISBacktracking2(t *testing.T) {
	arr := []int{4, 1, 3, 2, 3, 1, 5, 9, 0}
	len := lisdiscover(arr)
	fmt.Printf("length: %d\n", len)
	assert.Equal(t, 4, len)
}
