package helpers

func Includes(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func SameArr(arr1 []int, arr2 []int) bool {
	len1 := len(arr1)
	len2 := len(arr2)

	if len1 != len2 {
		return false
	}

	for index, elem := range arr1 {
		if elem != arr2[index] {
			return false
		}
	}

	return true
}
