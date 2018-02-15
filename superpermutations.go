package superpermutations

import (
	"index/suffixarray"
	"math"
	"strings"
	"sync"
)

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

// Check verifies that the the second argument is a superpermutation of the first.
func Check(input, superpermutation string) bool {
	ps := permutations(strings.Split(input, ""))

	// seperating permutations into buckets
	bucketCount := int(math.Min(32.0, float64(len(ps))))
	bucketSize := len(ps) / bucketCount
	bucketDrop := len(ps) % bucketCount
	buckets := make([][]string, bucketCount)
	for i := 0; i < bucketCount; i++ {
		start := i*bucketSize + int(math.Min(float64(i), float64(bucketDrop)))
		end := start + bucketSize - 1
		if i < bucketDrop {
			end++
		}
		buckets = append(buckets, ps[start:end])
	}

	index := suffixarray.New([]byte(superpermutation))

	// checking bucket that all bucket values are in the index
	status := true
	wg := sync.WaitGroup{}
	for _, b := range buckets {
		subset := b
		wg.Add(1)
		go func() {
			for _, p := range subset {
				if len(index.Lookup([]byte(p), 1)) == 0 {
					status = false
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return status
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

// TODO improve perf
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
