package helpers

/*
*
	Deletes an index from an array and returns a new array.
*/
func DeleteIndexFromIntArray(arr []int, index int) []int {
	var newArr []int
	for i, elem := range arr {
		if i != index {
			newArr = append(newArr, elem)
		}
	}

	return newArr
}
