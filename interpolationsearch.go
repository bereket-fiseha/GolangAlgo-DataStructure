package main

import "fmt"

func InterpolationSearch(arrayList []int, target int) bool {

	lowerBound := 0
	upperBound := len(arrayList) - 1

	for lowerBound <= upperBound {
		arrayLen := (upperBound - lowerBound)
		indexOfTarget := lowerBound + ((target-arrayList[lowerBound])*arrayLen)/(arrayList[upperBound]-arrayList[lowerBound])
		fmt.Println(lowerBound)
		if arrayList[indexOfTarget] == target {
			return true

		} else if target < arrayList[indexOfTarget] {
			upperBound = indexOfTarget - 1
		} else {
			lowerBound = indexOfTarget + 1

		}

	}
	return false
}
