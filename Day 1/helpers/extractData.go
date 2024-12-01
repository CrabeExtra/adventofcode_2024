package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() map[string][]int {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	var setOne []int
	var setTwo []int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// extract data from the line
		lineData := formatInput(line)

		// add data to the arrays
		setOne = append(setOne, lineData[0])
		setTwo = append(setTwo, lineData[1])
	}

	defer f.Close()

	m := map[string][]int{
		"setOne": setOne,
		"setTwo": setTwo,
	}

	return m
}

/*
*

	Formats the string input, I just copied and pasted it straight from the advent of code link so no manual formatting required.
*/
func formatInput(x string) [2]int {
	// remove the 3 spaces between the two numbers
	x = strings.Replace(x, "   ", " ", 1)

	// replace the newline fields.
	x = strings.Replace(x, "\r", "", 1)
	x = strings.Replace(x, "\n", "", 1)

	// convert to array
	strArr := strings.Split(x, " ")
	var intArr [2]int
	var err error

	// populate the 2 val array with number versions
	// honestly I think this is inefficient, but it's neat.
	for index, val := range strArr {
		intArr[index], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}

	return intArr
}
