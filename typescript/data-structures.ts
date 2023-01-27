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
      if (this.heap.length > 0) {
        this.bubbleDown(0);
      }
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
    if (index < this.heap.length) {
      return index;
    } else {
      return null;
    }
  }
  private getRightChildIDx(i: number) {
    const index = i * 2 + 2;
    if (index < this.heap.length) {
      return index;
    } else {
      return null;
    }
  }
  printHeap() {
    console.log(this.heap.toString());
  }
}

// const maxHeap = new MaxHeap();

// maxHeap.insert(1);
// maxHeap.insert(2);
// maxHeap.insert(3);

// maxHeap.printHeap();

// maxHeap.insert(6);
// maxHeap.insert(0);
// maxHeap.insert(10);

// maxHeap.printHeap();

// console.log(maxHeap.getMax());
// maxHeap.printHeap();
// console.log(maxHeap.getMax());
// maxHeap.printHeap();
// console.log(maxHeap.getMax());

// maxHeap.printHeap();

// Max Heap or Priority queue with Max Heap END
// -----------------------------------------------------------------------------

// Max Heap or Priority queue with Max Heap

class MinHeap {
  private heap: number[] = [];
  insert(v: number) {
    this.heap.push(v);
    this.bubbleUp(this.heap.length - 1);
  }
  getMin() {
    if (this.heap.length > 0) {
      const currentMin = this.heap[0];
      this.swap(0, this.heap.length - 1);
      this.heap.pop();
      if (this.heap.length > 0) {
        this.bubbleDown(0);
      }
      return currentMin;
    } else {
      return 0;
    }
  }
  private bubbleUp(idx: number) {
    const inserted = this.heap[idx];
    while (idx > 0) {
      const parentIDx = this.getParentIDx(idx) as number;
      if (inserted < this.heap[parentIDx]) {
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
        (leftChildIDx && leftChild && rightChild && leftChild <= rightChild)
      ) {
        if (inserted > this.heap[leftChildIDx]) {
          this.swap(idx, leftChildIDx);
          idx = leftChildIDx;
          continue;
        }
      }
      if (rightChildIDx) {
        if (inserted > this.heap[rightChildIDx]) {
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
    if (index < this.heap.length) {
      return index;
    } else {
      return null;
    }
  }
  private getRightChildIDx(i: number) {
    const index = i * 2 + 2;
    if (index < this.heap.length) {
      return index;
    } else {
      return null;
    }
  }
  printHeap() {
    console.log(this.heap.toString());
  }
}

// const minHeap = new MinHeap();

// minHeap.insert(1);
// minHeap.insert(2);
// minHeap.insert(3);

// minHeap.printHeap();

// minHeap.insert(6);
// minHeap.insert(0);
// minHeap.insert(10);

// minHeap.printHeap();

// console.log(minHeap.getMin());
// minHeap.printHeap();
// console.log(minHeap.getMin());
// minHeap.printHeap();
// console.log(minHeap.getMin());
// minHeap.printHeap();
// console.log(minHeap.getMin());
// minHeap.printHeap();
// console.log(minHeap.getMin());
// minHeap.printHeap();
// console.log(minHeap.getMin());
// minHeap.printHeap();

// Min Heap or reverse Priority queue with Min Heap END
//------------------------------------------------------------------------------

// Binary Search Tree with traversal and search

class BSTNode {
  data: number;
  left: null | BSTNode;
  right: null | BSTNode;
  constructor(data: number) {
    this.data = data;
    this.left = null;
    this.right = null;
  }
}

class BST {
  private root: null | BSTNode;
  constructor() {
    this.root = null;
  }

  Insert(data: number) {
    const newNode = new BSTNode(data);
    if (this.root === null) {
      this.root = newNode;
    } else {
      this.insertNode(this.root, newNode);
    }
  }

  private insertNode(node: BSTNode, newNode: BSTNode) {
    if (newNode.data < node.data) {
      if (node.left === null) {
        node.left = newNode;
      } else {
        this.insertNode(node.left, newNode);
      }
    } else {
      if (node.right === null) {
        node.right = newNode;
      } else {
        this.insertNode(node.right, newNode);
      }
    }
  }

  Remove(data: number) {
    this.root = this.removeNode(this.root, data);
  }

  private removeNode(node: BSTNode | null, data: number): BSTNode | null {
    if (node === null) return null;

    if (data < node.data) {
      node.left = this.removeNode(node.left, data);
      return node;
    }

    if (data > node.data) {
      node.right = this.removeNode(node.right, data);
      return node;
    }

    // if current node is the one to be deleted

    if (node.left === null && node.right === null) {
      node = null;
      return node;
    }

    if (node.left === null) {
      node = node.right;
      return node;
    }

    if (node.right === null) {
      node = node.left;
      return node;
    }

    // Deleting node with two children
    const aux = this.findMinNode(node.right);
    node.data = aux.data;

    node.right = this.removeNode(node.right, aux.data);
    return node;
  }

  Inorder(node: BSTNode | null) {
    if (node !== null) {
      this.Inorder(node.left);
      console.log(node.data);
      this.Inorder(node.right);
    }
  }

  Preorder(node: BSTNode | null) {
    if (node !== null) {
      console.log(node.data);
      this.Preorder(node.left);
      this.Preorder(node.right);
    }
  }

  Postorder(node: BSTNode | null) {
    if (node !== null) {
      this.Postorder(node.left);
      this.Postorder(node.right);
      console.log(node.data);
    }
  }

  private findMinNode(node: BSTNode): BSTNode {
    if (node.left === null) return node;
    return this.findMinNode(node.left);
  }

  GetRootNode() {
    return this.root;
  }

  Search(node: BSTNode | null, data: number): BSTNode | null {
    if (node === null) return null;

    if (data < node.data) return this.Search(node.left, data);

    if (data > node.data) return this.Search(node.right, data);

    return node;
  }
}

// const bst = new BST();

// bst.Insert(40);
// bst.Insert(29);
// bst.Insert(47);
// bst.Insert(12);
// bst.Insert(41);
// bst.Insert(98);
// bst.Insert(1);
// bst.Insert(14);
// bst.Insert(28);
// bst.Insert(57);
// bst.Insert(52);

// bst.Inorder(bst.GetRootNode());

// console.log(bst.Search(bst.GetRootNode(), 1));

// bst.Remove(12);
// bst.Remove(40);

// bst.Inorder(bst.GetRootNode());

// Binary Search Tree with traversal and search END
//------------------------------------------------------------------------------
