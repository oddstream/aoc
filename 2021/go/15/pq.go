package main

// https://pkg.go.dev/container/heap

type Item struct {
	Point
	risk  int // cumulative risk
	index int // internal use by heap
}

type PriorityQueue []*Item

// need to implement sort interface (Len, Less, Swap)
// and heap interface (Push, Pop)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].risk < pq[j].risk
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, point Point, priority int) {
// 	item.Point = point
// 	item.risk = priority
// 	heap.Fix(pq, item.index)
// }
