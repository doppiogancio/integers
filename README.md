Package integers [![GoDoc](https://godoc.org/github.com/doppiogancio/integers?status.png)](https://godoc.org/github.com/doppiogancio/integers)
=======

Package integers implements simple functions to manipulate integer numbers

Examples:
```
for _, x := range integers.MakeRange(1, 9) {
		...
}

for _, evenNumber := range integers.MakeRangeWithStep(0, 20, 2) {
		...
}

integers.Join([]int{1,2,3}, "-") // 1-2-3

integers.Find([]Find{11,22,33}, 33) // 2

integers.Find([]Find{11,22,33}, 44) // -1

integers.Contains([]Find{11,22,33}, 33) // true

integers.Contains([]Find{11,22,33}, 44) // false

integers.RemoveIndex([]Find{11,22,33}, 1) // [11,33]

integers.RemoveValue([]Find{11,22,33}, 22) // [11,33]

integers.UniqueValues([]Find{11,22,33,22,22,11,33,33,33}) // [11,22,33]
```
