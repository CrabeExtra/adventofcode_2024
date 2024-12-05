package helpers

func MergeSortByRules(update []int, rules map[int][]int) []int {
	arrLen := len(update)

	var mergedResult []int

	if arrLen > 1 {
		halfLen := arrLen / 2
		leftSlice := update[:halfLen]
		rightSlice := update[halfLen:]

		leftResult := MergeSortByRules(leftSlice, rules)
		rightResult := MergeSortByRules(rightSlice, rules)

		mergedResult = mergeArraysByRules(leftResult, rightResult, rules)
	} else {
		return update
	}

	return mergedResult
}

func mergeArraysByRules(left []int, right []int, rules map[int][]int) []int {
	iL := 0
	iR := 0
	lenL := len(left)
	lenR := len(right)
	var result []int
	exitCond := false

	// merge each line until complete
	for !exitCond {
		// current vals
		var leftCurr int
		var rightCurr int
		// if no more values to iterate, arrays are merged.
		if iL >= lenL && iR >= lenR {
			exitCond = true

			// if left side reached the end, append with right side
		} else if iL >= lenL {
			rightCurr = right[iR]
			result = append(result, rightCurr)
			iR++

			// vice versa
		} else if iR >= lenR {
			leftCurr = left[iL]
			result = append(result, leftCurr)
			iL++
			// otherwise, there is data to compare
		} else {
			leftCurr = left[iL]
			rightCurr = right[iR]

			leftFirst := Includes(rules[leftCurr], rightCurr) // does left take precedent
			// if left first, add left first and increment iL
			if leftFirst {
				result = append(result, leftCurr)
				iL++

				// vice versa (simple else statement, prevents having to iterate the rules array again)
				// edge cases: neither first? then it doesn't matter. Both first? Also doesn't matter.
			} else {
				result = append(result, rightCurr)
				iR++
			}
		}
	}

	return result
}

func CheckForIncorrectSorting(updates [][]int, rules map[int][]int) [][]int {
	//loop through updates. for each update, iterate from len - 1, for each element, check for all other elements in rules.

	var incorrectUpdates [][]int

	for _, update := range updates {
		logUpdate := false
		for i1, entry1 := range update {
			for i2, entry2 := range update {
				if i2 < i1 {
					// for each entry1 find all earlier entry2's that can't be earlier
					if Includes(rules[entry1], entry2) {
						logUpdate = true
					}
				}
			}
		}
		if logUpdate {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	return incorrectUpdates
}
