package structures

import (
	"fmt"
	"math"
)

/**
 * Just some useful data structures in go
 */

// MaxHeap OR PRIORITY QUEUE AS MAX HEAP
type MaxHeap struct {
	heap []int
}

// Insert ...
func (mh *MaxHeap) Insert(v int) {
	mh.heap = append(mh.heap, v)
	mh.bubbleUp(len(mh.heap) - 1)
}

// GetMax ...
func (mh *MaxHeap) GetMax() int {
	if len(mh.heap) > 0 {
		currentMax := mh.heap[0]
		mh.swap(0, len(mh.heap)-1)
		mh.heap = append(mh.heap[:len(mh.heap)-1])
		if len(mh.heap) > 0 {
			mh.bubbleDown(0)
		}
		return currentMax
	}
	return 0
}

func (mh *MaxHeap) bubbleUp(idx int) {
	inserted := mh.heap[idx]
	for idx > 0 {
		parentIDx := mh.getParentIDx(idx)
		if inserted > mh.heap[parentIDx] {
			mh.swap(parentIDx, idx)
			idx = parentIDx
			continue
		}
		break
	}
}
func (mh *MaxHeap) bubbleDown(idx int) {
	inserted := mh.heap[idx]
	for idx < len(mh.heap)-1 {
		leftChildIDx := mh.getLeftChildIDx(idx)
		rightChildIDx := mh.getRightChildIDx(idx)
		validLCIDx := leftChildIDx > 0
		validRCIDx := rightChildIDx > 0

		if (!validRCIDx && validLCIDx) || (validLCIDx && validRCIDx && mh.heap[leftChildIDx] >= mh.heap[rightChildIDx]) {
			if inserted < mh.heap[leftChildIDx] {
				mh.swap(idx, leftChildIDx)
				idx = leftChildIDx
				continue
			}
		}
		if validRCIDx {
			if inserted < mh.heap[rightChildIDx] {
				mh.swap(idx, rightChildIDx)
				idx = rightChildIDx
				continue
			}
		}
		break
	}
}

func (mh *MaxHeap) swap(idx1, idx2 int) {
	mh.heap[idx1], mh.heap[idx2] = mh.heap[idx2], mh.heap[idx1]
}

func (mh *MaxHeap) getParentIDx(i int) int {
	if i != 0 {
		return int(math.Ceil(float64(i)/2) - 1)
	}
	return -1
}

func (mh *MaxHeap) getLeftChildIDx(i int) int {
	index := i*2 + 1
	if index < len(mh.heap) {
		return index
	}
	return 0
}

func (mh *MaxHeap) getRightChildIDx(i int) int {
	index := i*2 + 2
	if index < len(mh.heap) {
		return index
	}
	return 0
}

// PrintHeap ...
func (mh *MaxHeap) PrintHeap() {
	fmt.Println(mh.heap)
}

// import (
// 	"fmt"
// 	ds "leetcode/structures"
// )

// func main() {
// 	mHeap := &ds.MaxHeap{}
// 	mHeap.Insert(1)
// 	mHeap.Insert(2)
// 	mHeap.Insert(3)
// 	mHeap.Insert(6)
// 	mHeap.Insert(0)
// 	mHeap.Insert(10)

// 	mHeap.PrintHeap()

// 	fmt.Println(mHeap.GetMax())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMax())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMax())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMax())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMax())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMax())

// 	mHeap.PrintHeap()
// }

// MAX HEAP OR PRIORITY QUEUE AS MAX HEAP END
// ------------------------------------------------------------------------------

// MinHeap OR REVERSE PRIORITY QUEUE AS MIN HEAP
type MinHeap struct {
	heap []int
}

// Insert inserts an int into MinHeap and bubbles it up if necessery
func (mh *MinHeap) Insert(v int) {
	mh.heap = append(mh.heap, v)
	mh.bubbleUp(len(mh.heap) - 1)
}

// GetMin removes top item, sets last item to top and bubbles it down if necessery
func (mh *MinHeap) GetMin() int {
	if len(mh.heap) > 0 {
		currentMin := mh.heap[0]
		mh.swap(0, len(mh.heap)-1)
		mh.heap = append(mh.heap[:len(mh.heap)-1])
		if len(mh.heap) > 0 {
			mh.bubbleDown(0)
		}
		return currentMin
	}
	return 0
}

func (mh *MinHeap) bubbleUp(idx int) {
	inserted := mh.heap[idx]
	for idx > 0 {
		parentIDx := mh.getParentIDx(idx)
		if inserted < mh.heap[parentIDx] {
			mh.swap(parentIDx, idx)
			idx = parentIDx
			continue
		}
		break
	}
}
func (mh *MinHeap) bubbleDown(idx int) {
	inserted := mh.heap[idx]
	for idx < len(mh.heap)-1 {
		leftChildIDx := mh.getLeftChildIDx(idx)
		rightChildIDx := mh.getRightChildIDx(idx)

		validLCIDx := leftChildIDx > 0
		validRCIDx := rightChildIDx > 0

		if (!validRCIDx && validLCIDx) || (validLCIDx && validRCIDx && mh.heap[leftChildIDx] <= mh.heap[rightChildIDx]) {
			if inserted > mh.heap[leftChildIDx] {
				mh.swap(idx, leftChildIDx)
				idx = leftChildIDx
				continue
			}
		}
		if validRCIDx {
			if inserted > mh.heap[rightChildIDx] {
				mh.swap(idx, rightChildIDx)
				idx = rightChildIDx
				continue
			}
		}
		break
	}
}

func (mh *MinHeap) swap(idx1, idx2 int) {
	mh.heap[idx1], mh.heap[idx2] = mh.heap[idx2], mh.heap[idx1]
}

func (mh *MinHeap) getParentIDx(i int) int {
	if i != 0 {
		return int(math.Ceil(float64(i)/2) - 1)
	}
	return -1
}

func (mh *MinHeap) getLeftChildIDx(i int) int {
	index := i*2 + 1
	if index < len(mh.heap) {
		return index
	}
	return 0
}
func (mh *MinHeap) getRightChildIDx(i int) int {
	index := i*2 + 2
	if index < len(mh.heap) {
		return index
	}
	return 0
}

// PrintHeap prints current heap
func (mh *MinHeap) PrintHeap() {
	fmt.Println(mh.heap)
}

// import (
// 	"fmt"
// 	ds "leetcode/structures"
// )

// func main() {
// 	mHeap := &ds.MinHeap{}
// 	mHeap.Insert(1)
// 	mHeap.Insert(2)
// 	mHeap.Insert(3)
// 	mHeap.Insert(6)
// 	mHeap.Insert(0)
// 	mHeap.Insert(10)

// 	mHeap.PrintHeap()

// 	fmt.Println(mHeap.GetMin())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMin())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMin())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMin())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMin())
// 	mHeap.PrintHeap()
// 	fmt.Println(mHeap.GetMin())

// 	mHeap.PrintHeap()
// }

// MIN HEAP OR REVERSE PRIORITY QUEUE AS MIN HEAP END
// ------------------------------------------------------------------------------

// Binary Search Tree with traversal and search

// BSTNode ...
type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

// NewBSTNode creates a new bst node
func NewBSTNode(data int) *BSTNode {
	return &BSTNode{data: data}
}

// BST ...
type BST struct {
	root *BSTNode
}

// NewBST creates a new binary search tree
func NewBST() *BST {
	return &BST{}
}

// Insert data into the BST
func (bst *BST) Insert(data int) {
	newNode := NewBSTNode(data)
	if bst.root == nil {
		bst.root = newNode
	} else {
		bst.insertNode(bst.root, newNode)
	}
}

func (bst *BST) insertNode(node, newNode *BSTNode) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			bst.insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			bst.insertNode(node.right, newNode)
		}
	}
}

// Remove node with specific data
func (bst *BST) Remove(data int) {
	bst.root = bst.removeNode(bst.root, data)
}

func (bst *BST) removeNode(node *BSTNode, data int) *BSTNode {
	if node == nil {
		return nil
	}

	if data < node.data {
		node.left = bst.removeNode(node.left, data)
		return node
	}

	if data > node.data {
		node.right = bst.removeNode(node.right, data)
		return node
	}

	// if current node is the one to be deleted
	if node.left == nil && node.right == nil {
		node = nil
		return node
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	// Deleting node with two children
	aux := bst.findMinNode(node.right)
	node.data = aux.data

	node.right = bst.removeNode(node.right, aux.data)
	return node
}

// Inorder bst traversal
func (bst *BST) Inorder(node *BSTNode) {
	if node != nil {
		bst.Inorder(node.left)
		fmt.Println(node.data)
		bst.Inorder(node.right)
	}
}

// Preorder bst traversal
func (bst *BST) Preorder(node *BSTNode) {
	if node != nil {
		fmt.Println(node.data)
		bst.Inorder(node.left)
		bst.Inorder(node.right)
	}
}

// Postorder bst traversal
func (bst *BST) Postorder(node *BSTNode) {
	if node != nil {
		bst.Inorder(node.left)
		bst.Inorder(node.right)
		fmt.Println(node.data)
	}
}

// Inorder bst traversal
func (bst *BST) findMinNode(node *BSTNode) *BSTNode {
	if node.left == nil {
		return node
	}
	return bst.findMinNode(node.left)
}

// GetRootNode returns root node of the BST
func (bst *BST) GetRootNode() *BSTNode {
	return bst.root
}

// Search returns node with specified data or nil
func (bst *BST) Search(node *BSTNode, data int) *BSTNode {
	if node == nil {
		return nil
	}

	if data < node.data {
		return bst.Search(node.left, data)
	}

	if data > node.data {
		return bst.Search(node.right, data)
	}

	return node
}

// import (
// 	"fmt"
// 	ds "leetcode/structures"
// )

// func main() {
// 	bst := ds.NewBST()
// 	bst.Insert(40)
// 	bst.Insert(29)
// 	bst.Insert(47)
// 	bst.Insert(12)
// 	bst.Insert(41)
// 	bst.Insert(98)
// 	bst.Insert(1)
// 	bst.Insert(14)
// 	bst.Insert(28)
// 	bst.Insert(57)
// 	bst.Insert(52)

// 	bst.Inorder(bst.GetRootNode())

// 	fmt.Println(bst.Search(bst.GetRootNode(), 1))

// 	bst.Remove(12)
// 	bst.Remove(40)

// 	bst.Inorder(bst.GetRootNode())
// }

// Binary Search Tree with traversal and search END
//------------------------------------------------------------------------------
