package datastructures

import (
	"fmt"
	"math"
)

/**
 * To run in vscode:
 * ctrl+shift+d -> select GO Launch -> f5
 */

// MAX HEAP OR PRIORITY QUEUE AS MAX HEAP
type maxHeap struct {
	heap []int
}

func (mh *maxHeap) Insert(v int) {
	mh.heap = append(mh.heap, v)
	mh.bubbleUp(len(mh.heap) - 1)
}

func (mh *maxHeap) GetMax() int {
	if len(mh.heap) > 0 {
		currentMax := mh.heap[0]
		mh.swap(0, len(mh.heap)-1)
		mh.heap = append(mh.heap[:len(mh.heap)-1])
		mh.bubbleDown(0)
		return currentMax
	}
	return 0
}

func (mh *maxHeap) bubbleUp(idx int) {
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
func (mh *maxHeap) bubbleDown(idx int) {
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

func (mh *maxHeap) swap(idx1, idx2 int) {
	mh.heap[idx1], mh.heap[idx2] = mh.heap[idx2], mh.heap[idx1]
}

func (mh *maxHeap) getParentIDx(i int) int {
	if i != 0 {
		return int(math.Ceil(float64(i)/2) - 1)
	}
	return -1
}

func (mh *maxHeap) getLeftChildIDx(i int) int {
	index := i*2 + 1
	if index < len(mh.heap)-1 {
		return index
	}
	return 0
}
func (mh *maxHeap) getRightChildIDx(i int) int {
	index := i*2 + 2
	if index < len(mh.heap)-1 {
		return index
	}
	return 0
}
func (mh *maxHeap) PrintHeap() {
	fmt.Println(mh.heap)
}

// func main() {
// 	mHeap := &maxHeap{}
// 	mHeap.Insert(1)
// 	mHeap.Insert(2)
// 	mHeap.Insert(3)

// 	mHeap.PrintHeap()

// 	mHeap.Insert(6)
// 	mHeap.Insert(0)
// 	mHeap.Insert(10)

// 	mHeap.PrintHeap()

// 	fmt.Println(mHeap.GetMax())
// 	fmt.Println(mHeap.GetMax())
// 	fmt.Println(mHeap.GetMax())

// 	mHeap.PrintHeap()
// }

// MAX HEAP OR PRIORITY QUEUE AS MAX HEAP END
