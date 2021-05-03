package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//PLACE N QUEENS ON BOARD SO THAT NO ONE CAN ATTACK THE OTHER

func TestNQueens1(t *testing.T) {
	r := nqueens(4, 4)
	assert.Equal(t, 4, len(r))
	fmt.Printf("%v\n", r)

	r = nqueens(8, 8)
	assert.Equal(t, 8, len(r))
	fmt.Printf("%v\n", r)

	r = nqueens(11, 11)
	assert.Equal(t, 11, len(r))
	fmt.Printf("%v\n", r)

}
