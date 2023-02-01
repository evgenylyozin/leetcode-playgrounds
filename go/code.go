package main

/**
 * To run in vscode:
 * ctrl+shift+d -> select GO Launch -> f5
 */

import (
	"fmt"
	algs "leetcode/algorithms"
)

func main() {
	arr := []int{3, 2, 4, 1, 0}
	algs.MergeSort(arr, 0, 4)
	fmt.Println(arr) // => [0,1,2,3,4]
}
