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
		bst.Preorder(node.left)
		bst.Preorder(node.right)
	}
}

// Postorder bst traversal
func (bst *BST) Postorder(node *BSTNode) {
	if node != nil {
		bst.Postorder(node.left)
		bst.Postorder(node.right)
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

// Directed graph structure

// DGraph structure
type DGraph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// NewDirectedGraph creates new directed graph and returns the pointer to it
func NewDirectedGraph() *DGraph {
	return &DGraph{}
}

// AddVertex will add a vertex to a graph
func (g *DGraph) AddVertex(vertex int) error {
	if g.contains(g.vertices, vertex) {
		err := fmt.Errorf("Vertex %d already exists", vertex)
		return err
	}
	v := &Vertex{
		key: vertex,
	}
	g.vertices = append(g.vertices, v)
	return nil
}

// AddEdge will add ad endge from a vertex to a vertex
func (g *DGraph) AddEdge(to, from int) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("Not a valid edge from %d ---> %d", from, to)
	} else if g.contains(fromVertex.adjacent, toVertex.key) {
		return fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.key, toVertex.key)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		return nil
	}
}

// RemoveVertex deletes the whole vertex with specified data in its key from the graph
func (g *DGraph) RemoveVertex(key int) {
	vertexToRemove := g.getVertex(key)
	g.removePointersToVertex(vertexToRemove)
	g.removeVertexItself(vertexToRemove)
}

// RemoveEdge deletes specific edge from one vertex to another if any
func (g *DGraph) RemoveEdge(from, to int) {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if fromVertex != nil && toVertex != nil {
		indexOfTheEdgeToRemove := g.getVertexIndex(fromVertex.adjacent, toVertex)
		if indexOfTheEdgeToRemove > -1 {
			var rest []*Vertex
			if indexOfTheEdgeToRemove+1 < len(fromVertex.adjacent) {
				rest = fromVertex.adjacent[indexOfTheEdgeToRemove+1:]
			}
			fromVertex.adjacent = append(fromVertex.adjacent[:indexOfTheEdgeToRemove], rest...)
		}
	}
}

// removePointersToVertex will find and remove all pointers to due for deletion vertex
func (g *DGraph) removePointersToVertex(vertex *Vertex) {
	for _, v := range g.vertices {
		for j, adj := range v.adjacent {
			if adj == vertex {
				var rest []*Vertex
				if j+1 < len(v.adjacent) {
					rest = v.adjacent[j+1:]
				}
				v.adjacent = append(v.adjacent[:j], rest...)
			}
		}
	}
}

func (g *DGraph) removeVertexItself(vertex *Vertex) {
	for i, v := range g.vertices {
		if v == vertex {
			var rest []*Vertex
			if i+1 < len(g.vertices) {
				rest = g.vertices[i+1:]
			}
			g.vertices = append(g.vertices[:i], rest...)
		}
	}
}

// getVertexIndex will return a vertex index if exists in current slice of vertices or return -1
func (g *DGraph) getVertexIndex(vertices []*Vertex, vertex *Vertex) int {
	for i, v := range vertices {
		if v == vertex {
			return i
		}
	}
	return -1
}

// getVertex will return a vertex point if exists or return nil
func (g *DGraph) getVertex(vertex int) *Vertex {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *DGraph) contains(v []*Vertex, key int) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}
	return false
}

// Print the directed graph
func (g *DGraph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%d ", v.key)
		}
		fmt.Println()
	}
}

// Traverse the graph
func (g *DGraph) Traverse() {
	var seen []*Vertex
	for _, v := range g.vertices {
		g.traverse(seen, v)
	}
}

func (g *DGraph) traverse(seen []*Vertex, vertex *Vertex) {
	seen = append(seen, vertex)
	fmt.Print(vertex.key)
	fmt.Print(",")
	if g.getVertexIndex(seen, vertex) == -1 {
		g.traverse(seen, vertex)
	}
}

// ReturnDGraphVertices returns a slice of vertices currently in the graph
func (g *DGraph) ReturnDGraphVertices() []*Vertex {
	return g.vertices
}

// BFS - breadth first search in the directed graph from starting Vertex, doesn't see disconnected vertices
func (g *DGraph) BFS(v *Vertex) {
	var visited []*Vertex
	var queue []*Vertex

	visited = append(visited, v)
	queue = append(queue, v)
	for len(queue) > 0 {
		vertex := queue[0]
		fmt.Printf("BFS stumbled upon %v\n", vertex)
		if len(queue) > 1 {
			queue = append(queue[1:])
		} else {
			queue = []*Vertex{}
		}
		for _, adj := range v.adjacent {
			if g.getVertexIndex(visited, adj) == -1 {
				visited = append(visited, adj)
				queue = append(queue, adj)
			}
		}
	}
}

// DFS - depth first search in the directed graph from starting Vertex, doesn't see disconnected vertices
func (g *DGraph) DFS(v *Vertex) {
	var visited []*Vertex

	g.dfs(v, visited)
}

func (g *DGraph) dfs(v *Vertex, visited []*Vertex) {
	visited = append(visited, v)
	fmt.Printf("DFS stumbled upon %v\n", v)

	for _, adj := range v.adjacent {
		if g.getVertexIndex(visited, adj) == -1 {
			g.dfs(adj, visited)
		}
	}
}

// import (
// 	"fmt"
// 	ds "leetcode/structures"
// )

// func main() {
// 	dg := ds.NewDirectedGraph()
// 	dg.AddVertex(1)
// 	dg.AddVertex(2)
// 	dg.AddVertex(3)
// 	dg.AddVertex(4)
// 	dg.AddEdge(1, 4)
// 	dg.AddEdge(2, 1)
// 	dg.AddEdge(3, 2)
// 	dg.AddEdge(4, 3)
// 	fmt.Println("After adding vertices and edges:")
// 	dg.Print()
// 	dg.RemoveVertex(1)
// 	fmt.Println("After removing vertex 1:")
// 	dg.Print()

// 	dg.RemoveEdge(3, 4)
// 	fmt.Println("After removing edge from 3 to 4:")
// 	dg.Print()
// 	fmt.Println("Traversing:")
// 	dg.Traverse()
// 	fmt.Println()
// 	fmt.Println("BFS:")
// 	dg.BFS(dg.ReturnDGraphVertices()[0])
// 	fmt.Println("DFS:")
// 	dg.DFS(dg.ReturnDGraphVertices()[0])
// }

// Directed graph structure END
//------------------------------------------------------------------------------

// Undirected Graph structure

// Graph - undirected graph structure
type Graph struct {
	vertices []*Vertex
}

// NewUndirectedGraph creates new undirected graph and returns the pointer to it
func NewUndirectedGraph() *Graph {
	return &Graph{}
}

// AddVertex will add a vertex to a graph
func (g *Graph) AddVertex(vertex int) error {
	if g.contains(g.vertices, vertex) {
		err := fmt.Errorf("Vertex %d already exists", vertex)
		return err
	}
	v := &Vertex{
		key: vertex,
	}
	g.vertices = append(g.vertices, v)
	return nil
}

// AddEdge will add ad endge from a vertex to a vertex
func (g *Graph) AddEdge(to, from int) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("Not a valid edge from %d ---> %d", from, to)
	} else if g.contains(fromVertex.adjacent, toVertex.key) {
		return fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.key, toVertex.key)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
		return nil
	}
}

// RemoveVertex deletes the whole vertex with specified data in its key from the graph
func (g *Graph) RemoveVertex(key int) {
	vertexToRemove := g.getVertex(key)
	g.removePointersToVertex(vertexToRemove)
	g.removeVertexItself(vertexToRemove)
}

// RemoveEdge deletes specific edge from one vertex to another if any
func (g *Graph) RemoveEdge(from, to int) {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if fromVertex != nil && toVertex != nil {
		indexOfTheEdgeInFromVertex := g.getVertexIndex(fromVertex.adjacent, toVertex)
		indexOfTheEdgeInToVertex := g.getVertexIndex(toVertex.adjacent, fromVertex)
		if indexOfTheEdgeInFromVertex > -1 {
			var rest []*Vertex
			if indexOfTheEdgeInFromVertex+1 < len(fromVertex.adjacent) {
				rest = fromVertex.adjacent[indexOfTheEdgeInFromVertex+1:]
			}
			fromVertex.adjacent = append(fromVertex.adjacent[:indexOfTheEdgeInFromVertex], rest...)
		}

		if indexOfTheEdgeInToVertex > -1 {
			var rest []*Vertex
			if indexOfTheEdgeInToVertex+1 < len(toVertex.adjacent) {
				rest = toVertex.adjacent[indexOfTheEdgeInToVertex+1:]
			}
			toVertex.adjacent = append(toVertex.adjacent[:indexOfTheEdgeInToVertex], rest...)
		}
	}
}

// removePointersToVertex will find and remove all pointers to due for deletion vertex
func (g *Graph) removePointersToVertex(vertex *Vertex) {
	for _, v := range g.vertices {
		for j, adj := range v.adjacent {
			if adj == vertex {
				var rest []*Vertex
				if j+1 < len(v.adjacent) {
					rest = v.adjacent[j+1:]
				}
				v.adjacent = append(v.adjacent[:j], rest...)
			}
		}
	}
}

func (g *Graph) removeVertexItself(vertex *Vertex) {
	for i, v := range g.vertices {
		if v == vertex {
			var rest []*Vertex
			if i+1 < len(g.vertices) {
				rest = g.vertices[i+1:]
			}
			g.vertices = append(g.vertices[:i], rest...)
		}
	}
}

// getVertexIndex will return a vertex index if exists in current slice of vertices or return -1
func (g *Graph) getVertexIndex(vertices []*Vertex, vertex *Vertex) int {
	for i, v := range vertices {
		if v == vertex {
			return i
		}
	}
	return -1
}

// getVertex will return a vertex point if exists or return nil
func (g *Graph) getVertex(vertex int) *Vertex {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) contains(v []*Vertex, key int) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}
	return false
}

// Print the directed graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%d ", v.key)
		}
		fmt.Println()
	}
}

// Traverse the graph
func (g *Graph) Traverse() {
	var seen []*Vertex
	for _, v := range g.vertices {
		g.traverse(seen, v)
	}
}

func (g *Graph) traverse(seen []*Vertex, vertex *Vertex) {
	seen = append(seen, vertex)
	fmt.Print(vertex.key)
	fmt.Print(",")
	if g.getVertexIndex(seen, vertex) == -1 {
		g.traverse(seen, vertex)
	}
}

// ReturnGraphVertices returns a slice of vertices currently in the graph
func (g *Graph) ReturnGraphVertices() []*Vertex {
	return g.vertices
}

// BFS - breadth first search in the directed graph from starting Vertex, doesn't see disconnected vertices
func (g *Graph) BFS(v *Vertex) {
	var visited []*Vertex
	var queue []*Vertex

	visited = append(visited, v)
	queue = append(queue, v)
	for len(queue) > 0 {
		vertex := queue[0]
		fmt.Printf("BFS stumbled upon %v\n", vertex)
		if len(queue) > 1 {
			queue = append(queue[1:])
		} else {
			queue = []*Vertex{}
		}
		for _, adj := range v.adjacent {
			if g.getVertexIndex(visited, adj) == -1 {
				visited = append(visited, adj)
				queue = append(queue, adj)
			}
		}
	}
}

// DFS - depth first search in the directed graph from starting Vertex, doesn't see disconnected vertices
func (g *Graph) DFS(v *Vertex) {
	var visited []*Vertex

	g.dfs(v, visited)
}

func (g *Graph) dfs(v *Vertex, visited []*Vertex) {
	visited = append(visited, v)
	fmt.Printf("DFS stumbled upon %v\n", v)

	for _, adj := range v.adjacent {
		if g.getVertexIndex(visited, adj) == -1 {
			g.dfs(adj, visited)
		}
	}
}

// import (
// 	"fmt"
// 	ds "leetcode/structures"
// )

// func main() {
// 	ug := ds.NewUndirectedGraph()
// 	ug.AddVertex(1)
// 	ug.AddVertex(2)
// 	ug.AddVertex(3)
// 	ug.AddVertex(4)
// 	ug.AddEdge(1, 4)
// 	ug.AddEdge(2, 1)
// 	ug.AddEdge(3, 2)
// 	ug.AddEdge(4, 3)
// 	fmt.Println("After adding vertices and edges:")
// 	ug.Print()
// 	ug.RemoveVertex(1)
// 	fmt.Println("After removing vertex 1:")
// 	ug.Print()

// 	ug.RemoveEdge(3, 4)
// 	fmt.Println("After removing edge from 3 to 4:")
// 	ug.Print()
// 	fmt.Println("Traversing:")
// 	ug.Traverse()
// 	fmt.Println()
// 	fmt.Println("BFS:")
// 	ug.BFS(ug.ReturnGraphVertices()[0])
// 	fmt.Println("DFS:")
// 	ug.DFS(ug.ReturnGraphVertices()[0])
// }

// Undirected Grapg structure END
//------------------------------------------------------------------------------
