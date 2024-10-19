package main

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestBits(t *testing.T) {
	var n uint64
	n = SetBit(n, 5)
	n = SetBit(n, 10)

	fmt.Println(IsBitOne(n, 1))
	fmt.Println(IsBitOne(n, 5))
	fmt.Println(IsBitOne(n, 10))
	fmt.Printf("%b\n", n)
	fmt.Println(bits.OnesCount64(n))
}

func TestIntersectionOfBitMap(t *testing.T) {
	min := 5
	bm1 := NewBitMap(min, []int{10, 15, 20})
	bm2 := NewBitMap(min, []int{10, 11, 20})

	fmt.Println(IntersectionOfBitMap(bm1, bm2, min))
}
