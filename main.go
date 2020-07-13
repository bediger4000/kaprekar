package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	k, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	orig := k
	if k >= 10000 {
		fmt.Printf("%d too many digits\n", k)
		return
	}
	n := 0
	for k != 6174 {
		k = kaprekar(k)
		n++
		if n > 100 {
			break
		}
	}
	fmt.Printf("%d %d\n", orig, n)
}

// kaprekar For a given input x, create two new numbers that consist of the digits in x in ascending and descending order.
// Subtract the smaller number from the larger number.
func kaprekar(x int64) int64 {
	var digits []int64
	for x > 0 {
		d := x % 10
		digits = append(digits, d)
		x /= 10
	}

	sort.Sort(SliceFwd(digits))
	var n1 int64
	var n2 int64
	l := len(digits) - 1
	for i := range digits {
		n1 *= 10
		n1 += digits[i]

		n2 *= 10
		n2 += digits[l-i]

	}
	r := n1 - n2
	if n2 > n1 {
		r = n2 - n1
	}
	return r
}

type SliceFwd []int64

func (s SliceFwd) Len() int { return len(s) }
func (s SliceFwd) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}
func (s SliceFwd) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type SliceBkwd []int64

func (s SliceBkwd) Len() int { return len(s) }
func (s SliceBkwd) Less(i, j int) bool {
	if s[i] < s[j] {
		return false
	}
	return true
}
func (s SliceBkwd) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
