package algorithms

import "math"

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

// TernarySearchIterative - iterative version of ternary search
func TernarySearchIterative(arr []int, whatToSearchFor int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == whatToSearchFor {
			return mid1
		}

		if arr[mid2] == whatToSearchFor {
			return mid2
		}

		if arr[mid1] > whatToSearchFor {
			right = mid1 - 1
			continue
		}
		if arr[mid2] < whatToSearchFor {
			left = mid2 + 1
			continue
		}
		left = mid1 + 1
		right = mid2 - 1
	}
	return -1
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.TernarySearchIterative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.TernarySearchIterative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }

// TernarySearchRecursive - recursive version of ternary search
func TernarySearchRecursive(arr []int, whatToSearchFor, left, right int) int {
	if left > right {
		return -1
	}

	mid1 := left + (right-left)/3
	mid2 := right - (right-left)/3

	if arr[mid1] == whatToSearchFor {
		return mid1
	}
	if arr[mid2] == whatToSearchFor {
		return mid2
	}

	if arr[mid1] > whatToSearchFor {
		return TernarySearchRecursive(arr, whatToSearchFor, left, mid1-1)
	}

	if arr[mid2] < whatToSearchFor {
		return TernarySearchRecursive(arr, whatToSearchFor, mid2+1, right)
	}

	return TernarySearchRecursive(arr, whatToSearchFor, mid1+1, mid2-1)
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.TernarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3,0,8))  // => 2
// 	fmt.Println(algs.TernarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10,0,8)) // => -1
// }

// JumpSearch - search on sorted arrays with specified jumps
func JumpSearch(arr []int, whatToSearchFor int) int {
	length := len(arr)
	step := int(math.Sqrt(float64(length)))
	start := 0
	end := step
	for arr[int(math.Min(float64(length), float64(end)))-1] < whatToSearchFor {
		start, end = end, end+step
		if start >= length {
			return -1
		}
	}

	for arr[start] < whatToSearchFor {
		start++
		if start == int(math.Min(float64(length), float64(end))) {
			return -1
		}
	}

	if arr[start] == whatToSearchFor {
		return start
	}
	return -1
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.JumpSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.JumpSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }

// ExponentialSearch - determine the range then use binary search on it
func ExponentialSearch(arr []int, whatToSearchFor int) int {
	if arr[0] == whatToSearchFor {
		return 0
	}
	i := 1
	for i < len(arr) && arr[i] <= whatToSearchFor {
		i = i * 2
	}
	return BinarySearchRecursive(arr, whatToSearchFor, i/2, int(math.Min(float64(i), float64(len(arr)-1))))
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.ExponentialSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.ExponentialSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }

// FibonacciSearch - search in ranges defined by fibonacci sequence
func FibonacciSearch(arr []int, whatToSearchFor int) int {
	fibMm2 := 0
	fibMm1 := 1
	fibM := fibMm2 + fibMm1

	for fibM < len(arr) {
		fibMm2, fibMm1 = fibMm1, fibM
		fibM = fibMm2 + fibMm1
	}

	offset := -1

	for fibM > 1 {
		i := int(math.Min(float64(fibMm2+offset), float64(len(arr)-1)))

		if arr[i] < whatToSearchFor {
			fibM = fibMm1
			fibMm1 = fibMm2
			fibMm2 = fibM - fibMm1
			offset = i
			continue
		}
		if arr[i] > whatToSearchFor {
			fibM = fibMm2
			fibMm1 = fibMm1 - fibMm2
			fibMm2 = fibM - fibMm1
			continue
		}
		return i
	}
	if fibMm1 > 0 && arr[len(arr)-1] == whatToSearchFor {
		return len(arr) - 1
	}
	return -1
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.FibonacciSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
// 	fmt.Println(algs.FibonacciSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
// }

// BubbleSort ...
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// import (
// 	"fmt"
// 	algs "leetcode/algorithms"
// )

// func main() {
// 	fmt.Println(algs.BubbleSort([]int{3, 2, 4, 1, 0})) // => [0,1,2,3,4]
// }
