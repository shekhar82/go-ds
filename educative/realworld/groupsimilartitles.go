package realworld

import "strconv"

func computeHashKey(s string) string {
	count := make([]int, 26)

	for _, c := range s {
		idx := c - 'a'
		count[idx]++
	}

	key := ""
	for i := 0; i < 26; i++ {
		key += "#"
		key += strconv.Itoa(count[i])
	}

	return key
}

// prepareGroupMap Will group together strings that have same hashkey into a slice
func prepareGroupMap(strs []string) *map[string][]string {
	res := make(map[string][]string)

	for _, s := range strs {
		sKey := computeHashKey(s)
		res[sKey] = append(res[sKey], s)
	}

	return &res
}

func FindGroupForAWord(words []string, query string) (bool, []string) {
	wordsMap := prepareGroupMap(words)

	searchKey := computeHashKey(query)

	val, found := (*wordsMap)[searchKey]

	if found {
		return true, val
	} else {
		return false, nil
	}
}
