Package integers [![GoDoc](https://godoc.org/github.com/doppiogancio/integers?status.png)](https://godoc.org/github.com/doppiogancio/integers)
=======

Package integers implements simple functions to manipulate integer numbers

A list of functions included in this package:
```
Join(elems []int, separator string) string
Find(haystack []int, needle int) int
Contains(haystack []int, needle int) bool
MakeRange(min, max int) []int
MakeRangeWithStep(min, max, step int) []int
RemoveIndex(array []int, index int) []int
RemoveValue(array []int, value int) []int
UniqueValues(values []int) []int
```