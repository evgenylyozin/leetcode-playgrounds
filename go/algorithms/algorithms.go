package algorithms

// BinarySearchIterative - iterative version of binary search
func BinarySearchIterative(arr []int, whatToSearchFor int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (right + left) / 2
		if arr[middle] == whatToSearchFor {
			return middle
		}
		if arr[middle] > whatToSearchFor {
			right = middle - 1
			continue
		}
		left = middle + 1
	}
	return -1
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.BinarySearchIterative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.BinarySearchIterative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }

// BinarySearchRecursive - recursive version of binary search
func BinarySearchRecursive(arr []int, whatToSearchFor, left, right int) int {
	if left > right {
		return -1
	}
	middle := (right + left) / 2
	if arr[middle] == whatToSearchFor {
		return middle
	}
	if arr[middle] > whatToSearchFor {
		return BinarySearchRecursive(arr, whatToSearchFor, left, middle-1)
	}
	return BinarySearchRecursive(arr, whatToSearchFor, middle+1, right)
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 0, 8))  // => 2
// 	fmt.Println(algs.BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0, 8)) // => -1
// }

// BinarySearchMeta - Meta Binary Search (one sided binary search)
func BinarySearchMeta(arr []int, whatToSearchFor int) int {
	bitsToStoreArrLength := 0
	for (1 << bitsToStoreArrLength) < len(arr)-1 {
		bitsToStoreArrLength++
	}

	pos := 0
	for i := bitsToStoreArrLength; i >= 0; i-- {
		if arr[pos] == whatToSearchFor {
			return pos
		}

		newPos := pos | (1 << i)

		if newPos < len(arr) && arr[newPos] <= whatToSearchFor {
			pos = newPos
		}
	}
	if arr[pos] == whatToSearchFor {
		return pos
	}
	return -1
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.BinarySearchMeta([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.BinarySearchMeta([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }
