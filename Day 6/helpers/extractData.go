package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadMap() [][]string {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := bufio.NewReader(f)

	var data [][]string

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		lineArr := strings.Split(line, "")

		// add data to the arrays
		data = append(data, lineArr)
	}

	defer f.Close()

	return data
}
