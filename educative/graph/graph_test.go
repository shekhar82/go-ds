package graph

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {

	g := CreateEmmptygraph()
	if g.Network == nil && g.Vertices != nil {
		t.Errorf("Graph isn't created")
	}

	g1 := Graph{}

	if g1.Network != nil && g.Vertices != nil {
		t.Errorf("Graph isn't created")
	}

}

func TestAddVertex(t *testing.T) {
	g := CreateEmmptygraph()

	g.AddVertex("DEL", "Delhi Airport")
	g.AddVertex("BLR", "Bengalore Airport")
	g.AddVertex("CHENNAI", "Chennai Airport")
	g.AddVertex("KOLKATTA", "Kolkatta Airport")
	g.AddVertex("MUM", "Mumbai Airport")
	g.AddVertex("LUCKNOW", "Lucknow airport")

	airports := g.ListAllVertices()

	if len(airports) != 6 {
		t.Errorf("Airports not added")
	} else {
		t.Logf("t: %v\n", airports)
	}

}

func TestAddEdges(t *testing.T) {
	g := CreateEmmptygraph()

	g.AddVertex("DEL", "Delhi Airport")
	g.AddVertex("BLR", "Bengalore Airport")
	g.AddVertex("CHENNAI", "Chennai Airport")
	g.AddVertex("KOLKATTA", "Kolkatta Airport")
	g.AddVertex("MUM", "Mumbai Airport")
	g.AddVertex("LUCKNOW", "Lucknow airport")

	g.AddEdge("DEL", "BLR", 1000)
	g.AddEdge("DEL", "MUM", 2000)
	g.AddEdge("BLR", "MUM", 3000)
	t.Logf("t: %v\n", g.PrintGraph())

}
