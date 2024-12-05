package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadRules() map[int][]int {
	f, err := os.Open("rules.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	data := make(map[int][]int)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		rulesArr := strings.Split(line, "|")
		var rulesArrI []int

		for _, elem := range rulesArr {
			rulesArrI = append(rulesArrI, strToInt(elem))
		}

		data[rulesArrI[0]] = append(data[rulesArrI[0]], rulesArrI[1])
	}

	defer f.Close()

	return data
}

func ReadUpdates() [][]int {
	f, err := os.Open("updates.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	var data [][]int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		updatesArr := strings.Split(line, ",")

		var intUpdates []int

		for _, elem := range updatesArr {
			intUpdates = append(intUpdates, strToInt(elem))
		}
		// add data to the arrays
		data = append(data, intUpdates)
	}

	defer f.Close()

	return data
}

/*
*

	Formats the string input, I just copied and pasted it straight from the advent of code link so no manual formatting required.
*/
func strToInt(x string) int {
	x = strings.ReplaceAll(x, "\r\n", "")
	newInt, err := strconv.Atoi(x)

	if err != nil {
		return -1
	}
	return newInt
}
