package main

func LinearSearch(arrayList []int, target int) bool {

	for _, el := range arrayList {
		if el == target {
			return true

		}
	}
	return false

}
