package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//PRINT ALL COMBINATIONS OF K ELEMENTS FROM INPUT

func TestAllAnagrams1(t *testing.T) {
	a := allAnagrams("god")
	fmt.Printf("%v\n", a)
	assert.Equal(t, 6, len(a))
}
