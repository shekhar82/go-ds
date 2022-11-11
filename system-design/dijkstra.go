package systemdesign

import "math"

type DNode struct {
	ID string
}

type DEdge struct {
	Source *DNode
	Dest   *DNode
	Weight float32
}
type DGraph struct {
	DNodes  map[string]*DNode
	Network map[string][]DEdge
}

func CreateDGraph() *DGraph {
	return &DGraph{
		DNodes:  make(map[string]*DNode),
		Network: make(map[string][]DEdge),
	}
}

func (dg *DGraph) AddNode(id string) bool {
	if _, ok := dg.DNodes[id]; !ok {
		dg.DNodes[id] = &DNode{ID: id}
		return true
	}
	return false
}

func (dg *DGraph) AddEdge(sourceId, destId string, weight float32, isDirected bool) {
	sourceNode := dg.DNodes[sourceId]
	destNode := dg.DNodes[destId]
	dg.Network[sourceId] = append(dg.Network[sourceId], DEdge{
		Source: sourceNode,
		Dest:   destNode,
		Weight: weight,
	})

	if !isDirected {
		dg.Network[destId] = append(dg.Network[destId], DEdge{
			Source: destNode,
			Dest:   sourceNode,
			Weight: weight,
		})
	}
}

func (dg *DGraph) ShortestPath(startNodeId, endNodeId string) (float32, string) {
	costs := make(map[string]float32)
	parent := make(map[string]string)
	processed := make(map[string]bool)
	//init cost map for all nodes and assign weights
	parent[startNodeId] = startNodeId
	for id := range dg.DNodes {
		if id == startNodeId {
			costs[startNodeId] = 0
		} else {
			costs[id] = math.MaxFloat32
		}
	}

	nodeId := findLowestCostNode(costs, processed)

	for len(nodeId) != 0 {
		cost := costs[nodeId]
		for _, edge := range dg.Network[nodeId] {
			newCost := cost + edge.Weight
			if costs[edge.Dest.ID] > newCost {
				costs[edge.Dest.ID] = newCost
				parent[edge.Dest.ID] = nodeId
			}
		}
		processed[nodeId] = true
		nodeId = findLowestCostNode(costs, processed)
	}

}

func findLowestCostNode(costs map[string]float32, processed map[string]bool) string {
	var lowestCost float32 = math.MaxFloat32
	lowestCostNode := ""

	for n, c := range costs {
		_, ok := processed[n]
		if c < float32(lowestCost) && !ok {
			lowestCost = c
			lowestCostNode = n
		}
	}

	return lowestCostNode
}
