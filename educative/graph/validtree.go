package graph

type ValidTree struct {
	rootNodes []int
	rank      []int
}

func CreateValidTree(noOfNodes int) *ValidTree {
	localNodes := make([]int, noOfNodes)
	localRanks := make([]int, noOfNodes)

	for i := 0; i < noOfNodes; i++ {
		localNodes[i] = i
		localRanks[i] = 1
	}

	return &ValidTree{
		rootNodes: localNodes,
		rank:      localRanks,
	}
}

func (t *ValidTree) FindRoot(nodeId int) int {
	if nodeId == t.rootNodes[nodeId] {
		return nodeId
	}

	t.rootNodes[nodeId] = t.FindRoot(t.rootNodes[nodeId])
	return t.rootNodes[nodeId]
}

func (t *ValidTree) ConnectNodeAndDetectCycle(node1, node2 int) bool {
	rootOfNode1 := t.FindRoot(node1)
	rootOfNode2 := t.FindRoot(node2)

	if rootOfNode1 != rootOfNode2 {

		if t.rank[rootOfNode1] > t.rank[rootOfNode2] {
			t.rootNodes[rootOfNode2] = rootOfNode1
		} else if t.rank[rootOfNode1] < t.rank[rootOfNode2] {
			t.rootNodes[rootOfNode1] = rootOfNode2
		} else {
			t.rootNodes[rootOfNode2] = rootOfNode1
			t.rank[rootOfNode1] += 1
		}
		return true
	}

	return false
}

func ValidTreeDetection(n int, edges [][]int) bool {
	tree := CreateValidTree(n)

	for _, edge := range edges {
		if !tree.ConnectNodeAndDetectCycle(edge[0], edge[1]) {
			return false
		}
	}

	countOfDistinctRoots := 0

	for i := 0; i < n; i++ {
		if tree.rootNodes[i] == i {
			countOfDistinctRoots += 1
		}
	}

	return countOfDistinctRoots == 1
}
