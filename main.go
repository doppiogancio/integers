package integers

import (
	"fmt"
	"sort"
	"strings"
)

func Join(elems []int, separator string) string {
	listOfStrings := []string{}
	for _, value := range elems {
		listOfStrings = append(listOfStrings, fmt.Sprintf("%d", value))
	}
	return strings.Join(listOfStrings, separator)
}

func Find(haystack []int, needle int) int {
	if len(haystack) == 0 {
		return -1
	}

	for index, n := range haystack {
		if n == needle {
			return index
		}
	}
	return -1
}

func Contains(haystack []int, needle int) bool {
	return Find(haystack, needle) > -1
}

func MakeRange(min, max int) []int {
	a := []int{}

	if max < min {
		return a
	}

	for {
		a = append(a, min)
		min += 1

		if min > max {
			break
		}
	}

	return a
}

func MakeRangeWithStep(min, max, step int) []int {
	a := []int{}

	if max < min {
		return a
	}

	for {
		a = append(a, min)
		min += step

		if min > max {
			break
		}
	}

	return a
}

func RemoveIndex(array []int, index int) []int {
	if index > len(array) {
		return array
	}

	copy(array[index:], array[index+1:])
	array[len(array)-1] = 0
	array = array[:len(array)-1]
	return array
}

// RemoveValue will remove only the first occurrence of value
func RemoveValue(array []int, value int) []int {
	index := Find(array, value)
	if index < 0 {
		return array
	}

	return RemoveIndex(array, index)
}

func UniqueValues(values []int) []int {
	valueMap := map[int]int{}
	list := []int{}

	for _, value := range values {
		valueMap[value] = value
	}

	for _, value := range valueMap {
		list = append(list, value)
	}

	sort.Ints(list)

	return list
}

