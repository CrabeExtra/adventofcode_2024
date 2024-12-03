package main

import (
	"day3/helpers"
	"fmt"
	"regexp"
	"strconv"
)

/*
*

	Goes through data, finds proper mul format and sums them.
*/
func partOne(data string) {
	fmt.Println("Part One: ")
	sum := 0

	reg := regexp.MustCompile(`mul\((\d{1,3}),\s*(\d{1,3})\)`)

	matches := reg.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		// Convert the first and second number from string to int
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		// Multiply the numbers and add to result
		product := num1 * num2
		sum += product
	}

	fmt.Println("Sum total: ")
	fmt.Println(sum)
}

func partTwo(data string) {
	fmt.Println("Part Two:")
	regMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	regDont := regexp.MustCompile(`don't\(\)`)
	regDo := regexp.MustCompile(`do\(\)`)

	ignoreMul := false
	sum := 0
	i := 0

	for i < len(data) {
		if regDont.MatchString(data[i:]) && data[i:i+len("don't()")] == "don't()" {
			ignoreMul = true
			i += len("don't()")
		} else if regDo.MatchString(data[i:]) && data[i:i+len("do()")] == "do()" {
			ignoreMul = false
			i += len("do()")
		} else if !ignoreMul {
			if loc := regMul.FindStringIndex(data[i:]); loc != nil && loc[0] == 0 {
				matched := data[i+loc[0] : i+loc[1]]
				matches := regMul.FindStringSubmatch(matched)
				num1, _ := strconv.Atoi(matches[1])
				num2, _ := strconv.Atoi(matches[2])
				sum += num1 * num2
				i += loc[1]
			} else {
				i++
			}
		} else {
			i++ // Mov to next char if ignoring
		}
	}

	fmt.Println("Sum total:", sum)
}

func main() {
	data := helpers.ReadInput()

	partOne(data)

	partTwo(data)
}
