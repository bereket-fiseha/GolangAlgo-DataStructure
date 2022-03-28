package main

import "fmt"

func main() {
	arrayList := []int{5, 10, 7, 4, 9}
	devide(arrayList)

}

func devide(arrayList []int) {
	arrayLen := len(arrayList)
	mid := arrayLen / 2
	leftSideList := arrayList[0:mid]
	rightSideList := arrayList[mid:arrayLen]
	fmt.Println(leftSideList)
	fmt.Println(rightSideList)
	if len(leftSideList) > 1 {
		devide(leftSideList)
	}
	if len(rightSideList) > 1 {
		devide(rightSideList)
	}

}
