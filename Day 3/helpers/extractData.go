package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() string {
	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	var data string

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		// extract data from the line
		newLine := formatInput(line)

		// add data to the arrays
		data += newLine
	}

	defer f.Close()

	return data
}

/*
*

	Formats the string input, I just copied and pasted it straight from the advent of code link so no manual formatting required.
*/
func formatInput(x string) string {

	// replace the newline fields.
	x = strings.Replace(x, "\r", "", 1)
	x = strings.Replace(x, "\n", "", 1)

	return x
}
