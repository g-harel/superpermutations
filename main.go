package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(findSuperpermutation("1234567"))
}

func findSuperpermutation(value string) string {
	length := len(value)
	shifts := factorial(length-1)*length - 1
	sequence := make([]int, shifts)
	// populating the sequence with values.
	for i := 1; i < length; i++ {
		initial := int(math.Floor(float64(shifts) / float64(factorial(i+1))))
		interval := initial + 1
		for j := initial; j < shifts; j += interval {
			sequence[j]++
		}
	}
	sp := value
	for i := 0; i < shifts/2+1; i++ {
		sp += reverse(sp[len(sp)-length : len(sp)-length+sequence[i]])
	}
	sp = sp[:len(sp)-len(value)+1]
	ss := sp + reverse(sp)[1:]
	if isSuperpermutation(value, ss) {
		return ss
	}
	return "err"
}

func reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
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
