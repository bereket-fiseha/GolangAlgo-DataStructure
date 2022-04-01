package main

import "fmt"

func main() {
	//itemExists := LinearSearch([]int{8, 9, 3, 4}, 4)
	itemExists := BinarySearch([]int{3, 4, 5, 6, 7}, 7)

	fmt.Println(itemExists)

}
