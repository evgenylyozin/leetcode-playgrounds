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
	bst := ds.NewBST()
	bst.Insert(40)
	bst.Insert(29)
	bst.Insert(47)
	bst.Insert(12)
	bst.Insert(41)
	bst.Insert(98)
	bst.Insert(1)
	bst.Insert(14)
	bst.Insert(28)
	bst.Insert(57)
	bst.Insert(52)

	bst.Inorder(bst.GetRootNode())

	fmt.Println(bst.Search(bst.GetRootNode(), 1))

	bst.Remove(12)
	bst.Remove(40)

	bst.Inorder(bst.GetRootNode())
}
