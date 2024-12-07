package main

import (
	"day6/helpers"
	"day6/overviewFuncs"
	"fmt"
)

func partOne(overview [][]string) {
	var guard overviewFuncs.Guard

	totalDistinctLocations := 0

	// search for guard's location and set
	guard = guard.Find(overview)

	outOfBounds := false

	// continue while inbounds
	for !outOfBounds {
		var distinctLocationVisited int
		// count the current tile
		overview, distinctLocationVisited = guard.CountTile(overview)

		totalDistinctLocations += distinctLocationVisited

		// check the next tile
		nextTile := guard.LookAhead(overview)

		// handle navigation
		switch nextTile {
		case "outofbounds":
			outOfBounds = true
		case "#":
			guard = guard.TurnRight()
		default:
			guard = guard.MoveForwards()
		}
	}

	fmt.Println("Part one answer: ")
	fmt.Println(totalDistinctLocations)
}

func partTwo(overview [][]string) {

	var guard overviewFuncs.Guard
	visitedTiles := make(overviewFuncs.VisitedTile)

	totalPossibleObstructions := 0

	outOfBounds := false

	// search for guard's location and set
	guard = guard.Find(overview)

	// continue while inbounds
	for !outOfBounds {

		// log current tile as visited
		visitedTiles[guard.GetDirection()] = append(visitedTiles[guard.GetDirection()], guard.GetLocation())

		if guard.SearchForLoops(overview, visitedTiles) {
			totalPossibleObstructions++
		}

		// check the next tile
		nextTile := guard.LookAhead(overview)

		// handle navigation
		switch nextTile {
		case "outofbounds":
			outOfBounds = true
		case "#":
			guard = guard.TurnRight()
		default:
			guard = guard.MoveForwards()
		}
	}

	fmt.Println("Part two result: ")
	fmt.Println(totalPossibleObstructions)
}

func main() {
	overview := helpers.ReadMap()

	// partOne(overview)

	partTwo(overview)
}
