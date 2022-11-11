package graph

func CountComponents(n int, edges [][]int) int {
	cg := CreateComponentGraph(n)

	for _, edge := range edges {
		cg.ConnectNodes(edge[0], edge[1])
	}

	disjoinedCounts := 0
	for i := 0; i < len(cg.rootNodes); i++ {
		if cg.rootNodes[i] == i {
			disjoinedCounts += 1
		}
	}

	return disjoinedCounts
}

type ComponentGraph struct {
	rootNodes []int
	rank      []int
}

func CreateComponentGraph(size int) *ComponentGraph {
	localNodes := make([]int, size)
	localRanks := make([]int, size)

	for i := 0; i < size; i++ {
		localNodes[i] = i
		localRanks[i] = i
	}

	return &ComponentGraph{
		rootNodes: localNodes,
		rank:      localRanks,
	}
}

func (cg *ComponentGraph) FindRoot(nodeId int) int {
	if nodeId == cg.rootNodes[nodeId] {
		return nodeId
	}

	cg.rootNodes[nodeId] = cg.FindRoot(cg.rootNodes[nodeId])
	return cg.rootNodes[nodeId]
}

func (cg *ComponentGraph) ConnectNodes(node1, node2 int) bool {
	rootOfNode1 := cg.FindRoot(node1)
	rooOfNode2 := cg.FindRoot(node2)

	if rootOfNode1 != rooOfNode2 {
		if cg.rank[rootOfNode1] > cg.rank[rooOfNode2] {
			cg.rootNodes[rooOfNode2] = rootOfNode1
		} else if cg.rank[rootOfNode1] < cg.rank[rooOfNode2] {
			cg.rootNodes[rootOfNode1] = rooOfNode2
		} else {
			cg.rootNodes[rooOfNode2] = rootOfNode1
			cg.rank[rootOfNode1] += 1
		}

		return true
	}

	return false
}
