package graph

import (
	"fmt"
)

type Vertex struct {
	ID   string
	Name string
}

type Edge struct {
	Weight float32
}

type Graph struct {
	Vertices   map[string]*Vertex
	Network    map[string]map[string]Edge
	IsDirected bool
}

func CreateEmmptygraph() *Graph {
	g := Graph{
		Vertices:   make(map[string]*Vertex),
		Network:    make(map[string]map[string]Edge),
		IsDirected: false,
	}
	return &g
}

func (g *Graph) AddVertex(id, name string) bool {
	v := &Vertex{id, name}

	if _, ok := g.Vertices[id]; ok {
		return false
	} else {
		g.Vertices[id] = v
		g.Network[id] = make(map[string]Edge)
	}
	return true
}

func (g *Graph) ListAllVertices() []string {

	vertices := make([]string, 0)
	for k := range g.Vertices {
		vertices = append(vertices, k)
	}

	return vertices
}

func (g *Graph) isVertexInThere(vertextId string) bool {
	_, ok := g.Vertices[vertextId]
	return ok
}

func (g *Graph) AddEdge(source, destination string, weight float32) bool {
	return g.addEdge(source, destination, weight, true)
}

func (g *Graph) addEdge(source, destination string, weight float32, addReverse bool) bool {
	if g.isVertexInThere(source) && g.isVertexInThere(destination) {
		// Check if there are any edges from source vertex
		if edges, ok := g.Network[source]; ok {
			// Check if there is an edge already between source and destination.
			// If yes then modify the wieght
			if edge, ok := edges[destination]; ok {
				edge.Weight = weight
			} else {
				// If not then add an edge
				edges[destination] = Edge{Weight: weight}
			}
		} else {
			// Add en edge in there
			g.Network[source][destination] = Edge{Weight: weight}
		}

		// Connecting another reverse edge from destination to source if it isn't a directed graph
		if !g.IsDirected && addReverse {
			g.addEdge(destination, source, weight, false)
		}
		return true
	}
	return false
}

func (g *Graph) PrintGraph() string {

	s := ""
	for k, v := range g.Network {
		s += fmt.Sprintf("%s ->", k)
		if len(v) != 0 {
			for key, edge := range v {
				s += fmt.Sprintf("%s : %f, ", key, edge.Weight)
			}
		} else {
			s += "{}"
		}

		s += "\n"
	}
	return s
}


