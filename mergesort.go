package main

// func main() {
// 	//arrayList := []int{5, 10, 7, 4, 9}
// 	//	devide(arrayList)
// 	//	mergedArray := mergeTwoSortedArrays([]int{1, 3, 7, 9}, []int{2, 4, 5, 10, 12, 14})
// 	//	fmt.Println(mergedArray)
// 	mergeSort([]int{4, 3, 2, 8, 9, 7})

// }

func mergeTwoSortedArrays(firstSortedArray []int, secondSortedArray []int) []int {
	mergedSortedArray := []int{}
	temp := 0
	for i := 0; i < len(firstSortedArray); i++ {
		for j := temp; j < len(secondSortedArray); j++ {

			if firstSortedArray[i] < secondSortedArray[j] {
				mergedSortedArray = append(mergedSortedArray, firstSortedArray[i])
				if i == len(firstSortedArray)-1 {
					mergedSortedArray = append(mergedSortedArray, secondSortedArray[j:]...)
				}

				break
			} else {

				mergedSortedArray = append(mergedSortedArray, secondSortedArray[j])
				temp += 1
				if j == len(secondSortedArray)-1 {
					mergedSortedArray = append(mergedSortedArray, firstSortedArray[i:]...)
				}
			}

		}
	}
	return mergedSortedArray
}
func mergeSort(arrayList []int) []int {

	if len(arrayList) <= 1 {

		return arrayList
	}
	arrayLen := len(arrayList)
	mid := arrayLen / 2
	leftSideList := arrayList[:mid]
	rightSideList := arrayList[mid:]

	//fmt.Println(leftSideList)
	//fmt.Println(rightSideList)

	mergeSort(leftSideList)
	mergeSort(rightSideList)
	return mergeTwoSortedArrays(leftSideList, rightSideList)

}
