package main

import (
	"strings"
)

func isSuperpermutation(input, guess string) bool {
	for _, permutation := range permutations(strings.Split(input, "")) {
		if strings.Index(guess, permutation) == -1 {
			return false
		}
	}
	return true
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

func reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
