# superpermutations

A superpermutation of the string `n` is another string that contains all the permutations of the characters in `n`. For example, `2112` is a superpermutation of `12`. However, `121` is a shorter superpermutation of `12` because it leverages overlapping characters to shorten the total length. This repository contains code to produce superpermutations that are close to minimal or minimal (proving a superpermutation is minimal for any `n` is still an open problem).

This implementation functions by starting with the original string `n` and appending the next character(s). The process starts by generating a sequence of integers that represent how many characters need to be shifted at any point.

##### Shift sequences

| `len(n)` | sequence |
| --- | --- |
| 1 | 1 |
| 2 | 121 |
| 3 | 11211 |
| 4 | 11121112111311121112111 |
| 5 | 11112111121111211113111121111211112111131111211112111121111411112111121111211113111121111211112111131111211112111121111 |
| ... | ... |

When the shift is of 1, the first character of the most recent permutation can be appended to the end of the string. This adds a unique permutation of the string `n` to the last characters of the result string.

```
123     shift(1)
 231    shift(1)
  312   shift(1)
12312   ...
```

When the shift is of 2, the first two characters of the most recent permutation need to be mirrored before being appended to the end of the string to avoid repeating permutations.

```
123        shift(1)
 231       shift(1)
  312      shift(1)
    213    shift(2)
     132   shift(1)
12312132   ...
```

The generalisation of this process for shifts of size larger than 1 is to branch out recursively for all permutations of the shifted characters and to keep iterating through the shift sequence on each branch until a duplicate permutation of the original `n` is found or until the sequence is exhausted.

```
2431                                           ...
 4312                                          shift(1)
    2134---2143---2314---2341---2413---2431    shift(3)
     1342   1432   3142   3412   4132   4312   shift(1)
```

[more information about superpermutations](http://www.njohnston.ca/2013/04/the-minimal-superpermutation-problem/)

## Running

```shell
$ go run *.go
```