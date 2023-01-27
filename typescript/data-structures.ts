/**
 * Just some useful data structures
 */

// Max Heap or Priority queue with Max Heap

class MaxHeap {
  private heap: number[] = [];
  insert(v: number) {
    this.heap.push(v);
    this.bubbleUp(this.heap.length - 1);
  }
  getMax() {
    if (this.heap.length > 0) {
      const currentMax = this.heap[0];
      this.swap(0, this.heap.length - 1);
      this.heap.pop();
      this.bubbleDown(0);
      return currentMax;
    } else {
      return 0;
    }
  }
  private bubbleUp(idx: number) {
    const inserted = this.heap[idx];
    while (idx > 0) {
      const parentIDx = this.getParentIDx(idx) as number;
      if (inserted > this.heap[parentIDx]) {
        this.swap(parentIDx, idx);
        idx = parentIDx;
        continue;
      }
      break;
    }
  }
  private bubbleDown(idx: number) {
    const inserted = this.heap[idx];
    while (idx < this.heap.length - 1) {
      const leftChildIDx = this.getLeftChildIDx(idx);
      const rightChildIDx = this.getRightChildIDx(idx);

      let leftChild = null;
      let rightChild = null;

      if (leftChildIDx && rightChildIDx) {
        leftChild = this.heap[leftChildIDx];
        rightChild = this.heap[rightChildIDx];
      }

      if (
        (!rightChildIDx && leftChildIDx) ||
        (leftChildIDx && leftChild && rightChild && leftChild >= rightChild)
      ) {
        if (inserted < this.heap[leftChildIDx]) {
          this.swap(idx, leftChildIDx);
          idx = leftChildIDx;
          continue;
        }
      }
      if (rightChildIDx) {
        if (inserted < this.heap[rightChildIDx]) {
          this.swap(idx, rightChildIDx);
          idx = rightChildIDx;
          continue;
        }
      }
      break;
    }
  }

  private swap(idx1: number, idx2: number) {
    [this.heap[idx1], this.heap[idx2]] = [this.heap[idx2], this.heap[idx1]];
  }

  private getParentIDx(i: number) {
    if (i !== 0) {
      return Math.ceil(i / 2) - 1;
    } else {
      return null;
    }
  }
  private getLeftChildIDx(i: number) {
    const index = i * 2 + 1;
    if (index < this.heap.length - 1) {
      return index;
    } else {
      return null;
    }
  }
  private getRightChildIDx(i: number) {
    const index = i * 2 + 2;
    if (index < this.heap.length - 1) {
      return index;
    } else {
      return null;
    }
  }
  printHeap() {
    console.log(this.heap.toString());
  }
}

const maxHeap = new MaxHeap();

maxHeap.insert(1);
maxHeap.insert(2);
maxHeap.insert(3);

maxHeap.printHeap();

maxHeap.insert(6);
maxHeap.insert(0);
maxHeap.insert(10);

maxHeap.printHeap();

console.log(maxHeap.getMax());
maxHeap.printHeap();
console.log(maxHeap.getMax());
maxHeap.printHeap();
console.log(maxHeap.getMax());

maxHeap.printHeap();
