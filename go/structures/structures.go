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
