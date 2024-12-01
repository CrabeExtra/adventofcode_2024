package helpers

import (
	"sort"
)

/**
Literally just a wrapper around integer sorting.
*/

func Sort(inputArr []int) []int {
	sort.Ints(inputArr[:])
	return inputArr
}
