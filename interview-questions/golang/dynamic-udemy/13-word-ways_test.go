package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//HOUSE

func TestWordWays1(t *testing.T) {

	w := []string{"pine", "pen", "apple", "applepen", "pineapple"}

	// s := time.Now()
	// for a := 0; a < 100000; a++ {
	r := wordWaysRecur("pineapplepenapple", w)
	assert.Equal(t, 3, r)
	// }
	// fmt.Printf("111 time=%d\n", time.Since(s))

	// s = time.Now()
	// for a := 0; a < 100000; a++ {
	r = wordWaysTab("pineapplepenapple", w)
	assert.Equal(t, 3, r)
	// }
	// fmt.Printf("222 time=%d\n", time.Since(s))

}
