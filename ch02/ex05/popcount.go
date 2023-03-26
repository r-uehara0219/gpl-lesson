package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// ex) 10101 & 10100 = 10100
// ex) 10100 & 10011 = 10000
// ex) 10000 & 01111 = 00000
// => 10101 has 3 bits
func PopCountByClearing(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
