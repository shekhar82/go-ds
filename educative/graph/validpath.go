package graph

type node struct {
	id int
}

type edge struct {
	nodeId int
	weight int
}

type LocalGraph struct {
	nodes   map[int]*node
	network map[int][]edge
}

func CreateLocalGraph(nodeCount int) LocalGraph {
	localNodes := make(map[int]*node)

	for i := 0; i < nodeCount; i++ {
		node := new(node)
		node.id = i
		localNodes[i] = node
	}

	return LocalGraph{
		nodes:   localNodes,
		network: make(map[int][]edge),
	}
}

func CreateLocalGraphWithEdges(nodeCount int, edges [][]int) LocalGraph {
	lg := CreateLocalGraph(nodeCount)
	lg.AddEdges(edges)
	return lg
}

func (lg *LocalGraph) AddEdges(edges [][]int) {
	for _, ed := range edges {
		e := edge{nodeId: ed[1], weight: 1}

		if _, ok := lg.network[ed[0]]; !ok {
			lg.network[ed[0]] = make([]edge, 0)
		}

		lg.network[ed[0]] = append(lg.network[ed[0]], e)
	}
}

func (lg *LocalGraph) 
