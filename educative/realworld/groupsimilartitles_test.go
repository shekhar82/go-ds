package realworld

import (
	"reflect"
	"testing"
)

type GroupMapTestData struct {
	name string
	args []string
	want *map[string][]string
}

func Test_prepareGroupMap(t *testing.T) {
	tests := generateGroupMapTestData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareGroupMap(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareGroupMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateGroupMapTestData() []GroupMapTestData {
	testData := make([]GroupMapTestData, 0)

	input1 := []string{"duel", "dule", "deul", "abac", "aabc", "baca"}
	output1 := make(map[string][]string)

	output1["#2#1#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"] = []string{"abac", "aabc", "baca"}
	output1["#0#0#0#1#1#0#0#0#0#0#0#1#0#0#0#0#0#0#0#0#1#0#0#0#0#0"] = []string{"duel", "dule", "deul"}

	test1 := GroupMapTestData{
		name: "test1",
		args: input1,
		want: &output1,
	}

	testData = append(testData, test1)
	return testData
}

func Test_computeHashKey(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"abac"}, "#2#1#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
		{"Test2", args{"aabc"}, "#2#1#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
		{"Test3", args{"baca"}, "#2#1#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
		{"Test4", args{"caab"}, "#2#1#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
		{"Test5", args{"duel"}, "#0#0#0#1#1#0#0#0#0#0#0#1#0#0#0#0#0#0#0#0#1#0#0#0#0#0"},
		{"Test6", args{"dule"}, "#0#0#0#1#1#0#0#0#0#0#0#1#0#0#0#0#0#0#0#0#1#0#0#0#0#0"},
		{"Test7", args{"deul"}, "#0#0#0#1#1#0#0#0#0#0#0#1#0#0#0#0#0#0#0#0#1#0#0#0#0#0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeHashKey(tt.args.s); got != tt.want {
				t.Errorf("computeHashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGroupForAWord(t *testing.T) {
	type args struct {
		words []string
		query string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []string
	}{
		{"Test1", args{words: []string{"duel", "dule", "deul", "abac", "aabc", "baca"}, query: "acab"}, true, []string{"abac", "aabc", "baca"}},
		{"Test2", args{words: []string{"duel", "dule", "deul", "abac", "aabc", "baca"}, query: "ledu"}, true, []string{"duel", "dule", "deul"}},
		{"Test3", args{words: []string{"duel", "dule", "deul", "abac", "aabc", "baca"}, query: "hello"}, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindGroupForAWord(tt.args.words, tt.args.query)
			if got != tt.want {
				t.Errorf("FindGroupForAWord() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindGroupForAWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
