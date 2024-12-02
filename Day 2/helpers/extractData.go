package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	var reports [][]int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// extract data from the line
		report := formatInput(line)

		// add data to the arrays
		reports = append(reports, report)
	}

	defer f.Close()

	return reports
}

/*
*

	Formats the string input, I just copied and pasted it straight from the advent of code link so no manual formatting required.
*/
func formatInput(x string) []int {

	// replace the newline fields.
	x = strings.Replace(x, "\r", "", 1)
	x = strings.Replace(x, "\n", "", 1)

	// convert to array
	strArr := strings.Split(x, " ")
	var intArr []int
	var err error

	// populate the 2 val array with number versions
	// honestly I think this is inefficient, but it's neat.
	for _, val := range strArr {
		var intVal int
		intVal, err = strconv.Atoi(val)

		intArr = append(intArr, intVal)
		if err != nil {
			panic(err)
		}
	}

	return intArr
}
