package main

import (
	"fmt"
	ds "leetcode/structures"
)

/**
 * To run in vscode:
 * ctrl+shift+d -> select GO Launch -> f5
 */

func main() {
	mHeap := &ds.MinHeap{}
	mHeap.Insert(1)
	mHeap.Insert(2)
	mHeap.Insert(3)
	mHeap.Insert(6)
	mHeap.Insert(0)
	mHeap.Insert(10)

	mHeap.PrintHeap()

	fmt.Println(mHeap.GetMin())
	mHeap.PrintHeap()
	fmt.Println(mHeap.GetMin())
	mHeap.PrintHeap()
	fmt.Println(mHeap.GetMin())
	mHeap.PrintHeap()
	fmt.Println(mHeap.GetMin())
	mHeap.PrintHeap()
	fmt.Println(mHeap.GetMin())
	mHeap.PrintHeap()
	fmt.Println(mHeap.GetMin())

	mHeap.PrintHeap()
}
