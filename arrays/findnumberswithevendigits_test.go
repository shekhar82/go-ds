package arrays

import "testing"

type input struct {
	numbers []int
}

func TestFindNum(t *testing.T) {
	inputs := make([]input, 2)

	ip1 := input{
		numbers: []int{12, 345, 2, 6, 7896},
	}
	ip2 := input{
		numbers: []int{555, 901, 482, 1771},
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			inputs[i] = ip1
		} else if i == 1 {
			inputs[i] = ip2
		}
	}

	expected := []int{2, 1}

	for i := 0; i < 2; i++ {
		actual := FindNumbers(inputs[i].numbers)
		if actual != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], actual)
		}
	}

}
