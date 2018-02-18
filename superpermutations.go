package superpermutations

import (
	"fmt"
	"index/suffixarray"
	"math"
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
	// computing permutations
	// derived from https://codereview.stackexchange.com/a/101829/105607
	length := len(input)
	ps := make([][]byte, factorial(length))
	wg := distribute(len(ps), func(start, end int) {
		for i := start; i <= end; i++ {
			combination := i
			remainingBitmask := (1 << uint(length)) - 1
			ps[i] = make([]byte, length)

			for j := 0; j < length; j++ {
				whichNumber := combination / factorial(length-1-j)
				combination %= factorial(length - 1 - j)

				bits := remainingBitmask
				for whichNumber > 0 {
					bits -= (bits & -bits)
					whichNumber--
				}

				nextNum := trailingZeros(bits)
				remainingBitmask &= ^(1 << uint(nextNum))
				ps[i][j] = input[nextNum]
			}
		}
	})

	index := suffixarray.New([]byte(superpermutation))

	wg.Wait()

	// checking that all permutations are present in the index
	status := true
	distribute(len(ps), func(start, end int) {
		for _, p := range ps[start:end] {
			if len(index.Lookup(p, 1)) == 0 {
				status = false
			}
		}
	}).Wait()

	return status
}

// divides the values of count into almost equal buckets (max difference of 1)
// calls callback in goroutine for each bucket with start and end indecies
func distribute(count int, cb func(int, int)) *sync.WaitGroup {
	maxGoroutines := 32.00

	bucketCount := int(math.Min(maxGoroutines, float64(count)))
	bucketSize := count / bucketCount
	bucketDrop := count % bucketCount

	wg := sync.WaitGroup{}
	wg.Add(bucketCount)

	for i := 0; i < bucketCount; i++ {
		start := i*bucketSize + int(math.Min(float64(i), float64(bucketDrop)))
		end := start + bucketSize - 1
		if i < bucketDrop {
			end++
		}

		go func() {
			cb(start, end)
			wg.Done()
		}()
	}

	return &wg
}

// returns hardcoded factorials up to max int
func factorial(a int) int {
	f := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600, 6227020800, 87178291200, 1307674368000, 20922789888000, 355687428096000, 6402373705728000, 121645100408832000, 2432902008176640000}
	if a > len(f) {
		panic(fmt.Errorf("cannot compute factorial above %d", len(f)))
	}
	return f[a]
}

// threadsafe splice
func splice(input []byte, index int) []byte {
	subset := []byte{}
	subset = append(subset, input[:index]...)
	subset = append(subset, input[index+1:]...)
	return subset
}

// count trailing zeros in binary form
func trailingZeros(n int) int {
	if n == 0 {
		return 63
	}

	s := 0
	for (n & 1) == 0 {
		s++
		n >>= 1
	}

	return s
}
