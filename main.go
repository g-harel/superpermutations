package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(findSuperpermutation("1234567"))
}

func findSuperpermutation(value string) string {
	valLength := len(value)
	rolls := []int{}
	rollsLength := factorial(valLength-1)*valLength - 1
	for i := 0; i < rollsLength; i++ {
		rolls = append(rolls, 0)
	}
	for i := 1; i < valLength; i++ {
		initial := int(math.Floor(float64(rollsLength) / float64(factorial(i+1))))
		interval := initial + 1
		for j := initial; j < rollsLength; j += interval {
			rolls[j]++
		}
	}
	superpermutation := findNext(value, valLength, 0, &rolls)
	fmt.Println("|")
	if isSuperpermutation(value, superpermutation) {
		return superpermutation
	}
	return "err"
}

func findNext(s string, initialLength, index int, rolls *[]int) string {
	if rand.Float32() < 0.001 {
		fmt.Print("*")
	}
	if index >= len(*rolls) {
		return s
	}
	offset := (*rolls)[index]
	newstr := s[len(s)-initialLength : len(s)-initialLength+offset]
	var options []string
	if offset < 4 || index < len(*rolls)/2-1 {
		options = []string{reverse(newstr)}
	} else {
		options = permutations(strings.Split(newstr, ""))[1:]
	}
	for _, o := range options {
		if strings.Index(s, s[len(s)-initialLength+offset:]+o) != -1 {
			continue
		}
		guess := findNext(s+o, initialLength, index+1, rolls)
		if guess != "" {
			return guess
		}
	}
	return ""
}
