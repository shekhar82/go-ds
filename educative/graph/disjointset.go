package graph

type UnionFindV1 struct {
	root []int
}

func CreateUnionFindV1(size int) *UnionFindV1 {
	val := make([]int, size)
	for i := 0; i < size; i++ {
		val[i] = i
	}
	return &UnionFindV1{
		root: val,
	}
}

func (union *UnionFindV1) Find(x int) int {
	if x < len(union.root) {
		return union.root[x]
	}
	return -1
}

func (union *UnionFindV1) Union(x, y int) {
	rootX := union.Find(x)
	rootY := union.Find(y)

	if rootX != -1 && rootY != -1 {
		if rootX != rootY {
			for i := 0; i < len(union.root); i++ {
				if union.root[i] == rootY {
					union.root[i] = rootX
				}
			}
		}
	}
}

func (union *UnionFindV1) Connected(x, y int) bool {
	return union.Find(x) != -1 && union.Find(y) != -1 && (union.Find(x) == union.Find(y))
}

type UnionFindV2 struct {
	root []int
}

func CreateUnionFindV2(size int) *UnionFindV2 {
	val := make([]int, size)
	for i := 0; i < size; i++ {
		val[i] = i
	}

	return &UnionFindV2{
		root: val,
	}
}

func (union *UnionFindV2) Find(x int) int {
	if x < len(union.root) {
		for x != union.root[x] {
			x = union.root[x]
		}
		return union.root[x]
	}
	return -1
}

func (union *UnionFindV2) Union(x, y int) {
	rootX := union.Find(x)
	rootY := union.Find(y)

	if rootX != -1 && rootY != -1 && rootX != rootY {
		union.root[rootY] = rootX
	}
}

func (union *UnionFindV2) Connected(x, y int) bool {
	rootX := union.Find(x)
	rootY := union.Find(y)

	return rootX != -1 && rootY != -1 && rootX == rootY
}

type UnionFindV3 struct {
	root []int
	rank []int
}

func CreateUnionFindV3(size int) *UnionFindV3 {
	rootLocal := make([]int, size)
	rankLocal := make([]int, size)

	for i := 0; i < size; i++ {
		rootLocal[i] = i
		rankLocal[i] = 1
	}

	return &UnionFindV3{
		root: rootLocal,
		rank: rankLocal,
	}
}

func (union *UnionFindV3) Find(x int) int {
	if x < len(union.root) {
		if x == union.root[x] {
			return x
		}
		union.root[x] = union.Find(union.root[x])
		return union.root[x]
	}
	return -1
}

func (union *UnionFindV3) Union(x, y int) {
	rootX := union.Find(x)
	rootY := union.Find(y)

	if rootX != -1 && rootY != -1 && rootX != rootY {
		if union.rank[rootX] < union.rank[rootY] {
			union.root[rootX] = rootY
		} else if union.rank[rootX] > union.rank[rootY] {
			union.root[rootY] = rootX
		} else {
			union.root[rootY] = rootX
			union.rank[rootX] += 1
		}
	}
}

func (union *UnionFindV3) Connected(x, y int) bool {
	rootX := union.Find(x)
	rootY := union.Find(y)

	return rootX != -1 && rootY != -1 && rootX == rootY
}
