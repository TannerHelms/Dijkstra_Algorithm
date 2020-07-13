package dijkstra

import (
	"fmt"
	"sort"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

type Edge struct {
	Parent   *Node
	Child    *Node
	Distance int
}

type Node struct {
	Name string
}

const Infinity = int(^uint(0) >> 1)

// Insert the Nodes into the Graph

func (g *Graph) insertNodes(node ...*Node) {
	for _,v := range node {
		g.Nodes = append(g.Nodes, v)
	}
}

//Insert the edges between Node's into the Graph

func (g *Graph) insertEdges(edge ...*Edge) {
	for _,v := range edge {
		g.Edges = append(g.Edges, v)
	}
}

// Function that Prints the edges to the console

func (g *Graph) toString() {
	for _,v := range g.Edges {
		fmt.Printf("%v %v %v\n", *v.Parent, v.Distance, *v.Child)
	}
}


// Returns a table that shows each Node and the distance from the start Node

func (g *Graph) NewDistanceTable(startNode *Node) map[*Node]int {
	costTable := make(map[*Node]int)

	//Set your start Node as zero since the distance is zero

	costTable[startNode] = 0

	for _,node := range g.Nodes {

		//Set all Nodes other then the start Node to Infinity

		if node != startNode {
			costTable[node] = Infinity
		}


	}
	return costTable
}

// This function returns a slice of all the Nodes neighbors

func (g *Graph) getNodeEdges(node *Node) (edges []*Edge) {

	// Sort through all the edges in Graph

	for _, edge := range g.Edges {

		// If the Edge parent is equal to the Node

		if edge.Parent == node {
			edges = append(edges,edge)
		}
	}
	return edges
}

// Function that returns the closest Node with the lease amount of distance
// If the Node hasn't been visited yet

func getClosestNonVisitedNode (costTable map[*Node]int, visited []*Node) *Node {
	type CostTableToSort struct {
		Node     *Node
		Distance int
	}
	var sorted []CostTableToSort

	// Check if the node has already been visited

	for node,cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}

		//If not add it to the slice

		if !isVisited {
			sorted = append(sorted, CostTableToSort{node,cost})
		}
	}

	//Sort the slice node distances

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Distance < sorted[j].Distance
	})

	//Return the first value in the ascending ordered slice

	return sorted[0].Node

}


