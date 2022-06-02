package main

import "fmt"

func main() {
	//itemExists := LinearSearch([]int{8, 9, 3, 4}, 4)
	//itemExists := InterpolationSearch([]int{3, 10, 20, 22, 26, 33, 36, 38, 40, 100, 500}, 26)
	orginalArray := []int{6, 4, 5, 4, 7, 9, 1, 2}
	var array = []int{}
	array = append(array, orginalArray...)
	sortedArray := BubbleSort(array)

	fmt.Println(orginalArray)
	fmt.Println(sortedArray)
}
