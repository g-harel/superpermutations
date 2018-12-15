package superpermutations

import (
	"index/suffixarray"
	"math"
	"sync"
)

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
