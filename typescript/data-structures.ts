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

// Directed graph

class Vertex {
  key: number;
  adjacent: Vertex[] = [];
  constructor(key: number) {
    this.key = key;
  }
}

class DGraph {
  vertices: Vertex[] = [];

  contains(vertices: Vertex[], key: number) {
    for (let v of vertices) {
      if (v.key === key) return true;
    }
    return false;
  }
  getVertex(key: number) {
    for (let v of this.vertices) {
      if (v.key === key) {
        return v;
      }
    }
  }

  getVertexIndex(vertices: Vertex[], vertex: Vertex) {
    for (let i = 0; i < vertices.length; i++) {
      if (vertices[i] === vertex) {
        return i;
      }
    }
    return -1;
  }

  AddVertex(key: number) {
    if (this.contains(this.vertices, key)) return;
    const newVertex = new Vertex(key);
    this.vertices.push(newVertex);
  }

  AddEdge(to: number, from: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);

    if (
      !toVertex ||
      !fromVertex ||
      this.contains(fromVertex.adjacent, toVertex.key)
    )
      return;
    fromVertex.adjacent.push(toVertex);
  }

  RemoveVertex(key: number) {
    const vertexToRemove = this.getVertex(key);
    if (vertexToRemove) {
      this.removePointersToVertex(vertexToRemove);
      this.removeVertexItself(vertexToRemove);
    }
  }

  removePointersToVertex(vertex: Vertex) {
    for (let v of this.vertices) {
      for (let i = 0; i < v.adjacent.length; i++) {
        const adj = v.adjacent[i];
        if (adj == vertex) {
          v.adjacent.splice(i, 1);
        }
      }
    }
  }

  removeVertexItself(vertex: Vertex) {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      if (v == vertex) {
        this.vertices.splice(i, 1);
      }
    }
  }

  RemoveEdge(from: number, to: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);
    if (toVertex && fromVertex) {
      const indexOfTheEdgeToRemove = this.getVertexIndex(
        fromVertex.adjacent,
        toVertex
      );
      if (indexOfTheEdgeToRemove > -1) {
        fromVertex.adjacent.splice(indexOfTheEdgeToRemove, 1);
      }
    }
  }

  Print() {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      console.log(v.key, ': ');
      for (let j = 0; j < v.adjacent.length; j++) {
        const adj = v.adjacent[j];
        console.log(adj);
      }
    }
  }

  Traverse() {
    const seen: Vertex[] = [];
    for (let v of this.vertices) {
      this.traverse(seen, v);
    }
  }

  traverse(seen: Vertex[], vertex: Vertex) {
    seen.push(vertex);
    console.log(vertex.key, '->');
    if (this.getVertexIndex(seen, vertex) == -1) {
      this.traverse(seen, vertex);
    }
  }

  ReturnDGraphVertices() {
    return this.vertices;
  }

  BFS(v: Vertex) {
    const visited: Vertex[] = [];
    const queue: Vertex[] = [];

    visited.push(v);
    queue.push(v);

    while (queue.length > 0) {
      const vertex = queue[0];
      console.log('BFS stumbled upon ', vertex.key, '\n');
      queue.shift();
      for (let adj of v.adjacent) {
        if (this.getVertexIndex(visited, adj) == -1) {
          visited.push(adj);
          queue.push(adj);
        }
      }
    }
  }

  DFS(v: Vertex) {
    const visited: Vertex[] = [];
    this.dfs(v, visited);
  }

  dfs(v: Vertex, visited: Vertex[]) {
    visited.push(v);
    console.log('DFS stumbled upon ', v.key, '\n');
    for (let adj of v.adjacent) {
      if (this.getVertexIndex(visited, adj) == -1) {
        this.dfs(adj, visited);
      }
    }
  }
}

// const dg = new DGraph();
// dg.AddVertex(1);
// dg.AddVertex(2);
// dg.AddVertex(3);
// dg.AddVertex(4);
// dg.AddEdge(1, 4);
// dg.AddEdge(2, 1);
// dg.AddEdge(3, 2);
// dg.AddEdge(4, 3);
// console.log('After adding vertices and edges:');
// dg.Print();
// dg.RemoveVertex(1);
// console.log('After removing vertex 1:');
// dg.Print();
// dg.RemoveEdge(3, 4);
// console.log('After removing edge from 3 to 4:');
// dg.Print();
// console.log('Traversing:');
// dg.Traverse();
// console.log('');
// console.log('BFS:');
// dg.BFS(dg.ReturnDGraphVertices()[0]);
// console.log('DFS:');
// dg.DFS(dg.ReturnDGraphVertices()[0]);

// Directed graph END
//------------------------------------------------------------------------------

// Undirected graph

class Graph {
  vertices: Vertex[] = [];

  contains(vertices: Vertex[], key: number) {
    for (let v of vertices) {
      if (v.key === key) return true;
    }
    return false;
  }
  getVertex(key: number) {
    for (let v of this.vertices) {
      if (v.key === key) {
        return v;
      }
    }
  }

  getVertexIndex(vertices: Vertex[], vertex: Vertex) {
    for (let i = 0; i < vertices.length; i++) {
      if (vertices[i] === vertex) {
        return i;
      }
    }
    return -1;
  }

  AddVertex(key: number) {
    if (this.contains(this.vertices, key)) return;
    const newVertex = new Vertex(key);
    this.vertices.push(newVertex);
  }

  AddEdge(to: number, from: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);

    if (
      !toVertex ||
      !fromVertex ||
      this.contains(fromVertex.adjacent, toVertex.key)
    ) {
      return;
    }
    fromVertex.adjacent.push(toVertex);
    toVertex.adjacent.push(fromVertex);
  }

  RemoveVertex(key: number) {
    const vertexToRemove = this.getVertex(key);
    if (vertexToRemove) {
      this.removePointersToVertex(vertexToRemove);
      this.removeVertexItself(vertexToRemove);
    }
  }

  removePointersToVertex(vertex: Vertex) {
    for (let v of this.vertices) {
      for (let i = 0; i < v.adjacent.length; i++) {
        const adj = v.adjacent[i];
        if (adj == vertex) {
          v.adjacent.splice(i, 1);
        }
      }
    }
  }

  removeVertexItself(vertex: Vertex) {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      if (v == vertex) {
        this.vertices.splice(i, 1);
      }
    }
  }

  RemoveEdge(from: number, to: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);
    if (toVertex && fromVertex) {
      const indexOfTheEdgeInFromVertex = this.getVertexIndex(
        fromVertex.adjacent,
        toVertex
      );
      const indexOfTheEdgeInToVertex = this.getVertexIndex(
        toVertex.adjacent,
        fromVertex
      );
      if (indexOfTheEdgeInFromVertex > -1) {
        fromVertex.adjacent.splice(indexOfTheEdgeInFromVertex, 1);
      }
      if (indexOfTheEdgeInToVertex > -1) {
        toVertex.adjacent.splice(indexOfTheEdgeInToVertex, 1);
      }
    }
  }

  Print() {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      console.log(v.key, ': ');
      for (let j = 0; j < v.adjacent.length; j++) {
        const adj = v.adjacent[j];
        console.log(adj);
      }
    }
  }

  Traverse() {
    const seen: Vertex[] = [];
    for (let v of this.vertices) {
      this.traverse(seen, v);
    }
  }

  traverse(seen: Vertex[], vertex: Vertex) {
    seen.push(vertex);
    console.log(vertex.key, '->');
    if (this.getVertexIndex(seen, vertex) == -1) {
      this.traverse(seen, vertex);
    }
  }

  ReturnGraphVertices() {
    return this.vertices;
  }

  BFS(v: Vertex) {
    const visited: Vertex[] = [];
    const queue: Vertex[] = [];

    visited.push(v);
    queue.push(v);

    while (queue.length > 0) {
      const vertex = queue[0];
      console.log('BFS stumbled upon ', vertex.key, '\n');
      queue.shift();
      for (let adj of v.adjacent) {
        if (this.getVertexIndex(visited, adj) == -1) {
          visited.push(adj);
          queue.push(adj);
        }
      }
    }
  }

  DFS(v: Vertex) {
    const visited: Vertex[] = [];
    this.dfs(v, visited);
  }

  dfs(v: Vertex, visited: Vertex[]) {
    visited.push(v);
    console.log('DFS stumbled upon ', v.key, '\n');
    for (let adj of v.adjacent) {
      if (this.getVertexIndex(visited, adj) == -1) {
        this.dfs(adj, visited);
      }
    }
  }
}

// const ug = new Graph();
// ug.AddVertex(1);
// ug.AddVertex(2);
// ug.AddVertex(3);
// ug.AddVertex(4);
// ug.AddEdge(1, 4);
// ug.AddEdge(2, 1);
// ug.AddEdge(3, 2);
// ug.AddEdge(4, 3);
// console.log('After adding vertices and edges:');
// ug.Print();
// ug.RemoveVertex(1);
// console.log('After removing vertex 1:');
// ug.Print();
// ug.RemoveEdge(3, 4);
// console.log('After removing edge from 3 to 4:');
// ug.Print();
// console.log('Traversing:');
// ug.Traverse();
// console.log('');
// console.log('BFS:');
// ug.BFS(ug.ReturnGraphVertices()[0]);
// console.log('DFS:');
// ug.DFS(ug.ReturnGraphVertices()[0]);

// Undirected graph END
//------------------------------------------------------------------------------

// Weighted Directed graph

class WeightedEdge {
  weight: number;
  vertex: WeightedVertex;
  constructor(v: WeightedVertex, weight: number) {
    this.vertex = v;
    this.weight = weight;
  }
}

class WeightedVertex {
  key: number;
  adjacent: WeightedEdge[] = [];
  constructor(key: number) {
    this.key = key;
  }
}

class WDGraph {
  vertices: WeightedVertex[] = [];

  convertWeightedEdgesToWeightedVertices(edges: WeightedEdge[]) {
    const ret: WeightedVertex[] = [];
    for (let edge of edges) {
      ret.push(edge.vertex);
    }
    return ret;
  }

  contains(vertices: WeightedVertex[], key: number) {
    for (let v of vertices) {
      if (v.key === key) return true;
    }
    return false;
  }
  getVertex(key: number) {
    for (let v of this.vertices) {
      if (v.key === key) {
        return v;
      }
    }
  }

  getVertexIndex(vertices: WeightedVertex[], vertex: WeightedVertex) {
    for (let i = 0; i < vertices.length; i++) {
      if (vertices[i] === vertex) {
        return i;
      }
    }
    return -1;
  }

  AddVertex(key: number) {
    if (this.contains(this.vertices, key)) return;
    const newVertex = new WeightedVertex(key);
    this.vertices.push(newVertex);
  }

  AddEdge(to: number, from: number, weight: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);

    if (
      !toVertex ||
      !fromVertex ||
      this.contains(
        this.convertWeightedEdgesToWeightedVertices(fromVertex.adjacent),
        toVertex.key
      )
    ) {
      return;
    }
    const newWeightedEdge = new WeightedEdge(toVertex, weight);
    fromVertex.adjacent.push(newWeightedEdge);
  }

  RemoveVertex(key: number) {
    const vertexToRemove = this.getVertex(key);
    if (vertexToRemove) {
      this.removePointersToVertex(vertexToRemove);
      this.removeVertexItself(vertexToRemove);
    }
  }

  removePointersToVertex(vertex: WeightedVertex) {
    for (let v of this.vertices) {
      for (let i = 0; i < v.adjacent.length; i++) {
        const adj = v.adjacent[i];
        if (adj.vertex === vertex) {
          v.adjacent.splice(i, 1);
        }
      }
    }
  }

  removeVertexItself(vertex: WeightedVertex) {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      if (v == vertex) {
        this.vertices.splice(i, 1);
      }
    }
  }

  RemoveEdge(from: number, to: number) {
    const toVertex = this.getVertex(to);
    const fromVertex = this.getVertex(from);
    if (toVertex && fromVertex) {
      const indexOfTheEdgeToRemove = this.getVertexIndex(
        this.convertWeightedEdgesToWeightedVertices(fromVertex.adjacent),
        toVertex
      );
      if (indexOfTheEdgeToRemove > -1) {
        fromVertex.adjacent.splice(indexOfTheEdgeToRemove, 1);
      }
    }
  }

  Print() {
    for (let i = 0; i < this.vertices.length; i++) {
      const v = this.vertices[i];
      console.log(v.key, ': ');
      for (let j = 0; j < v.adjacent.length; j++) {
        const adj = v.adjacent[j];
        console.log(adj.vertex.key, ' weight: ', adj.weight);
      }
    }
  }

  Traverse() {
    const seen: WeightedVertex[] = [];
    for (let v of this.vertices) {
      this.traverse(seen, v);
    }
  }

  traverse(seen: WeightedVertex[], vertex: WeightedVertex) {
    seen.push(vertex);
    console.log(vertex.key, '->');
    if (this.getVertexIndex(seen, vertex) == -1) {
      this.traverse(seen, vertex);
    }
  }

  ReturnWDGraphVertices() {
    return this.vertices;
  }

  BFS(v: WeightedVertex) {
    const visited: WeightedVertex[] = [];
    const queue: WeightedVertex[] = [];

    visited.push(v);
    queue.push(v);

    while (queue.length > 0) {
      const vertex = queue[0];
      console.log('BFS stumbled upon ', vertex.key, '\n');
      queue.shift();
      for (let adj of v.adjacent) {
        if (this.getVertexIndex(visited, adj.vertex) == -1) {
          visited.push(adj.vertex);
          queue.push(adj.vertex);
        }
      }
    }
  }

  DFS(v: WeightedVertex) {
    const visited: WeightedVertex[] = [];
    this.dfs(v, visited);
  }

  dfs(v: WeightedVertex, visited: WeightedVertex[]) {
    visited.push(v);
    console.log('DFS stumbled upon ', v.key, '\n');
    for (let adj of v.adjacent) {
      if (this.getVertexIndex(visited, adj.vertex) == -1) {
        this.dfs(adj.vertex, visited);
      }
    }
  }
}

// const dg = new WDGraph();
// dg.AddVertex(1);
// dg.AddVertex(2);
// dg.AddVertex(3);
// dg.AddVertex(4);
// dg.AddEdge(1, 4, 1);
// dg.AddEdge(2, 1, 2);
// dg.AddEdge(3, 2, 3);
// dg.AddEdge(4, 3, 4);
// console.log('After adding vertices and edges:');
// dg.Print();
// dg.RemoveVertex(1);
// console.log('After removing vertex 1:');
// dg.Print();
// dg.RemoveEdge(3, 4);
// console.log('After removing edge from 3 to 4:');
// dg.Print();
// console.log('Traversing:');
// dg.Traverse();
// console.log('');
// console.log('BFS:');
// dg.BFS(dg.ReturnWDGraphVertices()[0]);
// console.log('DFS:');
// dg.DFS(dg.ReturnWDGraphVertices()[0]);

// Weighted Directed graph END
//------------------------------------------------------------------------------
