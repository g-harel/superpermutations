# superpermutations

A superpermutation of the string `n` is another string that contains all the permutations of the characters in `n`. For example, `1221` is a superpermutation of `12`. However, `121` is a shorter superpermutation of `12` because it overlaps characters.

This repository contains code to produce superpermutations that are close to minimal (proving a superpermutation is minimal for any `n` is still not a solved problem).

The algorithm starts with the original string `n` and progressively appends more character(s). The number of appended/shifted characters each iteration is taken from a sequence of integers which is generated from the input string's length.

#### Shift sequences

| `len(n)` | sequence |
| --- | --- |
| 1 | `-` |
| 2 | `1` |
| 3 | `112` |
| 4 | `111211121113` |
| 5 | `111121111211112111131111211112111121111311112111121111211114` |
| ... | `...` |

When the shift is of 1, the first character of the most recent permutation can be appended to the end of the string. This adds a unique permutation of the string `n` to the end of the result string.

```
123     shift(1)
 231    shift(1)
  312   shift(1)
12312   ...
```

When the shift (`s`) is larger than one, the first `s` characters of the most recent permutation are mirrored and appended to the end of the result string. This process repeats until the sequence is exhausted, which means half + 1 of the result has been computed. The rest of the string is produced by mirroring this first part.

```
1234         shift(1)
 2341        shift(1)
  3412       shift(1)
     2143    shift(3) 341 -> 143
      1432   shift(1)
1234121432   ...
```

[more information about superpermutations](http://www.njohnston.ca/2013/04/the-minimal-superpermutation-problem/)

[more information about the integer sequences](https://oeis.org/A235748)

## Usage

```shell
$ go get -u github.com/g-harel/superpermutations/superpermutations
```

### CLI

```
Usage:
  superpermutations [flags]

Flags:
  -c, --check          check correctness of result (big performance hit)
  -h, --help           help for superpermutations
  -l, --length int     set input string length (max 13) (default 5)
  -o, --out string     write result to a file
  -p, --print          print the result (may be very large)
  -r, --runes string   custom list of chars (looped if < length)
  -s, --silent         silence all output (except --print)
      --version        version for superpermutations
```

### Package

```go
import "github.com/g-harel/superpermutations"

func main() {
  // generate superpermutation
  s := superpermutations.Find("01234")

  // confirm that it contains all permutations
  fmt.Println(superpermutations.Check("01234", s))
}
```