package main

func BinarySearchRecursive(arrayList []int, target int) bool {
	if len(arrayList) == 1 {
		if arrayList[0] == target {
			return true
		}
	} else {
		indexMidEl := len(arrayList) / 2
		midEl := arrayList[indexMidEl]
		if target == midEl {
			return true
		} else if target < midEl {
			return BinarySearchRecursive(arrayList[:indexMidEl], target)

		} else {

			return BinarySearchRecursive(arrayList[indexMidEl+1:], target)
		}

	}
	return false

}
