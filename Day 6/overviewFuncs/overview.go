package overviewFuncs

type Guard struct {
	direction string
	location  [2]int
}

func (guard Guard) GetDirection() string {
	return guard.direction
}

func (guard Guard) GetLocation() [2]int {
	return guard.location
}

/*
*

	The guard moves forwards.
*/
func (guard Guard) MoveForwards() Guard {
	switch guard.direction {
	case "left":
		guard.location[1] -= 1
	case "right":
		guard.location[1] += 1
	case "down":
		guard.location[0] += 1
	case "up":
		guard.location[0] -= 1
	}

	return guard
}

/*
*

	The guard turns right 90 degrees.
*/
func (guard Guard) TurnRight() Guard {
	switch guard.direction {
	case "left":
		guard.direction = "up"
	case "right":
		guard.direction = "down"
	case "down":
		guard.direction = "left"
	case "up":
		guard.direction = "right"
	}

	return guard
}

/*
*

	Counts current guard tile if applicable. Returns a modified map and the new count.
*/
func (guard Guard) CountTile(overview [][]string) ([][]string, int) {
	row := guard.location[0]
	col := guard.location[1]
	count := 0
	if overview[row][col] == "." || overview[row][col] == "^" || overview[row][col] == "<" || overview[row][col] == ">" || overview[row][col] == "v" {
		overview[row][col] = "X"
		count++
	}

	return overview, count
}

/*
*

	The guard looks ahead.
*/
func (guard Guard) LookAhead(overview [][]string) string {
	row := guard.location[0]
	col := guard.location[1]
	nextRow := row
	nextCol := col

	switch guard.direction {
	case "left":
		nextCol = col - 1
		// return overview[row][col-1]
	case "right":
		nextCol = col + 1
	case "down":
		nextRow = row + 1
	case "up":
		nextRow = row - 1
	}

	height := len(overview)

	if nextRow > height-1 || nextRow < 0 || row > height-1 || row < 0 {
		return "outofbounds"
	}

	width := len(overview[row])

	if nextCol < 0 || nextCol > width-1 || col < 0 || col > width-1 {
		return "outofbounds"
	}

	return overview[nextRow][nextCol]
}

func (guard Guard) Find(overview [][]string) Guard {
	for rowInd, row := range overview {
		if guard.direction != "" {
			break
		}
		for colInd, loc := range row {
			if guard.direction != "" {
				break
			}

			guard = guard.SetLocation(rowInd, colInd)
			switch loc {
			case "^":
				guard = guard.SetDirection("up")
			case ">":
				guard = guard.SetDirection("right")
			case "<":
				guard = guard.SetDirection("left")
			case "v":
				guard = guard.SetDirection("down")
			}
		} // Is there a downwards caret? is this relevant?
	}

	return guard
}

func (guard Guard) SetLocation(rowIndex int, colIndex int) Guard {
	guard.location[0] = rowIndex
	guard.location[1] = colIndex

	return guard
}

func (guard Guard) SetDirection(dir string) Guard {
	guard.direction = dir

	return guard
}

// if direction left, check upwards for visited tiles in the same direction
// if direction right, check downwards for visited tiles in the same direction
// if direction up, check right for "
// if direction down, check left for "
type VisitedTile map[string][][2]int

/*
*

	Essentially just check for possible loops when placing obstructions.

	This impl is kind of messy, but feeling lazy.
*/
func (guard Guard) SearchForLoops(overview [][]string, visitedTiles VisitedTile) bool {

	switch guard.direction {
	case "left":
		for index := range overview {
			if index < guard.location[0] {
				for _, elem := range visitedTiles["up"] {
					// if there is a visited tile in the up direction, then a loop may be formed.
					if elem[0] == index && elem[1] == guard.location[1] {
						return true
					}
				}

			}
		}
	case "right":
		for index := range overview {
			if index > guard.location[0] {
				for _, elem := range visitedTiles["down"] {
					// if there is a visited tile in the up direction, then a loop may be formed.
					if elem[0] == index && elem[1] == guard.location[1] {
						return true
					}
				}
			}
		}
	case "down":
		for index := range overview[guard.location[0]] {
			if index < guard.location[1] {
				for _, elem := range visitedTiles["left"] {
					// if there is a visited tile in the down direction, then a loop may be formed.
					if elem[0] == guard.location[0] && elem[1] == index {
						return true
					}
				}

			}
		}
	case "up":
		for index := range overview[guard.location[0]] {
			if index > guard.location[1] {
				for _, elem := range visitedTiles["right"] {
					// if there is a visited tile in the down direction, then a loop may be formed.
					if elem[0] == guard.location[0] && elem[1] == index {
						return true
					}
				}
			}
		}
	}
	return false
}
