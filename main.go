package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(findSuperpermutation("12345"))
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
	// finding a superpermutation
	superpermutation := findNext(value, length, 0, &sequence)
	fmt.Println("|")
	// sanity check
	if isSuperpermutation(value, superpermutation) {
		return superpermutation
	}
	return "err"
}

func findNext(s string, length, index int, sequence *[]int) string {
	// visual progress feedback
	if rand.Float32() < 0.001 {
		fmt.Print("*")
	}
	if index >= len(*sequence) {
		return s
	}
	offset := (*sequence)[index]
	newstr := s[len(s)-length : len(s)-length+offset]
	var options []string
	if offset < 4 || index < len(*sequence)/2-1 {
		options = []string{reverse(newstr)}
	} else {
		options = permutations(strings.Split(newstr, ""))[1:]
	}
	for _, o := range options {
		if strings.Index(s, s[len(s)-length+offset:]+o) != -1 {
			continue
		}
		guess := findNext(s+o, length, index+1, sequence)
		if guess != "" {
			return guess
		}
	}
	return ""
}
