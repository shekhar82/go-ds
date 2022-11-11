package systemdesign

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	numsInStr := make([]string, 0)

	for _, val := range nums {
		numsInStr = append(numsInStr, strconv.Itoa(val))
	}

	sort.Slice(numsInStr, func(i int, j int) bool {
		x, _ := strconv.Atoi(numsInStr[i] + numsInStr[j])
		y, _ := strconv.Atoi(numsInStr[j] + numsInStr[i])

		return x > y
	})

	var b bytes.Buffer

	for _, val := range numsInStr {
		b.WriteString(val)
	}

	finalNumber := strings.TrimLeft(b.String(), "0")

	if len(finalNumber) == 0 {
		return "0"
	}
	return finalNumber

}
