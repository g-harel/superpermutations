package main

import (
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {
	var check bool
	var length int
	var print bool
	var write string

	rootCmd := &cobra.Command{
		Use: "superpermutations",
		Run: func(cmd *cobra.Command, args []string) {
			cli(check, length, print, write)
		},
	}

	rootCmd.PersistentFlags().BoolVar(&check, "check", false, "check correctness of result")
	rootCmd.PersistentFlags().BoolVar(&print, "print", false, "print the result (may be very large)")
	rootCmd.PersistentFlags().IntVar(&length, "length", 5, "set input string length (max 16)")
	rootCmd.PersistentFlags().StringVar(&write, "write", "", "write result to a file")

	rootCmd.Execute()
}

func cli(check bool, length int, print bool, write string) {
	min := 0
	max := 13
	if length <= min {
		color.Red("Error: length must be bigger than %d\n", min)
		return
	} else if length > max {
		color.Red("Error: lengths above %d are not supported (maximum slice size)\n", max)
		return
	}

	chars := "0123456789abcdef"
	value := ""
	for i := 0; i < length; i++ {
		value += string(chars[i])
	}

	color.White("Computing for length %d ...", length)

	sp := Find(value)

	if print {
		color.Magenta(sp)
	}

	color.Cyan("Found, size: %d chars\n", len(sp))

	if check {
		color.White("Checking ...")
		if isSuperpermutation(value, sp) {
			color.Cyan("Check has passed!")
		}
	}

	if write != "" {
		color.White("Writing ...")
		err := ioutil.WriteFile(write, []byte(sp), 0644)
		if err != nil {
			color.Red("Error: could not write to file \"%v\"\n", write)
		} else {
			color.Cyan("Written successfully to \"%v\"", write)
		}
	}
}

// Find computes a superpermutation of the input string.
func Find(value string) string {
	length := len(value)
	shifts := factorial(length) / 2
	sequence := make([]int, shifts)

	// populating the shift sequence with values.
	for i := 2; i <= length; i++ {
		initial := 2 * (shifts - 1) / factorial(i)
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
		for i := 0; i < inc; i++ {
			out[cur+inc-i-1] = out[cur-length+i]
			out[outlen-cur-inc+i] = out[cur-length+i]
		}
		cur += inc
	}

	return string(out)
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
