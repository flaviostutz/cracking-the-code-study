package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsackRecursiveInfItems1(t *testing.T) {
	items := make([][]int, 0)
	items = append(items, []int{6, 30})
	items = append(items, []int{3, 14})
	items = append(items, []int{4, 16})
	items = append(items, []int{2, 9})

	// items = append(items, []int{5, 29})
	// items = append(items, []int{7, 39})
	// items = append(items, []int{6, 29})
	// items = append(items, []int{9, 9})
	m := knapsackRecursiveInfItems(items, 10)
	assert.Equal(t, 74, m)
	// fmt.Printf("maxvalue=%d c=%d\n", m, c)
}

func TestKnapsackRecursiveInfItems2(t *testing.T) {
	items := make([][]int, 0)
	items = append(items, []int{1, 1})
	items = append(items, []int{3, 4})
	items = append(items, []int{4, 5})
	items = append(items, []int{5, 7})
	m := knapsackRecursiveInfItems(items, 7)
	assert.Equal(t, 9, m)
	// fmt.Printf("maxvalue=%d c=%d\n", m, c)
}

func TestKnapsackRecursiveItems1(t *testing.T) {
	wt := []int{6, 3, 4, 2}
	val := []int{30, 14, 16, 9}

	m := knapsackRecursiveItems(7, wt, val, len(wt))
	assert.Equal(t, 60, m)
	// fmt.Printf("maxvalue=%d c=%d\n", m, c)
}

func TestKnapsackRecursiveItems2(t *testing.T) {
	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}

	m := knapsackRecursiveItems(7, wt, val, len(wt))
	assert.Equal(t, 9, m)
	// fmt.Printf("maxvalue=%d c=%d\n", m, c)
}

func TestKnapsackTabulationItems1(t *testing.T) {
	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}

	m := knapsackBottomUp(7, wt, val)
	assert.Equal(t, 9, m)
	// fmt.Printf("maxvalue=%d c=%d\n", m, c)
}
