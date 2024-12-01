package main

import (
	"day1/helpers"
	"fmt"
	"math"
)

/*
*
This solves the first part of the issue.
*/
func partOne() {
	fmt.Println("Part one answer:")

	// read in the input data
	input := helpers.ReadInput()

	// sort the lists of data
	input["setOne"] = helpers.Sort(input["setOne"])
	input["setTwo"] = helpers.Sort(input["setTwo"])

	totalDistance := 0

	// assuming the lists are the same size
	for index, element := range input["setOne"] {

		// extract the same index from the second set of data points
		setTwoPair := input["setTwo"][index]

		// calculate the distance between
		totalDistance += int(math.Abs(float64(element - setTwoPair)))

	}
	fmt.Println("Total distance:")
	fmt.Println(totalDistance)
}

/*
*

	This solves the second part of the issue (didn't spend more than a few seconds thinking up the complexity so don't judge)
*/
func partTwo() {
	fmt.Println("Part two answer:")

	// read in the input data.
	input := helpers.ReadInput()

	visitedNumbers := make(map[int]int)

	similarity := 0

	//iterate the first array.
	for _, element1 := range input["setOne"] {

		numOccurrences := visitedNumbers[element1]

		// check if number in map.
		if numOccurrences == 0 {
			// doesn't exist yet, find it.

			// count number of occurrences.
			for _, element2 := range input["setTwo"] {
				if element1 == element2 {
					numOccurrences++
				}
			}
			// add to map in case it occurs again.
			visitedNumbers[element1] = numOccurrences
		}
		// if already exists, increment again for repeated digit.
		// if doesn't already exist, increment for counted occurrences.
		similarity += numOccurrences * element1
	}

	fmt.Println("Similarity Score:")
	fmt.Println(similarity)

}

func main() {
	partOne()
	partTwo()
}
