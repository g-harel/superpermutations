package main

import (
	"fmt"
	"strings"
)

// TODO update readme
// TODO length (number) from cli + make check an option
// TODO limit size
func main() {
	value := "12345678"
	sp := findSuperpermutation(value)
	if isSuperpermutation(value, sp) {
		fmt.Println(sp)
	}
}

func findSuperpermutation(value string) string {
	length := len(value)
	shifts := factorial(length) / 2
	sequence := make([]int, shifts)

	// populating the shift sequence with values.
	for i := 2; i <= length; i++ {
		initial := uint64(2 * (shifts - 1) / factorial(i))
		interval := initial + 1
		for j := uint64(initial); j < shifts; j += interval {
			sequence[j]++
		}
	}

	// creating an empty output array
	outlen := uint64(0)
	for i := 1; i <= length; i++ {
		outlen += factorial(i)
	}
	out := make([]rune, outlen)

	// adding initial values to output array
	for i, r := range value {
		out[i] = r
		out[outlen-uint64(i+1)] = r
	}

	// adding all remaining values to the output
	cur := length
	for _, inc := range sequence {
		for i := 0; i < inc; i++ {
			out[cur+inc-i-1] = out[cur-length+i]
			out[outlen-uint64(cur+inc-i)] = out[cur-length+i]
		}
		cur += inc
	}

	// sanity check
	return string(out)
}

func factorial(a int) uint64 {
	if a == 0 {
		return 1
	}
	b := uint64(a)
	for i := 2; i < a; i++ {
		b *= uint64(i)
	}
	return b
}

// TODO imporve perf (goroutines?)
func isSuperpermutation(input, guess string) bool {
	for _, permutation := range permutations(strings.Split(input, "")) {
		if strings.Index(guess, permutation) == -1 {
			return false
		}
	}
	return true
}

func permutations(input []string) []string {
	if len(input) == 1 {
		return input
	}
	p := []string{}
	for i, char := range input {
		subset := []string{}
		subset = append(subset, input[:i]...)
		subset = append(subset, input[i+1:]...)
		for _, s := range permutations(subset) {
			p = append(p, char+s)
		}
	}
	return p
}
