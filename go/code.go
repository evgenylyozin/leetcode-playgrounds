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
	dg := ds.NewDirectedGraph()
	dg.AddVertex(1)
	dg.AddVertex(2)
	dg.AddVertex(3)
	dg.AddVertex(4)
	dg.AddEdge(1, 4)
	dg.AddEdge(2, 1)
	dg.AddEdge(3, 2)
	dg.AddEdge(4, 3)
	fmt.Println("After adding vertices and edges:")
	dg.Print()
	dg.RemoveVertex(1)
	fmt.Println("After removing vertex 1:")
	dg.Print()

	dg.RemoveEdge(3, 4)
	fmt.Println("After removing edge from 3 to 4:")
	dg.Print()
	fmt.Println("Traversing:")
	dg.Traverse()
	fmt.Println()
	fmt.Println("BFS:")
	dg.BFS(dg.ReturnGraphVertices()[0])
	fmt.Println("DFS:")
	dg.DFS(dg.ReturnGraphVertices()[0])
}
