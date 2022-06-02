package main

// BubbleSort works by repeatedly swapping the adjacent elements if they are in the wrong order.

func BubbleSort(inputArray []int) []int {

	if len(inputArray) <= 1 {
		return inputArray

	}
	for i := 0; i < len(inputArray); i++ {
		for j := 0; j < len(inputArray)-1; j++ {
			if inputArray[j] > inputArray[j+1] {

				temp := inputArray[j+1]
				inputArray[j+1] = inputArray[j]
				inputArray[j] = temp

			}
		}
	}
	return inputArray
}
