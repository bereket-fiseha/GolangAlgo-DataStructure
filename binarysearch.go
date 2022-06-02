package main

func BinarySearch(arrayList []int, target int) (bool, int) {

	lowerBound := 0
	upperBound := len(arrayList) - 1
	for lowerBound <= upperBound {
		midIndex := (lowerBound + upperBound) / 2

		if arrayList[midIndex] == target {

			return true, midIndex

		}
		if target < arrayList[midIndex] {

			upperBound = midIndex - 1

		} else {
			lowerBound = midIndex + 1
		}
	}
	return false, -1
}
