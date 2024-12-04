package lib

// quick reference shorthand representation: [{rowPos: 1, colPos: 3}, ...]
type Stack []map[string]int

func (s Stack) Push(v map[string]int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, map[string]int) {

	if len(s) == 0 {
		return s, map[string]int{}
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func InBounds(row int, col int, arr [][]string) bool {
	if row >= 0 && col >= 0 && row < len(arr) && col < len(arr[0]) {
		return true
	}
	return false
}

func BackDiagonalMas(tlr int, tlc int, brr int, brc int, data [][]string) bool {
	if (data[tlr][tlc] == "S" && data[brr][brc] == "M") || (data[tlr][tlc] == "M" && data[brr][brc] == "S") {
		return true
	}
	return false
}

func ForeDiagonalMas(trr int, trc int, blr int, blc int, data [][]string) bool {
	if (data[trr][trc] == "S" && data[blr][blc] == "M") || (data[trr][trc] == "M" && data[blr][blc] == "S") {
		return true
	}
	return false
}
