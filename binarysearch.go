package main

import "fmt"

func BinarySearch(arrayList []int, target int) bool {

	lowerBound := 0
	upperBound := len(arrayList) - 1
	for i := 0; i < len(arrayList); i++ {
		fmt.Println(lowerBound)
		fmt.Println(upperBound)
		midIndex := (lowerBound + upperBound + 1) / 2

		if arrayList[midIndex] == target {

			return true

		}
		if target < arrayList[midIndex] {

			upperBound = midIndex - 1

		} else {
			lowerBound = midIndex + 1
		}
	}
	return false
}
