package main

import (
	"day2/helpers"
	"fmt"
	"math"
)

func partOne(reports [][]int) {
	fmt.Println("Part one answer: ")

	safeReports := 0

	for _, report := range reports {

		var ascending bool

		for index, element := range report {
			// only check for n-1 report elements
			if index < (len(report) - 1) {
				nextElement := report[index+1]
				signedDifference := nextElement - element
				difference := math.Abs(float64(signedDifference))
				// check for >3 && ==0 changes in elements
				if difference > 3 || difference == 0 {
					break // break from report and go to next report, no safe increment
				}

				// check for change in direction
				if index == 0 {
					ascending = true
					if signedDifference < 0 {
						ascending = false
					}
				} else {
					// case where prior ascending, but currently descending
					if ascending && signedDifference < 0 {
						break
						// case where prior descending, but current ascending
					} else if !ascending && signedDifference > 0 {
						break
					}
				}
				// last index reached? this must be a safe report.
			} else {
				safeReports++
			}

		}
	}

	fmt.Println("Safe Reports: ")
	fmt.Println(safeReports)
}

func checkIfReportSafe(report []int) bool {
	var ascending bool

	for index, element := range report {
		// only check for n-1 report elements
		if index < (len(report) - 1) {
			nextElement := report[index+1]
			signedDifference := nextElement - element
			difference := math.Abs(float64(signedDifference))
			// check for >3 && ==0 changes in elements
			if difference > 3 || difference == 0 {
				break // break from report and go to next report, no safe increment
			}

			// check for change in direction
			if index == 0 {
				ascending = true
				if signedDifference < 0 {
					ascending = false
				}
			} else {
				// case where prior ascending, but currently descending
				if ascending && signedDifference < 0 {
					break
					// case where prior descending, but current ascending
				} else if !ascending && signedDifference > 0 {
					break
				}
			}
			// last index reached? this must be a safe report.
		} else {
			return true
		}
	}
	return false
}

/*
*

	I started getting into the guts of it but it was tedious, so just doing a shitty brute force solution O(n*m^2).
	I guess it's fine because m is naturally not a large number.
*/
func partTwoBruteForce(reports [][]int) {
	fmt.Println("Part two answer: ")

	safeReports := 0

	for _, report := range reports {
		result := checkIfReportSafe(report)
		if result {
			safeReports++
		} else {
			for index := range report {
				newReport := helpers.DeleteIndexFromIntArray(report, index)
				if checkIfReportSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}

	fmt.Println("Safe Reports: ")
	fmt.Println(safeReports)
}

func main() {
	reports := helpers.ReadInput()

	partOne(reports)

	partTwoBruteForce(reports)
}
