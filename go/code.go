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
	fmt.Println(algs.SelectionSort([]int{3, 2, 4, 1, 0})) // => [0,1,2,3,4]
}
