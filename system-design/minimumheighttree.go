package systemdesign

import "container/list"

func FindMinHeightTrees(n int, edges [][]int) []int {
	result := make([]int, 0)
	graph := CreateGraph(n, edges)

	minHeight := 20000
	heights := make([]StackNode, 0)

	for key := range graph.Nodes {
		height := graph.Height(key)
		heights = append(heights, StackNode{ID: key, Height: height})
		if minHeight >= height {
			minHeight = height
		}
	}

	for _, h := range heights {
		if h.Height == minHeight {
			result = append(result, h.ID)
		}
	}
	return result
}

type Node struct {
	ID int
}

type Edge struct {
	DestinationNodeID int
	Weight            int
}

type Graph struct {
	Nodes map[int]*Node
	Edges map[int][]Edge
}

func CreateGraph(nodeCount int, edges [][]int) *Graph {
	nodes := make(map[int]*Node)

	network := make(map[int][]Edge)

	for i := 0; i < nodeCount; i++ {
		node := &Node{ID: i}
		nodes[i] = node
		network[i] = make([]Edge, 0)
	}

	for _, values := range edges {
		// This is where edge would be constructed and edges network would be established
		sourceToDestinationEdge := Edge{DestinationNodeID: values[1], Weight: 1}
		destinationToSourceNode := Edge{DestinationNodeID: values[0], Weight: 1}

		network[values[0]] = append(network[values[0]], sourceToDestinationEdge)
		network[values[1]] = append(network[values[1]], destinationToSourceNode)
	}
	return &Graph{
		Nodes: nodes,
		Edges: network,
	}
}

type Stack struct {
	dll *list.List
}

func NewStack() *Stack {
	return &Stack{dll: list.New()}
}

func (s *Stack) Push(x interface{}) {
	s.dll.PushBack(x)
}

func (s *Stack) Pop() interface{} {
	if s.dll.Len() == 0 {
		return nil
	}

	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.dll.Len() == 0
}

type StackNode struct {
	ID     int
	Height int
}

func (g *Graph) Height(nodeId int) int {
	node := StackNode{ID: nodeId, Height: 0}
	height := 0
	s := NewStack()
	s.Push(node)

	visited := make(map[int]bool)

	for !s.IsEmpty() {
		popped := s.Pop()
		if popped != nil {
			poppedNode := popped.(StackNode)

			for _, connectedNodes := range g.Edges[poppedNode.ID] {
				if _, ok := visited[connectedNodes.DestinationNodeID]; !ok {
					stackedConnectedNode := StackNode{ID: connectedNodes.DestinationNodeID, Height: poppedNode.Height + 1}
					s.Push(stackedConnectedNode)
				}
			}
			if height < poppedNode.Height {
				height = poppedNode.Height
			}

			visited[poppedNode.ID] = true

		}
	}

	return height
}
