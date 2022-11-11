package graph

func CountNumberOfProvinces(isConnected [][]int) int {
	totalProvinceCount := 0

	unionFind := CreateUnionFindV3(len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				unionFind.Union(i, j)
			}
		}
	}

	for i := 0; i < len(unionFind.root); i++ {
		if unionFind.root[i] == i {
			totalProvinceCount += 1
		}
	}

	return totalProvinceCount

}
