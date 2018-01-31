package main

import (
	"fmt"
	"strings"
)

/*

1. create all permutations for input string of length n
2. for m between 1 and n-1 group permutations by overlap of n-m chars
3. solution has

> 1
1
# 1
= 1

> 12
12
 21
# 121
= 2-1^2

> 123
123
 231
  312
132
 321
  213
# 123121321
= (3^2-2^2)*2-1^2

> 1234
1234
 2341
  3412
   4123
    -
     2314
      3142
       1423
        4231
         -
          3124
           1243
            2431
             4312
1324
 3241
  2413
   4132
    -
     3214
      2143
       1432
        4321
         -
          2134
           1342
            3421
             4213
# 123412314231243121342132413214321
= ((4^2-3^2)*3-2^2)*2-1^2

> 12345
#
= ((((5)*5-4^2)*4-3^2)*3-2^2)*2-1^2

> 123456
# 12345612345162345126345123645132645136245136425136452136451234651234156234152634152364152346152341652341256341253641253461253416253412653412356412354612354162354126354123654132654312645316243516243156243165243162543162453164253146253142653142563142536142531645231465231456231452631452361452316453216453126435126431526431256432156423154623154263154236154231654231564213564215362415362145362154362153462135462134562134652134625134621536421563421653421635421634521634251634215643251643256143256413256431265432165432615342613542613452613425613426513426153246513246531246351246315246312546321546325146325416325461325463124563214563241563245163245613245631246532146532416532461532641532614532615432651436251436521435621435261435216435214635214365124361524361254361245361243561243651423561423516423514623514263514236514326541362541365241356241352641352461352416352413654213654123
= (((((6)*6-5^2)*5-4^2)*4-3^2)*3-2^2)*2-1^2

*/

func main() {
	val := "1234"
	res := findSuperpermutation(val)
	count := 0
	for _, v := range res {
		if isSuperpermutation(val, v) {
			count++
			fmt.Println(len(v), v)
		}
	}
	fmt.Println(count, len(res))
}

func findSuperpermutation(input string) []string {
	p := permutations(strings.Split(input, ""))
	for i := 1; i < len(input); i++ {
		temp := []string{}
		for j := 0; j < len(p); j++ {
			a := p[j]
			for k := 0; k < len(p); k++ {
				if j == k {
					continue
				}
				b := p[k]
				startA, startB, maxLen := longestSubstr(a, b)
				if maxLen >= len(input)-i {
					a = roll(a, -startA) + roll(b, startB)[:maxLen]
				}
			}
			temp = append(temp, a)
		}
		p = temp
	}
	return p
}

func roll(s string, amount int) string {
	amount = (len(s) - amount%len(s)) % len(s)
	res := (s + s)[amount : amount+len(s)]
	return res
}

func longestSubstr(a, b string) (startA, startB, maxLen int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			increment := 0
			for a[i+increment] == b[j+increment] {
				increment++
				if i+increment == len(a) {
					break
				}
				if j+increment == len(b) {
					break
				}
			}
			if increment > maxLen {
				startA = i
				startB = j
				maxLen = increment
			}
		}
	}
	return
}

func splice(arr *[]string, index int) string {
	val := (*arr)[index]
	(*arr) = append((*arr)[:index], (*arr)[index+1:]...)
	return val
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

// analysis

func isSuperpermutation(input, guess string) bool {
	for _, permutation := range permutations(strings.Split(input, "")) {
		if strings.Index(guess, permutation) == -1 {
			return false
		}
	}
	return true
}

type match struct {
	visits int
	valid  bool
}

func superpermutationStats(input, guess string) {
	if !isSuperpermutation(input, guess) {
		return
	}
	stats := map[string]*match{}
	for _, permutation := range permutations(strings.Split(input, "")) {
		stats[permutation] = &match{0, true}
	}
	useless := 0
	for i := 0; i <= len(guess)-len(input); i++ {
		substr := guess[i : i+len(input)]
		if stats[substr] == nil {
			stats[substr] = &match{0, false}
		}
		stats[substr].visits++
		if stats[substr].valid != true {
			if stats[substr].visits > 1 {
				fmt.Println(substr, stats[substr].visits)
			}
			useless++
		}
	}
	fmt.Println(useless)
}
