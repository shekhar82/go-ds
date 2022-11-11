package graph

import "testing"

func TestUnionFindType1(t *testing.T) {
	uf := CreateUnionFindV3(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)

	if uf.Connected(1, 5) != true {
		t.Errorf("Failed at connection between 1 and 5")
	}

	if uf.Connected(5, 7) != true {
		t.Errorf("Failed at connection between 5 and 7")
	}

	if uf.Connected(4, 9) != false {
		t.Errorf("Failed at connection between 4 and 9")
	}

	uf.Union(9, 4)
	if uf.Connected(4, 9) != true {
		t.Errorf("Failed at connection between 4 and 9")
	}
}
