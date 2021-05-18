package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//PRINT ONLY WORDS FROM STRING THAT ARE PRESENT IN DICTIONARY

func TestDictString1(t *testing.T) {
	a := dictString("catsanddog", []string{"cat", "cats", "dog", "and", "sand"})
	fmt.Printf("%v\n", a)
	assert.Equal(t, 2, len(a))

	// a = dictString("aabcdfcatsanddogaabdfds", []string{"cat", "cats", "dog", "and", "sand"})
	// fmt.Printf("%v\n", a)
	// assert.Equal(t, 2, len(a))

}
