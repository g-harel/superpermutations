package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findSuperpermutation("12345678"))
}

func findSuperpermutation(value string) string {
	length := len(value)
	shifts := factorial(length) / 2
	sequence := make([]int, shifts)

	// populating the shift sequence with values.
	for i := 2; i <= length; i++ {
		initial := int(2 * (shifts - 1) / factorial(i))
		interval := initial + 1
		for j := initial; j < shifts; j += interval {
			sequence[j]++
		}
	}

	// creating an empty output array
	outlen := 0
	for i := 1; i <= length; i++ {
		outlen += factorial(i)
	}
	out := make([]rune, outlen)

	// adding initial values to output array
	for i, r := range value {
		out[i] = r
		out[outlen-i-1] = r
	}

	// adding all remaining values to the output
	cur := length
	for _, inc := range sequence {
		newchars := out[cur-length : cur-length+inc]
		for j, c := range newchars {
			out[cur+inc-j-1] = c
			out[outlen-cur-inc+j] = c
		}
		cur += inc
	}

	// sanity check
	sp := string(out)
	if isSuperpermutation(value, sp) {
		return sp
	}
	return "err"
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	b := a
	for i := 2; i < a; i++ {
		b *= i
	}
	return b
}

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
