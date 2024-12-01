package main

import (
	"day1/helpers"
	"fmt"
	"math"
)

func main() {
	fmt.Println("start")

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

	fmt.Println(totalDistance)

}
