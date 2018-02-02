package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	findSuperpermutations("123456")
}

func findSuperpermutations(value string) string {
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
	guesses := findNext(value, valLength, 0, rolls)
	fmt.Println("|")
	for _, guess := range guesses {
		if isSuperpermutation(value, guess) {
			fmt.Println("found", guess)
		}
	}
	return ""
}

func findNext(s string, initialLength, index int, rolls []int) []string {
	if index >= len(rolls) {
		if rand.Float32() < 0.001 {
			fmt.Print("*")
		}
		return []string{s}
	}
	offset := rolls[index]
	newstr := s[len(s)-initialLength : len(s)-initialLength+offset]
	var options []string
	if offset < 4 || index < len(rolls)/2-1 {
		options = []string{reverse(newstr)}
	} else {
		options = permutations(strings.Split(newstr, ""))[1:]
	}
	var res []string
	for _, o := range options {
		for _, m := range findNext(s+o, initialLength, index+1, rolls) {
			res = append(res, m)
		}
	}
	return res
}
