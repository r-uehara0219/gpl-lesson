package main

import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("test"))
	c2 := sha256.Sum256([]byte("Test"))

	println(countDiffBits(c1, c2))
}

// 2つのハッシュ値の異なるビット数を数える
func countDiffBits(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		// XOR
		count += popCount(c1[i] ^ c2[i])
	}
	return count
}

func popCount(x byte) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
