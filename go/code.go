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
	fmt.Println(algs.BinarySearchMeta([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))  // => 2
	fmt.Println(algs.BinarySearchMeta([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // => -1
}
