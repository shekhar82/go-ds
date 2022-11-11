package graph

import "sort"

type AcquaintaceGraph struct {
	Root         []int
	Rank         []int
	Disconnected int
}

type ConnectionTime struct {
	timestamp int
	p1        int
	p2        int
}

type TimestampSorter []ConnectionTime

func (t TimestampSorter) Len() int           { return len(t) }
func (t TimestampSorter) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimestampSorter) Less(i, j int) bool { return t[i].timestamp < t[j].timestamp }

func EarliestAcq(logs [][]int, n int) int {
	connectedFriends := make(map[int]struct{})
	var exists = struct{}{}

	ag := CreateAcuaintaceGraph(n)
	times := make([]ConnectionTime, 0)
	for _, log := range logs {
		times = append(times, ConnectionTime{
			timestamp: log[0],
			p1:        log[1],
			p2:        log[2],
		})
	}

	sort.Sort(TimestampSorter(times))
	for _, connection := range times {
		if ag.ConnectFriends(connection.p1, connection.p2) {
			if _, ok := connectedFriends[connection.p1]; !ok {
				connectedFriends[connection.p1] = exists
			}

			if _, ok := connectedFriends[connection.p2]; !ok {
				connectedFriends[connection.p2] = exists
			}

			if (ag.Disconnected - len(connectedFriends)) == 1 {
				return connection.timestamp
			}

		}

	}

	return -1

}

func CreateAcuaintaceGraph(size int) *AcquaintaceGraph {
	localRoot := make([]int, size)
	localRank := make([]int, size)

	for i := 0; i < size; i++ {
		localRoot[i] = i
		localRank[i] = i
	}

	return &AcquaintaceGraph{
		Root:         localRoot,
		Rank:         localRank,
		Disconnected: size,
	}
}

func (acqG *AcquaintaceGraph) FindRoot(personId int) int {
	if acqG.Root[personId] == personId {
		return personId
	}

	acqG.Root[personId] = acqG.FindRoot(acqG.Root[personId])
	return acqG.Root[personId]
}

func (acqG *AcquaintaceGraph) ConnectFriends(p1Id, p2Id int) bool {
	rootOfP1 := acqG.FindRoot(p1Id)
	rootOfP2 := acqG.FindRoot(p2Id)

	if rootOfP1 != rootOfP2 {
		if acqG.Rank[rootOfP1] > acqG.Rank[rootOfP2] {
			acqG.Root[rootOfP2] = rootOfP1
		} else if acqG.Rank[rootOfP1] < acqG.Rank[rootOfP2] {
			acqG.Root[rootOfP1] = rootOfP2
		} else {
			acqG.Root[rootOfP2] = rootOfP1
			acqG.Rank[rootOfP1] += 1
		}
		return true
	}

	return false
}
