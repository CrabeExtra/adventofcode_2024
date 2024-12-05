package main

import (
	"day5/helpers"
	"fmt"
)

func partOne(rules map[int][]int, updates [][]int) {
	sumMiddle := 0

	// loop through updates
	for _, update := range updates {
		// sort the update by rules
		mergedArr := helpers.MergeSortByRules(update, rules)

		// did the sort do anything?
		noChanges := helpers.SameArr(update, mergedArr)

		// same array? sum the middle val
		if noChanges {
			sumMiddle += update[(len(update) / 2)]
		}

	}

	fmt.Println("Part one result")
	fmt.Println(sumMiddle)

}

func partTwo(rules map[int][]int, updates [][]int) {
	sumMiddle := 0

	// loop through updates
	for _, update := range updates {
		// sort the update by rules
		mergedArr := helpers.MergeSortByRules(update, rules)

		// did the sort do anything?
		noChanges := helpers.SameArr(update, mergedArr)

		// sorting did something? sum the middle val
		if !noChanges {
			sumMiddle += mergedArr[(len(mergedArr) / 2)]

		}

	}

	// incorrectSorting := helpers.CheckForIncorrectSorting(allSorted, rules)
	// incorrectSortingTotal := helpers.CheckForIncorrectSorting(updates, rules)

	// fmt.Println("Part two result")
	// fmt.Println("Total updates: ")
	// fmt.Println(len(updates))
	// fmt.Println("Total incorrect updates: ")
	// fmt.Println(len(incorrectSortingTotal))
	// fmt.Println("Total sorted: ")
	// fmt.Println(len(allSorted))
	// fmt.Println("Total calculated incorrect after sorting: ")
	// fmt.Println(len(incorrectSorting))
	fmt.Println(sumMiddle)
}

func main() {

	rules := helpers.ReadRules()
	updates := helpers.ReadUpdates()

	partOne(rules, updates)

	partTwo(rules, updates)
}
