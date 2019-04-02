package main

import (
	"fmt"
	"strconv"
)

// NumGen generates space padded, two digit numbers.
type NumGen struct {
	count int
}

func NewNumGen() NumGen {
	return NumGen{count: 0}
}

// Next generates the next number in line.  If the number
// has one digit it will be padded with a leading space.  Any numbers with
// two digits and above will be returned as is.  This implementation is specifically
// suited for numbers with 1 or 2 digits.
func (n *NumGen) Next() string {
	n.count += 1
	sCount := strconv.Itoa(n.count)
	if len(sCount) < 2 {
		return fmt.Sprintf("% d", n.count)
	} else {
		return string(sCount)
	}
}
