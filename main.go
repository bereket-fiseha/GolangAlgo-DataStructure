package main

import "fmt"

func main() {
	//itemExists := LinearSearch([]int{8, 9, 3, 4}, 4)
	itemExists := InterpolationSearch([]int{3, 10, 20, 22, 26, 33, 36, 38, 40, 100, 500}, 26)

	fmt.Println(itemExists)

}
