package superpermutations

import (
	"fmt"
	"index/suffixarray"
	"math"
	"strings"
	"sync"
	"time"
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
	start := time.Now()
	ps := permutations(input)

	fmt.Println("permut", time.Now().Sub(start))

	// calculating bucket indecies
	bucketCount := int(math.Min(32.0, float64(len(ps))))
	buckets := make([][2]int, bucketCount)
	bucketSize := len(ps) / bucketCount
	bucketDrop := len(ps) % bucketCount
	for i := 0; i < bucketCount; i++ {
		start := i*bucketSize + int(math.Min(float64(i), float64(bucketDrop)))
		end := start + bucketSize - 1
		if i < bucketDrop {
			end++
		}
		buckets = append(buckets, [2]int{start, end})
	}

	fmt.Println("bucket", time.Now().Sub(start))

	index := suffixarray.New([]byte(superpermutation))

	fmt.Println("index ", time.Now().Sub(start))

	// checking that all permutations are present in the index
	status := true
	wg := sync.WaitGroup{}
	for _, b := range buckets {
		indecies := b
		wg.Add(1)
		go func() {
			for _, p := range ps[indecies[0]:indecies[1]] {
				if len(index.Lookup([]byte(p), 1)) == 0 {
					status = false
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("done  ", time.Now().Sub(start))

	return status
}

func factorial(a int) int {
	f := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600, 6227020800, 87178291200, 1307674368000, 20922789888000, 355687428096000, 6402373705728000, 121645100408832000, 2432902008176640000}

	if a > len(f) {
		panic(fmt.Errorf("cannot compute factorial above %d", len(f)))
	}

	return f[a]
}

// TODO improve perf
func permutations(input string) []string {
	var pps func([]string) []string

	pps = func(input []string) []string {
		if len(input) == 1 {
			return input
		}
		p := []string{}
		for i, char := range input {
			subset := []string{}
			subset = append(subset, input[:i]...)
			subset = append(subset, input[i+1:]...)
			for _, s := range pps(subset) {
				p = append(p, char+s)
			}
		}
		return p
	}

	return pps(strings.Split(input, ""))
}
